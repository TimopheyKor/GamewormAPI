package main

import (
	"fmt"

	"google.golang.org/api/drive/v3"
)

// If a sheet for Gameworm's data was already created by the user, find and
// return its Id. If not, return an empty string and proper error.
func getExistingSheetId(srv *drive.Service, name string) (string, error) {
	// Query the root layer of user's Drive for db file:
	res, err := srv.Files.List().
		Q("name=\"" + name + "\" and trashed=false").Fields("files(id,parents)").
		Do()
	if err != nil {
		return "", fmt.Errorf("query files list for %s: %w", name, err)
	}

	// Handle errors & unexpected results:
	if len(res.Files) > 1 {
		return "", fmt.Errorf("%s: %w", name, ErrTooManyMatches)
	} else if len(res.Files) == 0 || len(res.Files[0].Id) == 0 {
		return "", fmt.Errorf("%s: %w", name, ErrNoMatchesFound)
	}
	return res.Files[0].Id, nil
}

// TODO: Add backup functionality for searching through other filesystems
// and/or finding more data files if needed.
