package sheets

type expense int
type day string

const (
	SummaryRange = "Summary!G9"

	// Meal Flags
	Lunch           = 0
	Breakfast       = 1
	Dinner          = 2
	Others          = 3
	Transportation  = 4

	// Daily Range
	BreakfastRange = "E5"
	LunchRange     = "E6"
	DinnerRange    = "E7"
	OtherRange     = "E8"
	TransportRange = "E9"
	TotalDailyExp  = "F10"

	// Month Half Expense
	FIRST_HALF_MONTH_EXPENSE  = "G12"
	SECOND_HALF_MONTH_EXPENSE = "G15"

	// Sheet Names
	SUMMARY_SHEET = "Summary"

	// TODO: Fill up sequence 1-30
	//Prefix Daily Sheetname
	FirstDay   = "1"
	SecondDay  = "2"
	ThirdDay
	FourthDay
	FifthDay
	SixthDay
)
