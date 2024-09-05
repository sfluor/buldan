package main

import (
	"context"
	"crypto/rand"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"sync"
	"time"

	"go.uber.org/zap"
	"nhooyr.io/websocket"
)


type lobbyManager struct {
	sync.Mutex
	lg        *zap.Logger
	instances map[string]*lobby
}

func (lm *lobbyManager) nextGameID() string {
	id := ""

	// find first unused ID
	for {
		if _, ok := lm.instances[id]; !ok && id != "" {
			return id
		}

		raw := make([]byte, 3)
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

	done := make(chan bool)

	lobby := &lobby{id: id,
		state: lobbyStateWaitRoom,
		lg:    lm.lg.With(zap.String(lobbyIDKey, id)),
		close: func() {
			lm.Lock()
			defer lm.Unlock()
			done <- true

			delete(lm.instances, id)
			lm.lg.Info("Closed lobby instance", zap.String(lobbyIDKey, id))
		}}

	// Async worker to notify users when waiting for an answer / at the end of rounds.
	go func() {
		period := 1
		ticker := time.Tick(time.Duration(period) * time.Second)
		for {
			select {
			case <-ticker:
				if err := lobby.maybeBroadcastTick(context.Background(), period); err != nil {
					lm.lg.Error("Failed to broadcast tick", zap.Error(err))
				}
			case <-done:
				lm.lg.Info("Exiting ticker", zap.String(lobbyIDKey, id))
				return
			}
		}
	}()

	lm.instances[id] = lobby
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

	logFields := []zap.Field{
		zap.String(lobbyIDKey, lobbyID),
		zap.String(playerKey, playerName),
	}

	lobby, found := lm.get(lobbyID)
	if !found {
		// Not internal technically but...
		errOutWS(lm.lg, conn, fmt.Sprintf("No lobby exist with id: %s", lobbyID), websocket.StatusInternalError)
		return
	}

	defer conn.Close(websocket.StatusNormalClosure, "")

	// Technically there could be a race condition changing player here if it becomes admin for instance.
	err := lobby.addPlayer(ctx, playerName, conn)
	if err != nil {
		errOutWS(lm.lg, conn, fmt.Sprintf("Couldn't join lobby: %s", err), websocket.StatusInternalError)
		return
	}

	defer func() {
		if err := lobby.disconnect(context.Background(), playerName); err != nil {
			lm.lg.Error("Failed to notify for player disconnection", append([]zap.Field{zap.Error(err)}, logFields...)...)
		}
	}()

	for {
		select {
		case <-ctx.Done():
			lm.lg.Info("Player disconnection", zap.String(playerKey, playerName), zap.String(lobbyIDKey, lobbyID))
		default:
			_, data, err := conn.Read(ctx)
			if err != nil {
				var status websocket.CloseError
				if errors.As(err, &status) && (status.Code == websocket.StatusGoingAway || status.Code == websocket.StatusNormalClosure) {
					// Expected just return
					lm.lg.Info("Websocket disconnection status", append([]zap.Field{zap.Any("status", status)}, logFields...)...)

					return
				}

				lm.lg.Error("Failed to read from websocket", append([]zap.Field{zap.Error(err)}, logFields...)...)
				return
			}

			clientEvent := ClientEvent{}
			if err := json.Unmarshal(data, &clientEvent); err != nil {
				lm.lg.Error("Failed to decode payload", append([]zap.Field{zap.Error(err), zap.String("payload", string(data))}, logFields...)...)
				errOutWS(lm.lg, conn, fmt.Sprintf("Couldn't decode client event: %s", err), websocket.StatusInternalError)
				return
			}

			if err := lobby.handle(ctx, playerName, clientEvent); err != nil {
				lm.lg.Error("Failed to handle payload", append([]zap.Field{zap.Error(err), zap.String("payload", string(data))}, logFields...)...)
				errOutWS(lm.lg, conn, fmt.Sprintf("Couldn't handle client event: %s", err), websocket.StatusInternalError)
				return
			}
		}
	}
}
