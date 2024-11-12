package sheetfuncs

import (
	"context"
	"fmt"

	"github.com/TimopheyKor/GamewormAPI/static"
	"google.golang.org/api/sheets/v4"
)

// TODO: Make this function pull from a schema rather than constantcs for
// creating the DB.
// TODO: Consider creating a "sheetHolder" type that holds the context,
// service, and sheetId (or sheet object itself if modifiable) rather than
// passing through nested functions.

// newSheetDB takes a context and a sheets service, creates a new spreadsheet to
// hold the Gameworm DB, initializes the schema of the DB, then returns the
// spreadsheet's SheetID.
func NewSheetDB(ctx context.Context, srv *sheets.Service) (string, error) {
	// Create a new spreadsheet with three sheets:
	spreadsheet, err := srv.Spreadsheets.Create(&sheets.Spreadsheet{
		Properties: &sheets.SpreadsheetProperties{
			Title: static.DbSpreadsheetName,
		},
		Sheets: []*sheets.Sheet{
			{
				Properties: &sheets.SheetProperties{
					Title: static.GameD,
				},
			},
			{
				Properties: &sheets.SheetProperties{
					Title: static.ReviewD,
				},
			},
			{
				Properties: &sheets.SheetProperties{
					Title: static.BacklogD,
				},
			},
		},
	}).Context(ctx).Do()
	if err != nil {
		return "", err
	}

	// Initialize the header values (column names) for the three sheets:
	spreadsheetId := spreadsheet.SpreadsheetId
	updateFn := prepInitUpdateCall(ctx, srv, spreadsheetId)
	res, err := updateFn(static.GameD+"!A1:E1", static.GameCols)
	if err != nil {
		return res, err
	}
	res, err = updateFn(static.ReviewD+"!A1:E1", static.ReviewCols)
	if err != nil {
		return res, err
	}
	res, err = updateFn(static.BacklogD+"!A1:C1", static.BacklogCols)
	if err != nil {
		return res, err
	}
	return spreadsheetId, nil
}

// prepInitUpdateCall takes a context, sheets service, and sheet id, and returns
// a function used to update database rows given a range and values.
func prepInitUpdateCall(ctx context.Context, srv *sheets.Service, sheetId string) func(string, [][]any) (string, error) {
	// Input: range string ("sheet_name!start_cell:end_cell"), 2D array of values
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

// prepUpdateCall takes a context, sheets service, and sheet id, and returns
// a function used to update a database row given a sheet

// TODO: Write body for prepAppendCall.
// prepAppendCall takes a context, sheets service, and sheet id, and returns
// a function used to append new rows to a database, given a sheet and values.
func prepAppendCall(ctx context.Context, srv *sheets.Service, sheetId string) func(string, []any) (string, error) {
	return func(sheetName string, values []any) (string, error) {
		return fmt.Sprint("Append call being implemented"), nil
	}
}

// TODO: Write body for prepReadCall.
// prepReadCall takes a context, sheets service, and sheet id, and returns
// a function used to read database rows given a range.
func prepReadCall(ctx context.Context, srv *sheets.Service, sheetId string) func(string) (string, error) {
	return func(sRange string) (string, error) {
		return fmt.Sprint("Read call being implemented"), nil
	}
}

// TODO: Write body for prepDeleteCall.
// prepDeleteCall takes a context, sheets service, and sheet id, and returns
// a function used to read database rows given a range.
func prepDeleteCall(ctx context.Context, srv *sheets.Service, sheetId string) func(string) (string, error) {
	return func(sRange string) (string, error) {
		return fmt.Sprint("Delete call being implemented"), nil
	}
}
