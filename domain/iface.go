package domain

type FinanceRepository interface {
	GetDayExpense(expense *DayExpense) (*DayExpense, error)
	GetFixedExpense() (*FixedExpenses, error)
	GetMiscExpense() (*MiscExpenses, error)
	SetDayExpense(expense *DayExpense) error
}

