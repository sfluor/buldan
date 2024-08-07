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
	"time"

	"nhooyr.io/websocket"
)

// TODO better logging

const (
	hardDisconnectDelay = 60 * time.Second
	newRoundDelaySec    = 10
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

// TODO: add symbol for player
type Player struct {
	Name   string
	Admin  bool
	Points uint32
	// Technically derived from conn below
	Connected bool
	// Internal
	conn *websocket.Conn
}

type GameOptions struct {
	Rounds             int
	MaxGuessesPerRound int
	GuessTimeSeconds   int
}

type EventType string

const (
	EventTypePlayers     = "players"
	EventTypeGuess       = "guess"
	EventTypeLoose       = "loose"
	EventTypeRoundUpdate = "round-update"
	EventTypeTick        = "tick"
	EventTypeEndRound    = "end-round"
	EventTypeEnd         = "end"
)

type ClientEventType string

const (
	ClientEventTypeGuess     = "guess"
	ClientEventTypeStartGame = "start-game"
)

type PlayerStatus struct {
	RemainingGuesses int
}

type Round struct {
	Round int

	PlayersStatuses map[string]*PlayerStatus

	CurrentPlayerIndex int

	currentPlayerRemainingSec int

	Letter string // string to ease the conversion to ASCII in JSON

	Guesses []Guess

	// Duplicate it since this one isn't serialized even
	// though we could infer it from the one below
	remainingCountries Countries
	Remaining          int
}

type EventTick struct {
	Type         EventType
	RemainingSec int
}

type EventRound struct {
	Type    EventType
	Round   *Round
	Players []*Player
}

type EventRoundEnd struct {
	Type      EventType
	Letter    string
	Countries []CountryStatus
	Guesses   []Guess
}

type Guess struct {
	Guess   string
	Player  string
	Flag    string
	Correct bool
}

type ClientEvent struct {
	Type    ClientEventType
	Options GameOptions
	Guess   string
}

func (p *Player) send(ctx context.Context, data []byte) error {
	if err := p.conn.Write(ctx, websocket.MessageText, data); err != nil {
		return fmt.Errorf("Error while sending %s to %s: %w", data, p.Name, err)
	}

	return nil
}

type lobbyState string

const (
	lobbyStateWaitRoom      = "wait-room"
	lobbyStateRound         = "round"
	lobbyStateBetweenRounds = "between-rounds"
)

type lobby struct {
	sync.Mutex

	started bool
	id      string
	players []*Player

	state lobbyState

	rounds []*Round

	letters []byte

	opts GameOptions

	close func()
}

func (l *lobby) broadcastJSON(ctx context.Context, data any) error {
	encoded, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("Failed to encode: %+v as JSON: %w", data, err)
	}

	log.Printf("[lobby:%s] Broadcasting: %+v", l.id, data)
	return l.broadcast(ctx, encoded)
}

func (l *lobby) maybeBroadcastTick(ctx context.Context, period int) error {
	// TODO
	l.Lock()
	defer l.Unlock()

	switch l.state {
	case lobbyStateWaitRoom:
		// Nothing to tick in the wait room
		return nil
	case lobbyStateRound:
		round := l.currentRound()
		if round == nil {
			return fmt.Errorf("Expected round to exist in state: %s", lobbyStateRound)
		}

		round.currentPlayerRemainingSec -= period
		if round.currentPlayerRemainingSec <= 0 {
			player := l.players[round.CurrentPlayerIndex]
			round.PlayersStatuses[player.Name].RemainingGuesses = 0
			return l.nextPlayerAndMaybeNewRound(ctx, round)
		}

		log.Printf("ticking from %s: %s (remaining: %d)", l.state, l.id, round.currentPlayerRemainingSec)
		return l.broadcastJSON(ctx, EventTick{
			Type:         EventTypeTick,
			RemainingSec: round.currentPlayerRemainingSec,
		})
	case lobbyStateBetweenRounds:
		return fmt.Errorf("not expected to tick between rounds")
	}

	return nil
}

func (l *lobby) currentRound() *Round {
	if len(l.rounds) > 0 {
		return l.rounds[len(l.rounds)-1]
	}

	return nil
}

func (l *lobby) broadcastPlayers(ctx context.Context) error {
	return l.broadcastJSON(ctx, EventRound{
		Type:    EventTypePlayers,
		Players: l.players,
	})
}

func (l *lobby) broadcast(ctx context.Context, content []byte) error {
	errs := []string{}

	for _, p := range l.players {
		if p.Connected {
			if err := p.send(ctx, content); err != nil {
				errs = append(errs, err.Error())
			}
		}
	}

	if len(errs) > 0 {
		return fmt.Errorf("Failed to broadcast to some users: %s", strings.Join(errs, "|"))
	}

	return nil
}

func (l *lobby) hardDisconnect(player string) error {
	l.Lock()
	defer l.Unlock()

	for i, p := range l.players {
		if p.Name == player {
			if p.Connected {
				// The player reconnected, no need to hard disconnect them.
				return nil
			}

			l.players = append(l.players[:i], l.players[i+1:]...)

			log.Printf("hard-disconnecting player %s from %s", player, l.id)

			if len(l.players) == 0 {
				// No more players let's close the lobby
				l.close()
				log.Printf("Closing lobby %s after all players left", l.id)
				return nil
			}

			round := l.currentRound()
			wasPlaying := round != nil && round.CurrentPlayerIndex >= len(l.players)
			// Offset the player index to avoid the FE from complaining, we'll see afterwards if we
			// need to create a new round.
			if wasPlaying {
				round.CurrentPlayerIndex--
			}

			if err := l.broadcastJSON(context.Background(), EventRound{
				Type:    EventTypeRoundUpdate,
				Players: l.players,
				Round:   round,
			}); err != nil {
				return err
			}

			// TODO dedupe this with above
			// If the player was this one move to the next player
			if wasPlaying {
				return l.nextPlayerAndMaybeNewRound(context.Background(), round)
			}

			return nil
		}
	}

	return fmt.Errorf("Player %s wasn't found for hard-disconnect, this isn't expected", player)
}

func (l *lobby) disconnect(ctx context.Context, player string) error {
	l.Lock()
	defer l.Unlock()

	for _, p := range l.players {
		if p.Name == player {
			p.Connected = false
			p.conn = nil

			// Schedule a hard disconnect
			go func() {
				<-time.Tick(hardDisconnectDelay)
				l.hardDisconnect(player)
			}()

			if p.Admin {
				// Mark as admin the first connected player.
				for _, p2 := range l.players {
					if p2.Connected {
						p2.Admin = true
						p.Admin = false
					}
				}
			}

			log.Printf("Player %s disconnected", player)

			// Broadcast the new players
			return l.broadcastPlayers(ctx)
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

func (l *lobby) reconnect(ctx context.Context, player *Player, conn *websocket.Conn) error {
	player.conn = conn
	player.Connected = true

	log.Printf("Player %s reconnected", player.Name)

	// Notify all the players of the new player
	if err := l.broadcastPlayers(ctx); err != nil {
		return err
	}

	round := l.currentRound()
	if round != nil {
		encoded, err := json.Marshal(EventRound{
			Type:    EventTypeRoundUpdate,
			Round:   l.currentRound(),
			Players: l.players,
		})
		if err != nil {
			return fmt.Errorf("Failed to encode as JSON: %w", err)
		}
		return player.send(ctx, encoded)
	}

	return nil
}

func (l *lobby) addPlayer(ctx context.Context, name string, conn *websocket.Conn) error {
	l.Lock()
	defer l.Unlock()

	for _, p := range l.players {
		if p.Name == name {
			if p.Connected {
				return fmt.Errorf("Player name %s is already taken", name)
			} else {
				return l.reconnect(ctx, p, conn)
			}
		}
	}

	// Not an existing player and the game has started, nope.
	if l.started {
		return fmt.Errorf("Game has started already, cannot accept more players")
	}

	newPlayer := Player{
		Name:      name,
		conn:      conn,
		Admin:     len(l.players) == 0,
		Connected: true,
	}
	l.players = append(l.players, &newPlayer)

	// Notify all the players of the new player
	return l.broadcastPlayers(ctx)
}

func (l *lobby) newRoundUnsafe(ctx context.Context) error {
	// TODO Careful of too many rounds
	letter := l.letters[len(l.rounds)]
	countries := countriesStartingWith(byte(letter))

	currentPlayerIndex := 0
	lastRound := l.currentRound()
	if lastRound != nil {
		currentPlayerIndex = (lastRound.CurrentPlayerIndex + 1) % len(l.players)

		endRound := EventRoundEnd{
			Type:      EventTypeEndRound,
			Letter:    lastRound.Letter,
			Guesses:   lastRound.Guesses,
			Countries: lastRound.remainingCountries.status(),
		}

		l.state = lobbyStateBetweenRounds

		if err := l.broadcastJSON(ctx, endRound); err != nil {
			return err
		}

		// TODO not great since we block the lobby
		// this could be a cond unlocked by the real ticker instead
		rem := newRoundDelaySec
		for range time.Tick(time.Second) {
			if rem == 0 {
				break
			}

			rem -= 1
			if err := l.broadcastJSON(ctx, EventTick{
				RemainingSec: rem,
				Type:         EventTypeTick,
			}); err != nil {
				return err
			}
		}
	}

	l.state = lobbyStateRound

	round := Round{
		Round:                     len(l.rounds) + 1,
		Letter:                    string(letter),
		CurrentPlayerIndex:        currentPlayerIndex,
		currentPlayerRemainingSec: l.opts.GuessTimeSeconds,
		PlayersStatuses:           map[string]*PlayerStatus{},
		Guesses:                   []Guess{},
		Remaining:                 countries.remaining(),
		remainingCountries:        countries,
	}

	for _, p := range l.players {
		round.PlayersStatuses[p.Name] = &PlayerStatus{
			RemainingGuesses: l.opts.MaxGuessesPerRound,
		}
	}

	l.rounds = append(l.rounds, &round)

	return l.broadcastJSON(ctx, EventRound{
		Type:    EventTypeRoundUpdate,
		Players: l.players,
		Round:   &round,
	})
}

func (l *lobby) newRound(ctx context.Context) error {
	l.Lock()
	defer l.Unlock()
	return l.newRoundUnsafe(ctx)
}

func (l *lobby) nextPlayerAndMaybeNewRound(ctx context.Context, currRound *Round) error {
	if l.nextPlayer() {
		return l.newRoundUnsafe(ctx)
	}

	return l.broadcastJSON(ctx, EventRound{
		Type:    EventTypeRoundUpdate,
		Round:   currRound,
		Players: l.players,
	})
}

func (l *lobby) nextPlayer() (done bool) {
	round := l.currentRound()

	start := round.CurrentPlayerIndex

	for {
		round.CurrentPlayerIndex++
		if round.CurrentPlayerIndex >= len(l.players) {
			round.CurrentPlayerIndex = 0
		}

		p := l.players[round.CurrentPlayerIndex]
		// We found a player that didn't loose yet and is still connected
		if round.PlayersStatuses[p.Name].RemainingGuesses > 0 && p.Connected {
			round.currentPlayerRemainingSec = l.opts.GuessTimeSeconds
			return false
		}

		// We navigated through all the players, let's start a new round
		if round.CurrentPlayerIndex == start {
			return true
		}
	}
}

func (l *lobby) handleGuess(ctx context.Context, from string, guess string) error {
	l.Lock()
	defer l.Unlock()

	round := l.currentRound()

	expectedPlayer := l.players[round.CurrentPlayerIndex].Name
	if expectedPlayer != from {
		return fmt.Errorf("Expected %s to play but received guess %s from %s", expectedPlayer, guess, from)
	}

	guess = strings.ToLower(guess)
	country, guessStr, correct := round.remainingCountries.guess(guess, from)

	round.Guesses = append(round.Guesses, Guess{
		Player:  from,
		Guess:   guessStr,
		Flag:    country.Flag,
		Correct: correct,
	})

	if correct {
		l.players[round.CurrentPlayerIndex].Points += 1
		round.Remaining = round.remainingCountries.remaining()
		if round.Remaining == 0 {
			return l.newRoundUnsafe(ctx)
		}
		return l.nextPlayerAndMaybeNewRound(ctx, round)
	} else {
		status := round.PlayersStatuses[from]
		status.RemainingGuesses--
		if status.RemainingGuesses <= 0 {
			return l.nextPlayerAndMaybeNewRound(ctx, round)
		}
	}

	// TODO: hints
	// TODO: time

	// TODO end-round
	// TODO flag

	return l.broadcastJSON(ctx, EventRound{
		Type:    EventTypeRoundUpdate,
		Round:   round,
		Players: l.players,
	})
}

func (l *lobby) handle(ctx context.Context, from string, clientEvent ClientEvent) error {
	log.Printf("Received client event: %+v from %s", clientEvent, from)

	switch clientEvent.Type {
	case ClientEventTypeStartGame:
		if l.isAdmin(from) {
			l.started = true
			l.opts = clientEvent.Options

			// TODO respect options
			log.Printf("Starting with game options: %+v", l.opts)

			if err := l.newRound(ctx); err != nil {
				return fmt.Errorf("failed to create new round: %w", err)
			}
		} else {
			log.Printf("Ignoring start game from %s who's not the admin", from)
		}
	case ClientEventTypeGuess:
		if err := l.handleGuess(ctx, from, clientEvent.Guess); err != nil {
			return fmt.Errorf("failed to handle guess from %s: %s: %w", from, clientEvent.Guess, err)
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
		letters: newLetters(),
		state:   lobbyStateWaitRoom,
		close: func() {
			lm.Lock()
			defer lm.Unlock()
			done <- true

			delete(lm.instances, id)
			log.Printf("Closed lobby instance: %s", id)
		}}

	// Async worker to notify users when waiting for an answer / at the end of rounds.
	go func() {
		period := 1
		ticker := time.Tick(time.Duration(period) * time.Second)
		for {
			select {
			case <-ticker:
				if err := lobby.maybeBroadcastTick(context.Background(), period); err != nil {
					log.Printf("[ERROR] Failed to broadcast tick: %s", err)
				}
			case <-done:
				log.Printf("exiting ticker from %s", id)
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
		errOutWS(conn, fmt.Sprintf("Couldn't join lobby: %s", err), websocket.StatusInternalError)
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
				errOutWS(conn, fmt.Sprintf("Couldn't decode client event: %s", err), websocket.StatusInternalError)
				return
			}

			if err := lobby.handle(ctx, playerName, clientEvent); err != nil {
				log.Printf("[ERROR] Failed to handle client event: %s: %s", data, err)
				errOutWS(conn, fmt.Sprintf("Couldn't handle client event: %s", err), websocket.StatusInternalError)
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
