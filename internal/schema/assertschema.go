package schema

import "strings"

// TODO: Write a function that generates a UNIQUE GAME ID for a game.
// This should maybe go in a different package.
// Maybe convert the values input to a JSON input when making the frontend.
func NewGameData(values []string) ([]any, error) {
	// resp := []any{values}
	return nil, nil
}

// TODO: Write a function that asserts that a game update or append call comes
// with a game ID, creating an ID if one is not provided.
func CheckGameID() {}

// TODO: Fill out AssertNewGameData so it checks if incoming game data is valid,
// and returns false if it's not.
func AssertNewGameData(values []string) bool {
	return false
}

// TrimExtendStr trims all spaces from a string and extends it to length 3 if
// it is not already.
func TrimExtendStr(s string) string {
	// TODO: Remove all whitespace rather than all single spaces.
	s = strings.ReplaceAll(s, " ", "")
	for len(s) < 3 {
		s = s + s[:1]
	}
	return s
}
