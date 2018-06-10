package sheets

import (
	"errors"
	"strings"
)

type ProvideRangeFunc func() (string, string, string)

type SheetRange struct {
	SheetName string
	From      string
	To        string
}

func (s *SheetRange) rangeAddress() string {
	if s.To == "" {
		return  s.SheetName + "!" + s.From
	}
	return s.SheetName + "!" + s.From + ":" + s.To
}

func newDayExpenseRanger(day string) sheetRanger {
	return &SheetRange{
		SheetName: day,
		From: "C4",
		To: "E9",
	}
}

func expenseRange(sheetname string, expense string) (sheetRanger, error) {
	sr := new(SheetRange)
	switch strings.ToLower(expense) {
	case Lunch:
		sr.SheetName = sheetname
		sr.From = LunchRange
		return  sr, nil
	case Breakfast:
		sr.SheetName = sheetname
		sr.From = BreakfastRange
		return sr, nil
	case Dinner:
		sr.SheetName = sheetname
		sr.From = DinnerRange
		return sr, nil
	case Others:
		sr.SheetName = sheetname
		sr.From = OtherRange
		return sr, nil
	case Transportation:
		sr.SheetName = sheetname
		sr.From = TransportRange
		return sr, nil
	default:
		return nil, errors.New("sheets.v2: not found expense type")
	}
}

func neutralizeString(prefix string, data interface{}) string {
	str := data.(string)
	return strings.Replace(str, prefix, "", -1)
}
