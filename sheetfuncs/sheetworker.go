package sheetfuncs

import (
	"context"

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

// TODO: Implement GameIdExists
// GameIdExists takes a SheetsHelper object and a GameID, returning True if
// it already exists in the database, false otherwise.
func (w *SheetWorker) GameIdExists(gameId string) bool {
	return false
}
