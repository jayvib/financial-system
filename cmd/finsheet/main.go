package main

import (
	"log"
	"financial-system/infrastructure/sheets"
	"fmt"
)

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
	exp, err := sheetService.GetDayExpense(2)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(exp)
}
