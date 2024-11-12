package sheetfuncs

import (
	"context"

	"github.com/TimopheyKor/GamewormAPI/static"
	"google.golang.org/api/sheets/v4"
)

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
