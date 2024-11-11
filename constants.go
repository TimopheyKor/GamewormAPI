package main

import (
	"errors"
)

// Define global constants:
// TODO: Rename unexported constants to prefix with _
// TODO: Replace Schema constants with a Schema file
const (
	// Google Sheet Temp Schema:
	dbSpreadsheetName = "GamewormDataDoNotEdit"
	readRange         = "Test Sheet!A1:C"
	gameD             = "Games"
	gameRange         = "A1:E"
	reviewD           = "Reviews"
	reviewRange       = "A1:C"
	backlogD          = "Backlog"
	backlogRange      = "A1:E"
	// TODO: Move this path to a more secure location when hosting.
	clientConfigPath = "C:/data/temp_credentials/credentials.json"
)

// Define column names:
// TODO: Replace Schema vars with a Schema file
var (
	gameCols = [][]any{{"Unique Game ID", "Game Title", "Game Image", "Game Developer",
		"Game Publisher"}}
	reviewCols  = [][]any{{"Unique Game ID", "Steam Review", "Steam Rating", "User Review", "User Rating"}}
	backlogCols = [][]any{{"Game ID", "Priority", "Steam Flag"}}
)

// Define global errors:
var (
	ErrTooManyMatches = errors.New("too many matches found")
	ErrNoMatchesFound = errors.New("no matches found")
)
