package main

import (
	"financial-system/config"

	fssheets "financial-system/internal/sheets"
	"log"
	"fmt"
)

const configFile = "client_secret.json"

var spreadsheetId = "1LBUAV99tHgD7RpdMekiDwVRE5A29B7L49wpHdJAmPSA"

func main() {
	// If modifying these scopes, delete your previously saved client_secret.json.
	client, err := config.NewClient(configFile, config.ReadWriteConfigFunc)

	service, err := fssheets.New(client, "testmonth", spreadsheetId)
	if err != nil {
		log.Fatalf("Unable to retrieve Sheets client: %v", err)
	}

	readRange := "Summary!G9"
	//
	resp, err := service.GetValues(readRange)
	if err != nil {
		log.Fatalf("Unable to retrieve data from sheet: %v", err)
	}

	totalExp := resp[0][0]
	fmt.Println(totalExp)
}