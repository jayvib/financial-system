package domain

import (
	"fmt"
	"strings"
	"text/tabwriter"
)

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

func (e *Expense) GetName() string {
	return e.name
}

func (e *Expense) GetValue() int {
	return e.value
}

func (e *Expense) GetMonth() int {
	return e.month
}

func (e *Expense) GetDay() int {
	return e.day
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

func (de *DayExpense) Total() int {
	return de.MealExpense.Total() + de.Transportation.value
}

func (de *DayExpense) String() string {
	strBuilder := &strings.Builder{}
	tabWriter := tabwriter.NewWriter(strBuilder, 0, 20, 1, ' ', tabwriter.TabIndent)
	fmt.Fprintln(tabWriter, "\tExpense\tAmount\tTotal")
	fmt.Fprintln(tabWriter, "Food\t--------------------")
	fmt.Fprintf(tabWriter, "\tBFast\tPhp%d\n", de.MealExpense.Breakfast.value)
	fmt.Fprintf(tabWriter, "\tLunch\tPhp%d\n", de.MealExpense.Lunch.value)
	fmt.Fprintf(tabWriter, "\tDinner\tPhp%d\n", de.MealExpense.Dinner.value)
	fmt.Fprintf(tabWriter, "\tOthers\tPhp%d\n", de.MealExpense.Others.value)
	fmt.Fprintln(tabWriter, "Transpo\t--------------------")
	fmt.Fprintf(tabWriter, "\tTranspo\tPhp%d\n", de.Transportation.value)
	fmt.Fprintf(tabWriter, "\t\t\tPhp%d\n", de.Total())
	tabWriter.Flush()
	return strBuilder.String()
}

type MealExpense struct {
	Breakfast *Expense
	Lunch *Expense
	Dinner *Expense
	Others *Expense
}

func (me *MealExpense) Total() int {
	return me.Breakfast.value + me.Lunch.value + me.Dinner.value + me.Others.value
}

type FixedExpenses map[string]*Expense

func (fe FixedExpenses) Total() int {
	total := 0
	for _, e := range fe {
		total += e.value
	}
	return total
}

func (fe FixedExpenses) String() string {
	return "<null>"
}

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