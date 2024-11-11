package main

import (
	"context"
	"fmt"

	"google.golang.org/api/sheets/v4"
)

// TODO: Make this function pull from a schema rather than constantcs for
// creating the DB.

// Using a context and a sheets service, creates a new spreadsheet to hold
// the Gameworm DB, initializes the schema of the DB, then returns the
// spreadsheet's SheetID.
func newSheetDB(ctx context.Context, srv *sheets.Service) (string, error) {
	// Create a new spreadsheet with three sheets:
	spreadsheet, err := srv.Spreadsheets.Create(&sheets.Spreadsheet{
		Properties: &sheets.SpreadsheetProperties{
			Title: dbSpreadsheetName,
		},
		Sheets: []*sheets.Sheet{
			{
				Properties: &sheets.SheetProperties{
					Title: gameD,
				},
			},
			{
				Properties: &sheets.SheetProperties{
					Title: reviewD,
				},
			},
			{
				Properties: &sheets.SheetProperties{
					Title: backlogD,
				},
			},
		},
	}).Context(ctx).Do()
	if err != nil {
		return "", err
	}
	// Initialize the header values (column names) for the three sheets:
	spreadsheetId := spreadsheet.SpreadsheetId
	updateFn := prepUpdateCall(ctx, srv, spreadsheetId)
	 if _, err = updateFn(gameD+"!A1:E1", gameCols) != nil { return "", err }
	return spreadsheetId, nil
}

func prepUpdateCall(ctx context.Context, srv *sheets.Service, sheetId string) func(string, [][]any) (string, error) {
	return func(sRange string, values [][]any) (string, error) {
		res, err := srv.Spreadsheets.Values.Update(sheetId, sRange, &sheets.ValueRange{
			MajorDimension: "ROWS",
			Values:         values,
		}).ValueInputOption("RAW").Context(ctx).Do()
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("update response: %+v\n", res), err
	}
}

// func updateRowCells(sRange string, values [][]any) (string, error) {
// 	res, err := srv.Spreadsheets
// }

// updateFn := prepSheetsCall(ctx, srv, id)
// s, err := updateFn(range, vals)
// s, err := updateFn(range2, vals2)
