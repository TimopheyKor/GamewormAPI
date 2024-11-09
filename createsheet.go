package main

import "google.golang.org/api/sheets/v4"

func newSpreadsheet(srv *sheets.Service, name string) (string, error) {
	sheet := srv.Spreadsheets.Create()
}
