package cli

import (
	"github.com/urfave/cli"
	"fmt"
	"github.com/pkg/errors"
	"financial-system/internal/sheets/v2"
	"financial-system/client"
)

var (
	isDebug = false
)

func FlagInitializer(a *App) {
	a.Name = "finsheet"
	a.Usage = "A simple command line tool for financial system in Google spreadsheet"
	a.App.Usage = "A command line version of Jayson's financial system in google spreadsheet"
	a.App.Flags = []cli.Flag {
		cli.StringFlag{
			Name: "mode",
			Usage: "mode is what operation to do. Possible values are either [get/set]",
		},
		cli.StringFlag{
			Name: "day, d",
			Usage: "day is the day where the expense came from",
		},
		cli.IntFlag{
			Name: "expense, e",
			Usage: "expense is expense of the day.",
		},
		cli.IntFlag{
			Name: "val",
			Usage: "value for the expense",
		},
		cli.StringFlag{
			Name: "month, mo",
			Usage: "month is what month expense will be set",
		},
	}
}

func ActionInitializer(a *App) {
	a.App.Action = func(c *cli.Context) error {
		expRange := ""
		value := c.GlobalInt("val")
		sheetname_ := c.GlobalString("day")
		month := c.GlobalString("month")

		ids, err := sheets.LoadSheetIDs()
		if err != nil {
			return errors.Wrap(err, "error while loading the sheet IDs")
		}
		cl, err := client.DefaultClient()
		if err != nil {
			return errors.Wrap(err, "error while initializing default client")
		}

		id, err := ids.Get(month)
		if err != nil {
			return err
		}

		service, err := sheets.New(cl, id)

		getExpenseHandler := func(day string) (*sheets.DayExpense, error) {
			sp := sheets.DayExpenseRangeProvider(day)
			de, err := sheets.GetDayExpense(service, sp)
			if err != nil {
				return nil, errors.Wrap(err, "unexpected error while getting the day expense")
			}
			return de, nil
		}


		switch c.GlobalString("mode") {
		case GET:
			de, err := getExpenseHandler(sheetname_)
			if err != nil {
				return err
			}
			fmt.Print(de)

		case SET:
			switch c.GlobalInt("expense") {
			case sheets.Breakfast:
				expRange = sheets.BreakfastRange
			case sheets.Lunch:
				expRange = sheets.LunchRange
			case sheets.Dinner:
				expRange = sheets.DinnerRange
			case sheets.Others:
				expRange = sheets.OtherRange
			case sheets.Transportation:
				expRange = sheets.TransportRange
			default:
				fmt.Println("Expense doesn't recognize")
				return errors.New("expense doesn't recognize")
			}

			sp := &sheets.SheetRange{
				SheetName: sheetname_,
				From:      expRange,
			}

			if isDebug {
				fmt.Printf("[Values] value: %d\n", value)
				fmt.Printf("[expense] expense: %d\n", c.GlobalInt("expense"))
			}
			valRange := sheets.NewValueRange(value)

			err := sheets.SetDayExpense(service, sp, valRange)
			if err != nil {
				return errors.Wrapf(err, "failed setting day expense")
			}

			de, err := getExpenseHandler(sheetname_)
			if err != nil {
				return err
			}

			fmt.Println(de)

		default:
			fmt.Println("Mode doesn't recognize")
			return errors.New("mode doesn't recognize")
		}
		return nil
	}
}