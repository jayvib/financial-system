package sheets

import (
	"testing"
	"financial-system/test/data"
	"text/tabwriter"
	"os"
	"fmt"
)

func TestNewSheet(t *testing.T) {
	for _, data := range data.SheetNameSheetID {
		sheet := New(data.Name, data.SheetID)
		if sheet == nil {
			t.Fatal("sheet must not be nil")
		}
	}
}

func TestMonthSheetID(t *testing.T) {
	sheetIds := MonthSheetID{sheetIdPath:sheetIDPath}

	t.Run("loadSheetIds", func(t *testing.T){
		t.Skip()
		err := sheetIds.loadSheetIds()
		if err != nil {
			t.Error(err)
		}

		if len(sheetIds.monthSheetId) == 0 {
			t.Error("sheet IDs must not be zero")
		}

		if len(sheetIds.monthSheetId) != 3 {
			t.Errorf("expecting the items of the sheetIds.monthSheetId to be 3 but got %d\n", len(sheetIds.monthSheetId))
		}

		tw := tabwriter.NewWriter(os.Stderr, 0, 0, 3, ' ', tabwriter.AlignRight|tabwriter.Debug)
		for mo, id := range sheetIds.monthSheetId {
			fmt.Fprintf(tw, "%s\t%s", mo, id)
		}
	})


	t.Run("setSheetIds", func(t *testing.T) {
		err := sheetIds.SetSheetId("April", "da4d4fa98e7a54df")
		if err != nil {
			t.Error(err)
		}

		err = sheetIds.loadSheetIds()
		if err != nil {
			t.Error(err)
		}

		if len(sheetIds.monthSheetId) != 4 {
			t.Errorf("expecting the items of the sheetIds.monthSheetId to be 4 but got %d\n", len(sheetIds.monthSheetId))
		}
	})
}