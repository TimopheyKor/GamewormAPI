package sheetfuncs

import (
	"context"
	"fmt"

	"github.com/TimopheyKor/GamewormAPI/internal/schema"
	"github.com/TimopheyKor/GamewormAPI/internal/static"
	"google.golang.org/api/sheets/v4"
)

// SheetsHolder retains essential and frequently reused data for interacting
// with a Google Spreadsheet via the Google Sheets v4 API.
type SheetsHolder struct {
	Ctx     context.Context //TODO: handle context differently as it can change with requests from frontend
	Srv     *sheets.Service
	SheetId string
}

// NewSheetsHolder takes a context, sheets service, and sheetId, and
// creates and initializes a new SheetsHolder, which is used
// to run functions on a given sheet.
func NewSheetsHolder(ctx context.Context, srv *sheets.Service, sheetId string) *SheetsHolder {
	return &SheetsHolder{
		Ctx:     ctx,
		Srv:     srv,
		SheetId: sheetId,
	}
}

// GameIdExists takes a gameId and table name, returning True if
// it already exists in the table, false otherwise.
func (w *SheetsHolder) GameIdExists(gameId, table string) (bool, error) {
	res, err := w.Srv.Spreadsheets.Values.Get(w.SheetId, table+"!A1:A").Do()
	if err != nil {
		return false, err
	}
	if len(res.Values) == 0 {
		return false, nil
	}
	for _, row := range res.Values {
		if row[static.GamePK] == gameId {
			return true, nil
		}
	}
	return false, nil
}

// TODO: Implement more schema assertion checks as part of the AddNewGame func.
// AddNewGame takes an array of values and attempts to append them to the
// Game table of the GamewormDB spreadsheet, returning the HTTP response and
// an error.
func (w *SheetsHolder) AddNewGame(g *schema.GameObject) (string, error) {
	// Check if the game already exists:
	gameExists, err := w.GameIdExists(g.GetID(), static.GameD)
	if err != nil {
		return "", err
	}
	if gameExists {
		return "", static.ErrDuplicateGameID
	}

	// Append a new row with that game's data:
	range_ := static.GameD + "!" + static.GameRange
	res, err := w.Srv.Spreadsheets.Values.Append(w.SheetId, range_, &sheets.ValueRange{
		MajorDimension: "ROWS",
		Values:         append([][]any{}, g.ToSlice()),
	}).ValueInputOption("RAW").Context(w.Ctx).Do()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("append response: %+v", res), nil
}

// DeleteGame removes the game provided from the tables provided.
// TODO: Add functionality so that if the Games table is provided, delete from
// all tables.
func (w *SheetsHolder) DeleteGame(gameId string, tables ...string) (string, error) {
	res := ""
	for _, table := range tables {
		readRange := table + "!A1:A"
		data, err := w.Srv.Spreadsheets.Values.Get(w.SheetId, readRange).Do()
		if err != nil {
			return "", err
		}
		if len(data.Values) == 0 {
			return "", static.ErrNoDataFound
		} else {
			for rowIdx, row := range data.Values {
				if row[static.GamePK] == gameId {
					res, err = w.deleteRow(int64(rowIdx), table)
					if err != nil {
						return fmt.Sprintf("delete row response: %+v", res), err
					}
				}
			}
		}
	}
	return fmt.Sprintf("delete game response: %+v", res), nil
}

// deleteRow is used to delete a row of data from a table given a row index.
func (w *SheetsHolder) deleteRow(rowIdx int64, table string) (string, error) {
	tableSheetId, err := w.getTableSheetId(table)
	if err != nil {
		return "get table sheet id response:", err
	}
	res, err := w.Srv.Spreadsheets.BatchUpdate(w.SheetId, &sheets.BatchUpdateSpreadsheetRequest{
		IncludeSpreadsheetInResponse: false,
		Requests: []*sheets.Request{
			{
				DeleteRange: &sheets.DeleteRangeRequest{
					Range: &sheets.GridRange{
						SheetId:          tableSheetId,
						StartRowIndex:    rowIdx,
						EndRowIndex:      rowIdx + 1,
						StartColumnIndex: 0,
						EndColumnIndex:   static.MaxColIdx + 1,
					},
					ShiftDimension: "ROWS",
				},
			},
		},
	}).Context(w.Ctx).Do()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("batch update response: %+v", res), nil
}

// TODO: Consider what formats should be returned for the GET calls.
// Maybe a map with an object representing each row of a table and the gameIDs as keys?
// TODO: Implement GetGames. Decide if it should be table-specific or split
// into different functions for different tables (as it would be returning
// different length sets of arrays based on the table.)
// If building Structs or JSON files to be cashed on the browser, there should
// be a different function for each table as pulling its data may have
// a different return type.
// GetGames returns a list of all game objects from the Games table in an
// api-readable formate.
func (w *SheetsHolder) GetGames() {}

// GetReviews returns a list of all reviews as objects tied to a gameID in an
// api-readable format.
func (w *SheetsHolder) GetReviews() {}

// GetBacklog returns a list of all backlog rows as objects tied to a gameID in
// an api-readable format.
func (w *SheetsHolder) GetBacklog() {}

// GetGame returns a single game in API-readable format.
// TODO: Decide if this should be by-name or by-id.
// At some point we'll need search functionality - but that might belong on
// the frontend.
func (w *SheetsHolder) GetGame() {}

// Update functions:
func (w *SheetsHolder) UpdateGame(gameId string)                     {}
func (w *SheetsHolder) AddReview(gameId string, options []string)    {}
func (w *SheetsHolder) UpdateReview(gameId string, options []string) {}

// TODO: Backlog additions and updates will need additional logic in order to
// keep the priority rankings of games in the backlog proper.
func (w *SheetsHolder) AddToBacklog(gameId string, options []string)  {}
func (w *SheetsHolder) UpdateBacklog(gameId string, options []string) {}

// getTableSheetId gets the SheetId of a specific table given the table's
// sheet name.
func (w *SheetsHolder) getTableSheetId(table string) (int64, error) {
	spreadsheetData, err := w.Srv.Spreadsheets.Get(w.SheetId).Context(w.Ctx).Do()
	if err != nil {
		return 0, err
	}
	for _, sheet := range spreadsheetData.Sheets {
		if sheet.Properties.Title == table {
			return sheet.Properties.SheetId, nil
		}
	}
	return 0, static.ErrSheetNotFound
}
