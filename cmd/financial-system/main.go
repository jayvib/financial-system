package main

import (
	"financial-system/config"
	"financial-system/internal/sheets/v2"
	"fmt"
	"log"
	"os"
)

const configFile = "client_secret.json"

var spreadsheetId = "1LBUAV99tHgD7RpdMekiDwVRE5A29B7L49wpHdJAmPSA"

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
	file, err := os.Open(configFile)
	if err != nil {
		log.Fatal(err)
	}

	client, err := config.NewClient(file, config.ReadWriteConfigFunc)
	if err != nil {
		fmt.Println("error: ", err.Error())
		log.Fatal(err)

	}

	sheetService, err := sheets.NewSpreadsheetService(client, spreadsheetId)
	if err != nil {
		log.Fatal(err)
	}

	dayExpRange := sheets.DayExpenseRangeProvider(sheets.SecondDay)

	dayExpense, err := sheets.GetDayExpense(sheetService, dayExpRange)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(dayExpense)
}
