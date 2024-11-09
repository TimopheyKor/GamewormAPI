package main

import (
	"fmt"
)

// Define global constants:
const (
	// Google Sheet Name to search for in user's Google Drive:
	searchSheetName = "LocalAPITestSheet"
	readRange       = "Test Sheet!A1:C"
	// TODO: Move this path to a more secure location when hosting.
	clientConfigPath = "C:/data/temp_credentials/credentials.json"
)

// Define global errors:
var (
	ErrTooManyMatchingSheets = fmt.Errorf(
		"multiple files found matching sheet query for filename %v",
		searchSheetName)
)
