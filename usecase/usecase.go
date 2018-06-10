package usecase

import "financial-system/domain"

// FinanceInteractor will be the object to provide functionalities to the application.
type FinanceInteractor struct {
	financeRepo domain.FinanceRepository
}

func (f *FinanceInteractor) GetDayExpense(day int) (*domain.DayExpense, error) {
	return f.financeRepo.GetDayExpense(day)
}


func (f *FinanceInteractor) SetExpense(exp *domain.Expense) error {
	return f.financeRepo.SetExpense(exp)
}

func NewFinanceInteractor(frepo domain.FinanceRepository) *FinanceInteractor {
	return &FinanceInteractor{
		financeRepo: frepo,
	}
}

// ===== An Example how to implement the ISP =======

// combination of ISP and closure FP
func YourExpenseFunc(getter domain.DayExpenseGetter) func(day int) (string, error) {
	return func(day int) (string, error) {
		de, err := getter.GetDayExpense(day)
		if err != nil {
			return "<nul>", err
		}
		return de.String(), nil
	}
}