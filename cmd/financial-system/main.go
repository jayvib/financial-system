package main

import (
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"financial-system/config"
	//"google.golang.org/api/sheets/v4"
	"log"
	"fmt"

	"financial-system/internal/sheets"
)

const configFile = "client_secret.json"

var spreadsheetId = "11RCS6J7S3vknme8EEi640Hxmf-xQv5XPGo0rEr7GVrA"

func main() {
	// If modifying these scopes, delete your previously saved client_secret.json.
	configFunc := func(b []byte) (*oauth2.Config, error) {
		conf, err := google.ConfigFromJSON(b, "https://www.googleapis.com/auth/spreadsheets.readonly")
		if err != nil {
			return nil, err
		}
		return conf, nil
	}

	client, err := config.NewClient(configFile, configFunc)

	sheet, err := sheets.New(client, "testmonth", spreadsheetId)
	if err != nil {
		log.Fatalf("Unable to retrieve Sheets client: %v", err)
	}

	readRange := "test!A1:A5"

	resp, err := sheet.Service.Spreadsheets.Values.Get(spreadsheetId, readRange).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve data from sheet: %v", err)
	}

	if len(resp.Values) == 0 {
		fmt.Println("No data found.")
	} else {
		fmt.Println("Data:")
		for _, row := range resp.Values {
			// Print columns A and E, which correspond to indices 0 and 4.
			fmt.Println(row[0])
		}
	}


}