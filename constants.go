package main

import (
	"errors"
)

// Define global constants:
// TODO: Rename unexported constants to prefix with _
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

// Define global errors:
var (
	ErrTooManyMatches = errors.New("too many matches found")
)
