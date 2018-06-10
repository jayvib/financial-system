package sheets

import "google.golang.org/api/sheets/v4"

func newValueRange(val int) *sheets.ValueRange {
	return &sheets.ValueRange{
		Values: [][]interface{}{{val}},
	}
}
