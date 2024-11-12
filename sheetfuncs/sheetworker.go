package sheetfuncs

import (
	"context"

	"github.com/TimopheyKor/GamewormAPI/static"
	"google.golang.org/api/sheets/v4"
)

// TODO: Think of a name that doesn't include Worker.
type SheetWorker struct {
	Ctx     context.Context
	Srv     *sheets.Service
	SheetId string
}

// NewSheetWorker takes a context, sheets service, and sheetId, and
// creates and initializes a new SheetWorker, which is used
// to run functions on a given sheet.
func NewSheetWorker(ctx context.Context, srv *sheets.Service, sheetId string) *SheetWorker {
	return &SheetWorker{
		Ctx:     ctx,
		Srv:     srv,
		SheetId: sheetId,
	}
}

// GameIdExists takes a gameId and table name, returning True if
// it already exists in the table, false otherwise.
func (w *SheetWorker) GameIdExists(gameId, table string) (bool, error) {
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

// TODO: Implement AddNewGame so that it replaces the Append functionality.
// This function should specifically be for adding a game to the first sheet -
// adding games to the Reviews section or Backlog should stem from an existing
// game in the first sheet.
func (w *SheetWorker) AddNewGame(range_ string, values []any) (string, error) {
	return "", nil
}

// return func(sheetName string, values []any) (string, error) {
// 	inputLen := len(values)
// 	switch {
// 	case inputLen > 5:
// 		return "", static.ErrInputOutOfRange
// 	case inputLen == 0:
// 		return "", static.ErrInputEmpty
// 	}
// 	res, err := srv.Spreadsheets.Values.Append(sheetId, sheetName, &sheets.ValueRange{
// 		MajorDimension: "ROWS",
// 		Values:         append([][]any{}, values),
// 	}).ValueInputOption("RAW").Context(ctx).Do()
// 	if err != nil {
// 		return "", err
// 	}
// 	return fmt.Sprintf("append response: %+v\n", res), err
// }
