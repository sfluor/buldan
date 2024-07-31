package main

import (
	"context"
	"crypto/rand"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"path"
	"strings"
	"sync"

	"nhooyr.io/websocket"
	// "nhooyr.io/websocket/wsjson"
)

func errOut(w http.ResponseWriter, msg string, code int) {
	log.Printf("[ERROR] Code:%d | %s", code, msg)
	http.Error(w, msg, code)
}

func errOutWS(c *websocket.Conn, msg string, code websocket.StatusCode) {
	log.Printf("[ERROR:WS] Code:%d | %s", code, msg)
	if err := c.Close(code, msg); err != nil {
		log.Printf("Failed to close websocket: %s", err)
	}
}

type Player struct {
	Name  string
	Admin bool
	// Internal
	conn *websocket.Conn
}

type EventType string

const (
	EventTypePlayers  = "players"
	EventTypeGuess    = "guess"
	EventTypeLoose    = "loose"
	EventTypeNewRound = "new-round"
	EventTypeEndRound = "end-round"
	EventTypeEnd      = "end"
)

type ClientEventType string

const (
	ClientEventTypeGuess     = "guess"
	ClientEventTypeStartGame = "start-game"
)

type Round struct {
	Round   int

	Players []Player

    CurrentPlayerIdx int

	Letter  string // string to ease the conversion to ASCII in JSON

	Guesses []Guess

	// Duplicate it since this one isn't serialized even
	// though we could infer it from the one below
	remainingCountries map[Country]struct{}
	Remaining          int
}

type EventRound struct {
	Type  EventType
	Round Round
}

type Guess struct {
	Guess   string
	Player  string
	Correct bool
}

type EventPlayers struct {
	Type    EventType
	Players []Player
}

type ClientEvent struct {
	Type  ClientEventType
	Guess string
}

func eventPlayers(players []Player) EventPlayers {
	return EventPlayers{
		Type:    EventTypePlayers,
		Players: players,
	}
}

func (p *Player) send(ctx context.Context, data []byte) error {
	if err := p.conn.Write(ctx, websocket.MessageText, data); err != nil {
		return fmt.Errorf("Error while sending %s to %s: %w", data, p.Name, err)
	}

	return nil
}

type lobby struct {
	sync.Mutex

	start   bool
	id      string
	players []Player

	rounds []Round

	close func()
}

func (l *lobby) broadcastJSON(ctx context.Context, data any) error {
	encoded, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("Failed to encode: %+v as JSON: %w", data, err)
	}

	return l.broadcast(ctx, encoded)
}

func (l *lobby) broadcast(ctx context.Context, content []byte) error {
	errs := []string{}

	for _, p := range l.players {
		if err := p.send(ctx, content); err != nil {
			errs = append(errs, err.Error())
		}
	}

	if len(errs) > 0 {
		return fmt.Errorf("Failed to broadcast to some users: %s", strings.Join(errs, "|"))
	}

	return nil
}

func (l *lobby) disconnect(ctx context.Context, player string) error {
	l.Lock()
	defer l.Unlock()

	for i, p := range l.players {
		if p.Name == player {
			l.players = append(l.players[:i], l.players[i+1:]...)

			if len(l.players) == 0 {
				// No more players let's close the lobby
				l.close()
				return nil
			}

			if p.Admin {
				// Mark as admin the first player that joined
				l.players[0].Admin = true
			}

			// Broadcast the new players
			return l.broadcastJSON(ctx, eventPlayers(l.players))
		}
	}

	return fmt.Errorf("Player %s not found while disconnecting", player)

}

func (l *lobby) isAdmin(name string) bool {
	l.Lock()
	defer l.Unlock()

	for _, p := range l.players {
		if p.Name == name {
			return p.Admin
		}
	}

	return false
}

func (l *lobby) addPlayer(ctx context.Context, name string, conn *websocket.Conn) error {
	l.Lock()
	defer l.Unlock()

	if l.start {
		return fmt.Errorf("Game has started already, cannot accept more players")
	}

	for _, p := range l.players {
		if p.Name == name {
			return fmt.Errorf("Player name %s is already taken", name)
		}
	}

	newPlayer := Player{Name: name, conn: conn, Admin: len(l.players) == 0}
	l.players = append(l.players, newPlayer)

	// Notify all the players of the new player
	return l.broadcastJSON(ctx, eventPlayers(l.players))
}

func (l *lobby) newRound(ctx context.Context) error {
	l.Lock()
	defer l.Unlock()

	letter := 'a'
	countries, err := countriesStartingWith(byte(letter))
	if err != nil {
		return err
	}

	round := Round{
		Round:              len(l.rounds) + 1,
		Letter:             string(letter), // TODO
        CurrentPlayerIdx: 0, // TODO; random
		Players:            l.players,
		Guesses:            []Guess{},
		Remaining:          len(countries),
		remainingCountries: countries,
	}

	l.rounds = append(l.rounds, round)

	return l.broadcastJSON(ctx, EventRound{
		Type:  EventTypeNewRound,
		Round: round,
	})
}

func (l *lobby) handle(ctx context.Context, from string, clientEvent ClientEvent) error {
	log.Printf("Received client event: %+v from %s", clientEvent, from)

	switch clientEvent.Type {
	case ClientEventTypeStartGame:
		if l.isAdmin(from) {
			if err := l.newRound(ctx); err != nil {
				return fmt.Errorf("failed to create new round: %w", err)
			}
		} else {
			log.Printf("Ignoring start game from %s who's not the admin", from)
		}
	}

	return nil
}

type lobbyManager struct {
	sync.Mutex
	instances map[string]*lobby
}

func (lm *lobbyManager) nextGameID() string {
	id := ""

	// find first unused ID
	for {
		if _, ok := lm.instances[id]; !ok && id != "" {
			return id
		}

		raw := make([]byte, 6)
		if _, err := rand.Read(raw); err != nil {
			panic(err)
		}
		id = fmt.Sprintf("%X", raw)
	}
}

func (lm *lobbyManager) create() string {
	lm.Lock()
	defer lm.Unlock()

	id := lm.nextGameID()
	lm.instances[id] = &lobby{id: id, close: func() {
		lm.Lock()
		defer lm.Unlock()

		delete(lm.instances, id)
		log.Printf("Closed lobby instance: %s", id)
	}}

	return id
}

func (lm *lobbyManager) get(id string) (*lobby, bool) {
	lm.Lock()
	defer lm.Unlock()

	lobby, ok := lm.instances[id]
	return lobby, ok
}

func (lm *lobbyManager) join(
	ctx context.Context,
	w http.ResponseWriter,
	lobbyID string,
	playerName string,
	conn *websocket.Conn) {

	lobby, found := lm.get(lobbyID)
	if !found {
		// Not internal technically but...
		errOutWS(conn, fmt.Sprintf("No lobby exist with id: %s", lobbyID), websocket.StatusInternalError)
		return
	}

	defer conn.Close(websocket.StatusNormalClosure, "")

	// Technically there could be a race condition changing player here if it becomes admin for instance.
	err := lobby.addPlayer(ctx, playerName, conn)
	if err != nil {
		errOutWS(conn, fmt.Sprintf("No lobby exist with id: %s", lobbyID), websocket.StatusInternalError)
		return
	}

	defer func() {
		if err := lobby.disconnect(context.Background(), playerName); err != nil {
			log.Printf("Failed to notify players from %s disconnecting: %s", playerName, err)
		}
	}()

	for {
		select {
		case <-ctx.Done():
			log.Printf("Player %s disconnected", playerName)
		default:
			_, data, err := conn.Read(ctx)
			if err != nil {
				var status websocket.CloseError
				if errors.As(err, &status) && (status.Code == websocket.StatusGoingAway || status.Code == websocket.StatusNormalClosure) {
					// Expected just return
					log.Printf("%s disconnected status: %+v", playerName, status)
					return
				}

				log.Printf("[ERROR] Reading from websocket for %s: %s", playerName, err)
				return
			}

			clientEvent := ClientEvent{}
			if err := json.Unmarshal(data, &clientEvent); err != nil {
				log.Printf("[ERROR] Failed to decode client event: %s: %s", data, err)
				return
			}

			if err := lobby.handle(ctx, playerName, clientEvent); err != nil {
				log.Printf("[ERROR] Failed to handle client event: %s: %s", data, err)
				return
			}
		}
	}
}

func main() {
	srv := http.NewServeMux()
	api := http.NewServeMux()

	statics := os.Args[1]
	listen := ":8080"
	if len(os.Args) == 3 {
		listen = os.Args[2]
	}

	fs := http.FileServer(http.Dir(statics))

	// https://stackoverflow.com/a/64687181, routing to SPA
	srv.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// If the requested file exists then return if; otherwise return index.html (fileserver default page)
		if r.URL.Path != "/" {
			fullPath := statics + strings.TrimPrefix(path.Clean(r.URL.Path), "/")
			_, err := os.Stat(fullPath)
			if err != nil {
				if !os.IsNotExist(err) {
					errOut(w, "File not found", 404)
				}
				// Requested file does not exist so we return the default (resolves to index.html)
				r.URL.Path = "/"
			}
		}
		fs.ServeHTTP(w, r)
	})

	srv.Handle("/api/", http.StripPrefix("/api", api))

	manager := &lobbyManager{instances: make(map[string]*lobby)}

	api.HandleFunc("POST /new-lobby", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")

		id := manager.create()
		log.Printf("Creating new lobby, id: %s", id)

		if err := json.NewEncoder(w).Encode(map[string]string{"id": id}); err != nil {
			errOut(w, err.Error(), 500)
			return
		}
	})

	api.HandleFunc("/lobby/{id}/{name}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")

		lobbyID := r.PathValue("id")
		playerName := r.PathValue("name")
		log.Printf("Player [%s] connecting to lobby with ID: %s", playerName, lobbyID)

		c, err := websocket.Accept(w, r, &websocket.AcceptOptions{InsecureSkipVerify: true})

		if err != nil {
			errOut(w, fmt.Sprintf("error accepting websocket: %s", err), 500)
			return
		}
		defer c.CloseNow()

		manager.join(r.Context(), w, lobbyID, playerName, c)
	})

	if err := http.ListenAndServe(listen, srv); err != nil {
		panic(err)
	}
}
