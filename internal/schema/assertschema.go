package schema

import "strings"

// GenerateGameID takes two strings (game title and developer) then returns
// the corresponding generated ID.
func GenerateGameID(title, dev string) string {
	// Create an ID from the required parameters:
	adjTitle, adjDev := TrimExtendStr(title), TrimExtendStr(dev)
	idx := len(adjTitle) - 3
	genId := adjTitle[:3] + adjTitle[idx:] + adjDev[:3]
	return genId
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
