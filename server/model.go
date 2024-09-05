package main

import (
	"context"
	"fmt"

	"nhooyr.io/websocket"
)

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
	Language           Language
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
	Round     int
	MaxRounds int

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

type EventEnd struct {
	Type    EventType
	Players []*Player
}

type EventRound struct {
	Type    EventType
	Round   *Round
	Players []*Player
}

type EventRoundEnd struct {
	Round       int
	MaxRounds   int
	IsLastRound bool
	Type        EventType
	Letter      string
	Countries   []CountryStatus
	Guesses     []Guess
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



