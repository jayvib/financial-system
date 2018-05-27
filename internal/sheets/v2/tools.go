package sheets

type ProvideRangeFunc func() (string, string, string)

type SheetRange struct {
	SheetName string
	From      string
	To        string
}

func (s *SheetRange) ProvideAddress() string {
	return s.SheetName + "!" + s.From + ":" + s.To
}

func DayExpenseRangeProvider(day string) *SheetRange {
	return &SheetRange{
		SheetName: day,
		From: "C4",
		To: "E9",
	}
}
