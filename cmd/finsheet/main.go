package main

import (
	"log"
	"financial-system/infrastructure/sheets"
	"financial-system/usecase"
	"financial-system/interfaces"
	"fmt"
	"os"
	"financial-system/domain"
	"time"
)

func init() {
	file, err := os.OpenFile("logs.log", os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(file)
	log.SetFlags(log.Ltime|log.LUTC)
}

func main() {

	sheetIds, err := sheets.LoadSheetIDs()
	if err != nil {
		log.Fatal(err)
	}

	id, err := sheetIds.Get("june")
	if err != nil {
		log.Fatal(err)
	}

	sheetService, err := sheets.Default(id)
	if err != nil {
		log.Fatal(err)
	}

	frepo := interfaces.NewFinanceRepo(sheetService)

	fi := usecase.NewFinanceInteractor(frepo)

	exp, err := fi.GetDayExpense(10)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(exp)

	otherExpense := domain.NewExpense().SetDay(10).SetValue(10).SetName("others")
	err = fi.SetExpense(otherExpense)
	if err != nil {
		log.Fatal(err)
	}

	dailyExpenseSummaryfunc := usecase.YourExpenseFunc(fi)
	expStr, err := dailyExpenseSummaryfunc(1)
	fmt.Println(expStr)

	time.Sleep(2 * time.Second)

	expStr, err = dailyExpenseSummaryfunc(2)
	fmt.Println(expStr)
}

func handleError(key int, err error, message string) {
	log.SetPrefix(fmt.Sprintf("[logID: %v]: ", key))
	log.Printf("%#v", err)
	fmt.Printf("[%v] %v", key, message)
}