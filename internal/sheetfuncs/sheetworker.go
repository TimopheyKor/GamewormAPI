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
	Ctx     context.Context
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

// Consider having the two delete functions be a single function with a variatic
// parameter of table names. Delete from all the table names provided, if the
// Games table is included, then delete from all tables.
// TODO: Implement DeleteGame.
// DeleteGame removes the game provided from the tables provided. If the Games
// table is provided, delete from all tables.
func (w *SheetsHolder) DeleteGame(gameId string, tables ...string) (string, error) {
	for _, table := range tables {

	}
}

// TODO: Implement GetGames. Decide if it should be table-specific or split
// into different functions for different tables (as it would be returning
// different length sets of arrays based on the table.)
func (w *SheetsHolder) GetGames(table string) {}
