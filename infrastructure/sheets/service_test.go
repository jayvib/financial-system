package sheets

import (
	"google.golang.org/api/sheets/v4"
	"financial-system/client"
	"testing"
)

func setup(t *testing.T) *sheets.Service {
	cl, err := client.DefaultClient()
	if err != nil {
		t.Fatal(err)
	}

	srvc, err := sheets.New(cl)
	if err != nil {
		t.Fatal(err)
	}
	return srvc
}


func TestSheetService_GetDayExpense(t *testing.T) {
	srvc := setup(t)
	_ = srvc

}

func TestSheetService_GetFixedExpenses(t *testing.T) {
	srvc := setup(t)
	_ = srvc
}

func TestSheetService_GetMiscExpenses(t *testing.T) {
	srvc := setup(t)
	_ = srvc
}

func TestSheetService_SetDayExpense(t *testing.T) {
	srvc := setup(t)
	_ = srvc
}