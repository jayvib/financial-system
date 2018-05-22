package main

import (
	"financial-system/internal/sheets"
	"log"
	"fmt"
)

const configFile = "client_secret.json"

var spreadsheetId = "1LBUAV99tHgD7RpdMekiDwVRE5A29B7L49wpHdJAmPSA"

func main() {

	ids, err := sheets.LoadSheetIDs()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(ids.GetAllSheetIDs())
}