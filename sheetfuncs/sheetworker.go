package sheetfuncs

import (
	"context"
	"fmt"

	"github.com/TimopheyKor/GamewormAPI/static"
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

// TODO: Implement schema assertion checks as part of the AddNewGame func.
// AddNewGame takes an array of values and attempts to append them to the
// Game table of the GamewormDB spreadsheet, returning the HTTP response and
// an error.
func (w *SheetsHolder) AddNewGame(values []any) (string, error) {
	range_ := static.GameD + "!" + static.GameRange
	res, err := w.Srv.Spreadsheets.Values.Append(w.SheetId, range_, &sheets.ValueRange{
		MajorDimension: "ROWS",
		Values:         append([][]any{}, values),
	}).ValueInputOption("RAW").Context(w.Ctx).Do()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("append response: %+v", res), nil
}
