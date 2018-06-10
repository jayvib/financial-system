package interfaces

import (
	"financial-system/domain"
	"errors"
)

var (
	notImplementedErr = errors.New("not implemented yet")
)

// FinanceRepo is the male end for the FinanceRepository in the domain.
// Its field 'storage' will be the female end where to get the source data.
type FinanceRepo struct {
	storage FinanceHandler
}

func (f *FinanceRepo) GetDayExpense(expense *domain.DayExpense) (*domain.DayExpense, error) {
	return nil, notImplementedErr
}

func (f *FinanceRepo) GetFixedExpense() (*domain.FixedExpenses, error) {
	return nil, notImplementedErr
}

func (f *FinanceRepo) GetMiscExpense() (*domain.MiscExpenses, error) {
	return nil, notImplementedErr
}

func (f *FinanceRepo) SetDayExpense(expense *domain.DayExpense) error {
	return notImplementedErr
}
