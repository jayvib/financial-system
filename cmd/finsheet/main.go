package main

import (
	"github.com/urfave/cli"
	"os"
	"log"
	"fmt"
	"github.com/pkg/errors"
	fcli "financial-system/internal/cli"
	"financial-system/internal/sheets"
	sheets2 "financial-system/internal/sheets/v2"
	"financial-system/config"
)

const configFile = "client_secret.json"

var spreadsheetId = "1mETJWembSRICxS_pvPnezMPEDL9nnRhtLLZ-zTJkjFc"

var (
	month string
	day string
	expense int
	mode string
	expRange string
)


func main() {
	file, err := os.Open(configFile)
	if err != nil {
		log.Fatal(err)
	}

	client, err := config.NewClient(file, config.ReadWriteConfigFunc)
	if err != nil {
		log.Fatal(err)
	}

	service, err := sheets2.NewSpreadsheetService(client, spreadsheetId)
	if err != nil {
		log.Fatal(err)
	}

	app := cli.NewApp()
	app.Name = "finsheet"
	app.Usage = "A command line version of Jayson's financial system in google spreadsheet"
	app.Flags = []cli.Flag {
		cli.StringFlag{
			Name: "month, mo",
			Usage: "month is the month where the expense came from",
			Destination: &month,
		},
		cli.StringFlag{
			Name: "mode",
			Usage: "mode is what operation to do. Possible values are either [get/set]",
		},
		cli.StringFlag{
			Name: "day, d",
			Usage: "day is the day where the expense came from",
			Destination: &day,
		},
		cli.IntFlag{
			Name: "expense, e",
			Usage: "expense is expense of the day.",
		},
	}


	app.Action = func(c *cli.Context) error {
		switch c.GlobalString("mode") {
		case fcli.GET:

			switch expense {
			case sheets2.Breakfast:
				expRange = sheets2.BreakfastRange
			case sheets2.Lunch:
				expRange = sheets2.LunchRange
			case sheets2.Dinner:
				expRange = sheets.DinnerRange
			case sheets2.Others:
				expRange = sheets2.OtherRange
			case sheets2.Transportation:
				expRange = sheets2.TransportRange
			default:
				fmt.Println("Expense doen't recognize")
				return errors.New("expense doesn't recognize")
			}

			fmt.Println("day value", day)
			sp := sheets2.DayExpenseRangeProvider(day)
			de, err := sheets2.GetDayExpense(service, sp)
			if err != nil {
				return errors.Wrap(err, "unexpected error while getting the day expense")
			}
			fmt.Print(de)

		case fcli.SET:

		default:
			fmt.Println("Mode doesn't recognize")
			return errors.New("mode doesn't recognize")
		}
		return nil
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
