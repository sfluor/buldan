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

func (p *Player) event(typ EventType) Event {
	return Event{
		Type:   typ,
		Player: *p,
	}
}

type EventType string

const (
	EventTypePlayers    = "players"
	EventTypeWrongGuess = "wrong-guess"
	EventTypeValidGuess = "valid-guess"
	EventTypeLoose      = "loose"
	EventTypeNewRound   = "new-round"
	EventTypeEndRound   = "end-round"
	EventTypeEnd        = "end"
)

type EventPlayers struct {
	Type    EventType
	Players []Player
}

func eventPlayers(players []Player) EventPlayers {
	return EventPlayers{
		Type:    EventTypePlayers,
		Players: players,
	}
}

type Event struct {
	Type   EventType
	Player Player
	data   string
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
	close   func()
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
	if err := lobby.addPlayer(ctx, playerName, conn); err != nil {
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
			typ, data, err := conn.Read(ctx)
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
			log.Printf("Received %s of type %s from %s", data, typ, playerName)
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
