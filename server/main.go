package main

import (
	"context"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
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

func (p *player) notifyNewPlayer(ctx context.Context, name string) error {
	encoded, err := json.Marshal(map[string]string{"new-player": name})
	if err != nil {
		return fmt.Errorf("(marshaling) Error when notifying %s of new player %s", p.name, name)
	}

	err = p.conn.Write(ctx, websocket.MessageText, encoded)
	if err != nil {
		return fmt.Errorf("(writing) Error when notifying %s of new player %s", p.name, name)
	}

	return nil
}

type lobby struct {
	sync.Mutex

	start   bool
	id      string
	players []player
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

	player := player{name: name, conn: conn}
	for _, p := range l.players {
		if err := p.notifyNewPlayer(ctx, name); err != nil {
			return err
		}

		if err := player.notifyNewPlayer(ctx, p.name); err != nil {
			return err
		}
	}

    // Notify the player itself
		if err := player.notifyNewPlayer(ctx, player.name); err != nil {
			return err
		}

	l.players = append(l.players, player)


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
