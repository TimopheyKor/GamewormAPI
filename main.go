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

	// Setup config and get a client with that config:
	// If modifying these scopes, delete your previously saved token.json.
	b, err := os.ReadFile(clientConfigPath)
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}
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

	// Create a Google Sheets Service from the Client:
	sSrvc, err := sheets.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		log.Fatalf("Unable to retrieve Sheets client: %v", err)
	}

	// Search for the Google Sheet we want:
	sheetId, err := getExistingSheetId(dSrvc, searchSheetName)
	if err != nil {
		log.Fatalf("getExistingSheetId error: %v", err)
	}

	// Read the test spreadsheet that was found:
	resp, err := sSrvc.Spreadsheets.Values.Get(sheetId, readRange).Do()
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
