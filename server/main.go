package main

import (
	"context"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"

	"nhooyr.io/websocket"
	// "nhooyr.io/websocket/wsjson"
)

func errOut(w http.ResponseWriter, msg string, code int) {
	log.Printf("[ERROR] Code:%d | %s", code, msg)
	http.Error(w, msg, code)
}

type player struct {
	name string
	conn *websocket.Conn
}

type EventType string

const (
	EventTypeNewPlayer  = "new-player"
	EventTypeWrongGuess = "wrong-guess"
	EventTypeValidGuess = "valid-guess"
	EventTypeLoose      = "loose"
	EventTypeNewRound   = "new-round"
    EventTypeEndRound = "end-round"
	EventTypeEnd        = "end"
)

type Event struct {
	Type   EventType
	Player string
	data   string
}

func (p *player) send(ctx context.Context, data []byte) error {
	if err := p.conn.Write(ctx, websocket.MessageText, data); err != nil {
		return fmt.Errorf("Error while sending %s to %s: %w", data, p.name, err)
	}

	return nil
}

type lobby struct {
	sync.Mutex

	start   bool
	id      string
	players []player
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

func (l *lobby) addPlayer(ctx context.Context, name string, conn *websocket.Conn) error {
	l.Lock()
	defer l.Unlock()

	if l.start {
		return fmt.Errorf("Game has started already, cannot accept more players")
	}

	for _, p := range l.players {
		if p.name == name {
			return fmt.Errorf("Player name %s is already taken", name)
		}
	}

	newPlayer := player{name: name, conn: conn}

	for _, p := range l.players {

		event := Event{Type: EventTypeNewPlayer, Player: p.name}
		encoded, err := json.Marshal(event)
		if err != nil {
			return fmt.Errorf("failed toe encode new player event: %w", err)
		}

		if err := newPlayer.send(ctx, encoded); err != nil {
			return err
		}
	}

	l.players = append(l.players, newPlayer)

	// Notify all the players of the new player
	return l.broadcastJSON(ctx, Event{Type: EventTypeNewPlayer, Player: newPlayer.name})
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
	lm.instances[id] = &lobby{id: id}

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
	defer conn.Close(websocket.StatusNormalClosure, "")

	lobby, found := lm.get(lobbyID)
	if !found {
		errOut(w, fmt.Sprintf("No lobby exist with id: %s", lobbyID), 404)
		return
	}

	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()

	if err := lobby.addPlayer(ctx, playerName, conn); err != nil {
		errOut(w, err.Error(), 400)
		return
	}

	for range time.NewTicker(time.Minute).C {
		log.Printf("Player: %s ping", playerName)
	}
}

func main() {
	srv := http.NewServeMux()
	api := http.NewServeMux()

	statics := os.Args[1]

	srv.Handle("/", http.FileServer(http.Dir(statics)))

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

	if err := http.ListenAndServe(":8080", srv); err != nil {
		panic(err)
	}
}
