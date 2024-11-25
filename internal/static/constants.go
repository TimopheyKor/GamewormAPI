package static

import (
	"errors"
)

// Define global constants:
// TODO: Rename unexported constants to prefix with _
// TODO: Replace Schema constants with a Schema file
const (
	// Google Sheet Temp Schema:
	DbSpreadsheetName = "GamewormDataDoNotEdit"
	ReadRange         = "Test Sheet!A1:C"
	GameD             = "Games"
	GameRange         = "A1:E"
	ReviewD           = "Reviews"
	ReviewRange       = "A1:C"
	BacklogD          = "Backlog"
	BacklogRange      = "A1:E"
	maxColIdx         = 6
	// TODO: Move this path to a more secure location when hosting.
	ClientConfigPath = "C:/data/temp_credentials/credentials.json"
)

// Define columns as enum constants:
// Game table:
const (
	GamePK = iota
	GTitle
	GImage
	GDev
	GPub
)

// Define column names:
// TODO: Replace Schema vars with a Schema file
var (
	GameCols = [][]any{{"Unique Game ID", "Game Title", "Game Image", "Game Developer",
		"Game Publisher"}}
	ReviewCols  = [][]any{{"Unique Game ID", "Steam Review", "Steam Rating", "User Review", "User Rating"}}
	BacklogCols = [][]any{{"Game ID", "Priority", "Steam Flag"}}
)

// Define global errors:
var (
	ErrTooManyMatches  = errors.New("too many matches found")
	ErrNoMatchesFound  = errors.New("no matches found")
	ErrInputOutOfRange = errors.New("input length exceeds schema definition")
	ErrInputEmpty      = errors.New("empty input")
	ErrDuplicateGameID = errors.New("game id already exists in games table")
)
