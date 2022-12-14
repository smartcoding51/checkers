package types

import "time"

const (
	// ModuleName defines the module name
	ModuleName = "checkers"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_checkers"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

const (
	NextGameKey = "NextGame-value-"
)

const (
	StoredGameEventKey     = "NewGameCreated"
	StoredGameEventCreator = "Creator"
	StoredGameEventIndex   = "Index"
	StoredGameEventRed     = "Red"
	StoredGameEventBlack   = "Black"
)

const (
	PlayMoveEventKey       = "MovePlayed"
	PlayMoveEventCreator   = "Creator"
	PlayMoveEventIdValue   = "IdValue"
	PlayMoveEventCapturedX = "CapturedX"
	PlayMoveEventCapturedY = "CapturedY"
	PlayMoveEventWinner    = "Winner"
	MovePlayedEventBoard   = "board"
)

const (
	RejectGameEventKey     = "GameRejected"
	RejectGameEventCreator = "Creator"
	RejectGameEventIdValue = "IdValue"
)

const (
	NoFifoIndex = "-1"
)

const (
	// MaxTurnDuration = time.Duration(24 * 3_600 * 1000_000_000) // 1 day
	// Temporarily set to 5 mins
	MaxTurnDuration = time.Duration(5 * 60 * 1000_000_000) // 5 minutes
	DeadlineLayout  = "2006-01-02 15:04:05.999999999 +0000 UTC"
)

const (
	GameForfeitedEventType      = "game-forfeited"
	GameForfeitedEventGameIndex = "game-index"
	GameForfeitedEventWinner    = "winner"
	GameForfeitedEventBoard     = "board"
)
