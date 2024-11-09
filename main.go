package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/drive/v3"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

func main() {
	// Google Cloud OAuth Process:
	fmt.Println("Processing Google Cloud OAuth...")
	ctx := context.Background()
	// TODO: When pushing to server, make this an encironment variable as direct filepaths are unsafe.
	b, err := os.ReadFile(clientConfigPath)
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	// Setupi config and get a client with that config:
	// If modifying these scopes, delete your previously saved token.json.
	config, err := google.ConfigFromJSON(b, sheets.SpreadsheetsReadonlyScope, sheets.DriveFileScope, drive.DriveReadonlyScope)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}
	client := getClient(config)

	// Creating a Google Drive service from the client:
	dSrvc, err := drive.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		log.Fatalf("new drive: %s", err)
	}

	// Search for the Google Sheet we want:
	// TODO: Make searchfile a constant string in a config file somewhere.
	res, err := dSrvc.Files.List().
		Q("name=\"" + searchSheetName + "\" and trashed=false").
		Fields("files(id,parents)").Do()
	if err != nil {
		log.Fatalf("Q.Fields.Do() call error: %v", err)
	}
	if len(res.Files) > 1 {
		log.Fatalf("Multiple matching filenames found, please delete"+
			"or rename incorrect files named: %v", searchSheetName)
	}

	// Download file content:
	fileId := res.Files[0].Id
	file, err := dSrvc.Files.Get(fileId).Do()
	if err != nil {
		log.Fatalf("Unable to download file on Do() call: %v", err)
	}
	fmt.Println("File Name Found:", file.Name)

	// Create a Google Sheets Service from the Client:
	sSrvc, err := sheets.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		log.Fatalf("Unable to retrieve Sheets client: %v", err)
	}

	// Read the test spreadsheet that was found:
	readRange := "Test Sheet!A1:C"
	resp, err := sSrvc.Spreadsheets.Values.Get(fileId, readRange).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve data from sheet: %v", err)
	}

	if len(resp.Values) == 0 {
		fmt.Println("No data found.")
	} else {
		fmt.Println("Example sheet data found.")
		// for _, row := range resp.Values {
		// 	// Print columns A and E, which correspond to indices 0 and 4.
		// 	fmt.Printf("%s, %s\n", row[0], row[4])
		// }
	}
}
