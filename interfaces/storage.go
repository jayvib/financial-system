package interfaces

import "financial-system/domain"

type FinanceHandler interface {
	GetDayExpense(day int) (*domain.DayExpense, error)
	GetFixedExpenses() (*domain.FixedExpenses, error)
	GetMiscExpenses() (*domain.MiscExpenses, error)
	SetDayExpense(exp *domain.Expense) error
}