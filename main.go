package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"

	"net/url"

	"github.com/TimopheyKor/GamewormAPI/internal/schema"
	"github.com/TimopheyKor/GamewormAPI/internal/sheetfuncs"
	"github.com/TimopheyKor/GamewormAPI/internal/static"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/drive/v3"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

func main() {
	// Google Cloud OAuth Process:
	fmt.Println("Testing from alternate machine...")
	// TODO: Auth code sometimes encodes as URL and needs to be decoded before working, figure out how
	// to automate this:
	encodedString := ""
	decodedString, err := url.QueryUnescape(encodedString)
	fmt.Printf("Auth Code Decode: %v\n", decodedString)

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

	// Testing GameIdExists function and SheetsHolder:
	testSheetsHolder := sheetfuncs.NewSheetsHolder(ctx, sSrvc, sheetId)
	val, err := testSheetsHolder.GameIdExists("BOB", static.GameD)
	if err != nil {
		log.Fatalf("unable to check for ID: %v\n", err)
	}
	fmt.Printf("GameIdExists(BOB, GameD): %v\n", val)
	val, err = testSheetsHolder.GameIdExists("TESME1FDEVFPUB", static.GameD)
	if err != nil {
		log.Fatalf("unable to check for ID: %v\n", err)
	}
	fmt.Printf("GameIdExists(TESME1FDEVFPUB, GameD): %v\n", val)

	// Testing AddNewGame function
	exGameInfo := schema.NewGameObject(
		"TEST GAME 1",
		"FAKE DEV",
		schema.WithImg("www.exampleimage.com/image.jpg"),
		schema.WithPub("FAKE PUBLISHER"),
	)

	res, err := testSheetsHolder.AddNewGame(exGameInfo)
	if err != nil {
		log.Fatalf("unable to append data to sheet: %v\n", err)
	}
	fmt.Printf("append response: %+v\n", res)
}
