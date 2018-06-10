package domain

type FinanceRepository interface {
	DayExpenseGetter
	FixedExpenseGetter
	MiscExpenseGetter
	ExpenseSetter
}

// =============An example for the Interface Segregation Principle =============

type DayExpenseGetter interface {
	GetDayExpense(day int) (*DayExpense, error)
}

type FixedExpenseGetter interface {
	GetFixedExpense() (*FixedExpenses, error)
}

type MiscExpenseGetter interface {
	GetMiscExpense() (*MiscExpenses, error)
}

type ExpenseSetter interface {
	SetExpense(expense *Expense) error
}
