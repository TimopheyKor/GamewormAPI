package main

import (
	"fmt"

	"google.golang.org/api/drive/v3"
)

// If a sheet for Gameworm's data was already created by the user, find and
// return its Id. If not, return nil.
func getExistingSheetId(srv *drive.Service, name string) (string, error) {
	// Query the file on top layer of user's Drive:
	res, err := srv.Files.List().
		Q("name=\"" + name + "\" and trashed=false").Fields("files(id,parents)").
		Do()

	// Handle errors & unexpected results:
	if err != nil {
		return "", fmt.Errorf("query files list for %s: %w", name, err)
	}
	if len(res.Files) > 1 {
		return "", fmt.Errorf("%s: %w", name, ErrTooManyMatches)
	}
	return res.Files[0].Id, nil
}

/*  Old code, may re-use for debugging:
// Download file content:
fileId := res.Files[0].Id
file, err := dSrvc.Files.Get(fileId).Do()
if err != nil {
	log.Fatalf("Unable to download file on Do() call: %v", err)
}
fmt.Println("File Name Found:", file.Name)
*/
