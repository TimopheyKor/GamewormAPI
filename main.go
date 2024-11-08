package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/drive/v3"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

// Retrieve a token, saves the token, then returns the generated client.
func getClient(config *oauth2.Config) *http.Client {
	// The file token.json stores the user's access and refresh tokens, and is
	// created automatically when the authorization flow completes for the first
	// time.
	tokFile := "token.json"
	tok, err := tokenFromFile(tokFile)
	if err != nil {
		tok = getTokenFromWeb(config)
		saveToken(tokFile, tok)
	}
	return config.Client(context.Background(), tok)
}

// Request a token from the web, then returns the retrieved token.
func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the "+
		"authorization code: \n%v\n", authURL)

	var authCode string
	if _, err := fmt.Scan(&authCode); err != nil {
		log.Fatalf("Unable to read authorization code: %v", err)
	}

	tok, err := config.Exchange(context.TODO(), authCode)
	if err != nil {
		log.Fatalf("Unable to retrieve token from web: %v", err)
	}
	return tok
}

// Retrieves a token from a local file.
func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	return tok, err
}

// Saves a token to a file path.
func saveToken(path string, token *oauth2.Token) {
	fmt.Printf("Saving credential file to: %s\n", path)
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatalf("Unable to cache oauth token: %v", err)
	}
	defer f.Close()
	json.NewEncoder(f).Encode(token)
}

// TODO: Create functionality to send requests to Google Docs
func main() {
	// Flag usage example:
	lvlFlag := flag.String("l", "info", "description of log level")
	flag.Parse()
	fmt.Printf("level is %s\n", *lvlFlag)

	// Google Cloud OAuth Process:
	fmt.Println("Processing Google Cloud OAuth...")
	ctx := context.Background()
	// TODO: When pushing to server, make this an encironment variable as direct filepaths are unsafe.
	b, err := os.ReadFile("C:/data/temp_credentials/credentials.json")
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

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
	searchfile := "LocalAPITestSheet"
	res, err := dSrvc.Files.List().
		Q("name=\"" + searchfile + "\" and trashed=false").
		Fields("files(id,parents)").Do()
	if err != nil {
		log.Fatalf("Q.Fields.Do() call error: %v", err)
	}
	if len(res.Files) > 1 {
		log.Fatalf("Multiple matching filenames found, please delete"+
			"or rename incorrect files named: %v", searchfile)
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
