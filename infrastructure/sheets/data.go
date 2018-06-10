package sheets

import (
	"fmt"
	"strconv"
	"financial-system/domain"
)

type foodExpense struct {
	breakfast int
	lunch int
	dinner int
	others int
}

func (fe *foodExpense) Total() int {
	return fe.breakfast + fe.lunch + fe.dinner + fe.others
}

type dayExpense struct {
	foodExpense    foodExpense
	transportation int
}

func (de *dayExpense) parse(data [][]interface{}) error {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("There's an expense that don't have value.")
		}
	}()
	breakfastIndex := 1
	lunchIndex     := 2
	dinnerIndex    := 3
	othersIndex    := 4
	transpoIndex   := 5
	prefix         := "Php"

	val := 2

	bfastValRaw := data[breakfastIndex][val]
	lunchValRaw := data[lunchIndex][val]
	dinnerValRaw := data[dinnerIndex][val]
	othersValRaw := data[othersIndex][val]
	transpoValRaw := data[transpoIndex][val]

	bfastVal, err := strconv.Atoi(neutralizeString(prefix, bfastValRaw))
	if err != nil {
		return err
	}

	lunchVal, err := strconv.Atoi(neutralizeString(prefix, lunchValRaw))
	if err != nil {
		return err
	}

	dinnerVal, err := strconv.Atoi(neutralizeString(prefix, dinnerValRaw))
	if err != nil {
		return err
	}

	othersVal, err := strconv.Atoi(neutralizeString(prefix, othersValRaw))
	if err != nil {
		return err
	}

	transpoVal, err := strconv.Atoi(neutralizeString(prefix, transpoValRaw))
	if err != nil {
		return err
	}

	de.foodExpense.breakfast = bfastVal
	de.foodExpense.lunch = lunchVal
	de.foodExpense.dinner = dinnerVal
	de.foodExpense.others = othersVal
	de.transportation = transpoVal
	return nil
}

func (de *dayExpense) toDayExpense() *domain.DayExpense {
	expense := &domain.DayExpense{
		MealExpense: domain.MealExpense{
			Breakfast: domain.NewExpense().SetName("Breakfast").SetValue(de.foodExpense.breakfast),
			Lunch: domain.NewExpense().SetName("Lunch").SetValue(de.foodExpense.lunch),
			Dinner: domain.NewExpense().SetName("Dinner").SetValue(de.foodExpense.dinner),
			Others: domain.NewExpense().SetName("Others").SetValue(de.foodExpense.others),
		},
		Transportation: domain.NewExpense().SetName("Transportation").SetValue(de.transportation),
	}
	return expense
}