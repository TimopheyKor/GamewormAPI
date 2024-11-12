package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/TimopheyKor/GamewormAPI/sheetfuncs"
	"github.com/TimopheyKor/GamewormAPI/static"
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
	b, err := os.ReadFile(static.ClientConfigPath)
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}
	config, err := google.ConfigFromJSON(b, sheets.SpreadsheetsReadonlyScope, sheets.DriveFileScope, drive.DriveReadonlyScope)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}
	client := getClient(ctx, config)

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

	// Search for the database sheet. If the sheet isn't found, return
	// the appropriate error, or create the databse sheet.
	sheetId, err := getExistingSheetId(dSrvc, static.DbSpreadsheetName)
	if err != nil && !errors.Is(err, static.ErrNoMatchesFound) {
		log.Fatalf("get existing sheet id: %s", err)
	} else if errors.Is(err, static.ErrNoMatchesFound) {
		fmt.Printf("%s, creating new sheet for db\n", err)
		sheetId, err = sheetfuncs.NewSheetDB(ctx, sSrvc)
		if err != nil {
			log.Fatalf("failed to create new sheet db: %v", err)
		}
	} else {
		fmt.Printf("existing sheet found: %s\n", static.DbSpreadsheetName)
	}

	// Read the first sheet of the sheetDB:
	resp, err := sSrvc.Spreadsheets.Values.Get(sheetId, static.GameD+"!"+static.GameRange).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve data from sheet: %v", err)
	}

	if len(resp.Values) == 0 {
		fmt.Println("No data found.")
	} else {
		fmt.Printf("Example sheet data found:\n%+v\n", resp.Values)
		// for _, row := range resp.Values {
		// 	// Print columns A and E, which correspond to indices 0 and 4.
		// 	fmt.Printf("%s, %s\n", row[0], row[4])
		// }
	}
}
