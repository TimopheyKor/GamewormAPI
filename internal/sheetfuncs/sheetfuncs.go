package sheetfuncs

import (
	"context"
	"fmt"

	"github.com/TimopheyKor/GamewormAPI/internal/static"
	"google.golang.org/api/sheets/v4"
)

// TODO: Make this function pull from a schema rather than constantcs for
// creating the DB.
// TODO: Make all the sheet-specific functions work off of the SheetWorker rather
// then a wrapper function.

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

// TODO: Convert prepAppendCall to AppendToSheet as a method on SheetWorker.
// prepAppendCall takes a context, sheets service, and sheet id, and returns
// a function used to append new rows to a database, given a sheet and values.
//
//	func PrepAppendCall(ctx context.Context, srv *sheets.Service, sheetId string) func(string, []any) (string, error) {
//		return func(sheetName string, values []any) (string, error) {
//			inputLen := len(values)
//			switch {
//			case inputLen > 5:
//				return "", static.ErrInputOutOfRange
//			case inputLen == 0:
//				return "", static.ErrInputEmpty
//			}
//			res, err := srv.Spreadsheets.Values.Append(sheetId, sheetName, &sheets.ValueRange{
//				MajorDimension: "ROWS",
//				Values:         append([][]any{}, values),
//			}).ValueInputOption("RAW").Context(ctx).Do()
//			if err != nil {
//				return "", err
//			}
//			return fmt.Sprintf("append response: %+v\n", res), err
//		}
//	}
func PrepAppendCall(ctx context.Context, srv *sheets.Service, sheetId string) func(string, [5]any) (string, error) {
	return func(sheetName string, values [5]any) (string, error) {
		// inputLen := len(values)
		// switch {
		// case inputLen > 5:
		// 	return "", static.ErrInputOutOfRange
		// case inputLen == 0:
		// 	return "", static.ErrInputEmpty
		// }
		res, err := srv.Spreadsheets.Values.Append(sheetId, sheetName, &sheets.ValueRange{
			MajorDimension: "ROWS",
			Values:         append([][]any{}, values[:]),
		}).ValueInputOption("RAW").Context(ctx).Do()
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("append response: %+v\n", res), err
	}
}

// TODO: Write body for prepReadCall.
// prepReadCall takes a context, sheets service, and sheet id, and returns
// a function used to read database rows given a range.
func PrepReadCall(ctx context.Context, srv *sheets.Service, sheetId string) func(string) (string, error) {
	return func(sRange string) (string, error) {
		return "Read call being implemented", nil
	}
}

// TODO: Write body for prepDeleteCall.
// prepDeleteCall takes a context, sheets service, and sheet id, and returns
// a function used to read database rows given a range.
func PrepDeleteCall(ctx context.Context, srv *sheets.Service, sheetId string) func(string) (string, error) {
	return func(sRange string) (string, error) {
		return "Delete call being implemented", nil
	}
}
