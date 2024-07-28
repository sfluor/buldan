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
	"nhooyr.io/websocket/wsjson"
)

func errOut(w http.ResponseWriter, msg string, code int) {
	log.Printf("[ERROR] Code:%d | %s", code, msg)
	http.Error(w, msg, code)
}

type lobby struct {
	sync.Mutex

	id      string
	players []string
}

func (l *lobby) addPlayer(player string) error {
	l.Lock()
	defer l.Unlock()

	for _, p := range l.players {
		if p == player {
			return fmt.Errorf("Player name %s is already taken", player)
		}
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

	api.HandleFunc("/lobby/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")

		lobbyID := r.PathValue("id")
		log.Printf("Connecting to lobby with ID: %s", lobbyID)

		c, err := websocket.Accept(w, r, &websocket.AcceptOptions{
			InsecureSkipVerify: true,
		})
		if err != nil {
			errOut(w, fmt.Sprintf("error accepting websocket: %s", err), 500)
			return
		}
		defer c.CloseNow()


		lobby, found := manager.get(lobbyID)
		if !found {
			errOut(w, fmt.Sprintf("No lobby exist with id: %s", lobbyID), 404)
			return
		}

		ctx, cancel := context.WithTimeout(r.Context(), time.Second*10)
		defer cancel()

		var v interface{}
		err = wsjson.Read(ctx, c, &v)
		if err != nil {
			errOut(w, fmt.Sprintf("error decoding first message: %s", err), 400)
			return
		}

		log.Printf("received: %v", v)

		if err := lobby.addPlayer(fmt.Sprintf("%+v", v)); err != nil {
			errOut(w, err.Error(), 400)
			return
		}

		c.Close(websocket.StatusNormalClosure, "")
	})

	if err := http.ListenAndServe(":8080", srv); err != nil {
		panic(err)
	}
}
