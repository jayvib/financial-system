package sheets

import (
	"net/http"
	"google.golang.org/api/sheets/v4"
)

type Sheet struct {
	Service *sheets.Service
	MonthSheet *MonthSheetID
}

func New(client *http.Client, month,sheetId string) (*Sheet, error) {
	servc, err := sheets.New(client)
	if err != nil {
		return nil, err
	}

	sheet :=  &Sheet{
		Service: servc,
		MonthSheet: &MonthSheetID{
			monthSheetId: make(monthSheetId),
			sheetIdPath: sheetIDPath,
		},
	}

	err = sheet.MonthSheet.SetSheetId(month, sheetId)
	if err != nil {
		return nil, err
	}

	return sheet, nil
}
