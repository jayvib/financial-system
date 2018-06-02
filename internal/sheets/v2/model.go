package sheets

type FoodExpense struct {
	Breakfast int
	Lunch int
	Dinner int
	Others int
}

func (fe *FoodExpense) Total() int {
	return fe.Breakfast + fe.Lunch + fe.Dinner + fe.Others
}

