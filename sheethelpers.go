package main

import (
	"context"

	"google.golang.org/api/sheets/v4"
)

// TODO: Make this function pull from a schema rather than constantcs for
// creating the DB.

// Using a context and a sheets service, creates a new spreadsheet to hold
// the Gameworm DB, then returns the spreadsheet's SheetID.
func newSheetDB(ctx context.Context, srv *sheets.Service) (string, error) {
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
	return spreadsheet.SpreadsheetId, nil
}

func prepUpdateCall(ctx context.Context, srv *sheets.Service, sheetId string) func(string, [][]any) (string, error) {
	return func(s string, a [][]any) (string, error) {
		return srv.Spreadsheets
	}
}

func updateRowCells(sRange string, values [][]any) (string, error) {
	res, err := srv.Spreadsheets
}

updateFn := prepSheetsCall(ctx, srv, id)
s, err := updateFn(range, vals)
s, err := updateFn(range2, vals2)