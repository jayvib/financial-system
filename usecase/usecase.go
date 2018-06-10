package usecase

import "financial-system/domain"

// FinanceInteractor will be the object to provide functionalities to the application.
type FinanceInteractor struct {
	financeRepo domain.FinanceRepository
}

func (f *FinanceInteractor) GetDayExpense(exp *domain.DayExpense) (*domain.DayExpense, error) {
	return f.financeRepo.GetDayExpense(exp)
}


func (f *FinanceInteractor) SetDayExpense(exp *domain.DayExpense) error {
	return f.financeRepo.SetDayExpense(exp)
}



