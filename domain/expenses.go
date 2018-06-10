package domain

import "fmt"

type Expenser interface {
	Total() int
	fmt.Stringer
}

type Expense struct {
	name string
	value int
	month int
	day int
}

func NewExpense() *Expense {
	return &Expense{}
}

func (e *Expense) SetName(name string) *Expense {
	e.name = name
	return e
}

func (e *Expense) SetValue(value int) *Expense {
	e.value = value
	return e
}

func (e *Expense) SetMonth(month int) *Expense {
	e.month = month
	return e
}

func (e *Expense) SetDay(day int) *Expense {
	e.day = day
	return e
}

type DayExpense struct {
	MealExpense
	Transportation *Expense
	Month int
	Day int
}

type MealExpense struct {
	Breakfast *Expense
	Lunch *Expense
	Dinner *Expense
	Others *Expense
}

type FixedExpenses map[string]*Expense

type MiscExpenses map[string]*Expense

type OverallExpenses struct {
	DailyExpenses []*Expense
	FixedExpenses FixedExpenses
	MiscExpenses MiscExpenses
}

type MonthExpenses struct {
	FirstHalf map[string]DayExpense
	SecondHalf map[string]DayExpense
}