package sheets

type expenses map[string]int

type OverAllExpenses struct {
	DailyExpenses DailyExpense
	FixedExpenses FixedExpense
	Misc          Miscellaneous
}

type DailyExpense struct {
	Breakfast      int
	Lunch          int
	Dinner         int
	Others         int
	Transportation int
}

type FixedExpense struct {
	firstHalf  expenses
	secondHalf expenses
}

type Miscellaneous struct {
	firstHalf  expenses
	secondHalf expenses
}
