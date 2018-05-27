package sheets

import (
	"net/http"
	"google.golang.org/api/sheets/v4"
	"errors"
	"strconv"
	"text/tabwriter"
	"fmt"
	"strings"
)

const (
	VALUE_OPTION_USER_ENTERED = "USER_ENTERED"
	VALUE_OPTION_RAW = "RAW"
	VALUE_OPTION_INPUT_VALUE_UNSPECIFIED = "INPUT_VALUE_OPTION_UNSPECIFIED"
)

type SheetInfo struct {
	Title string
	SheetId string
}

type ValueGetter struct {
	service *sheets.Service
}

func (vg *ValueGetter) Get(range_ string, sheetId string) (*sheets.ValueRange, error) {
	return vg.service.Spreadsheets.Values.Get(sheetId, range_).Do()
}

type ValueSetter struct {
	service *sheets.Service
}

func (vs *ValueSetter) Set(range_ string, sheetId string, valueRange *sheets.ValueRange) error {
	sheetUpdateCall := vs.service.Spreadsheets.Values.Update(sheetId, range_, valueRange).ValueInputOption(VALUE_OPTION_USER_ENTERED)
	_, err := sheetUpdateCall.Do()
	if err != nil {
		return err
	}
	return nil
}

type ValueGetterSetter struct {
	Setter
	Getter
}

type SpreadsheetService struct {
	SetGetter
	Service *sheets.Service
	SheetInfo SheetInfo
}

func NewSpreadsheetService(client *http.Client, sheetId string, opts ...func(service *SpreadsheetService)) (*SpreadsheetService, error) {
	service, err := sheets.New(client)
	if err != nil {
		return nil, err
	}

	vgs := &ValueGetterSetter{
		Setter: &ValueSetter{
			service: service,
		},
		Getter: &ValueGetter{
			service: service,
		},
	}

	spreadsheetService := &SpreadsheetService{
		vgs,
		service,
		SheetInfo{
			SheetId: sheetId,
		},
	}

	for _, opt := range opts {
		opt(spreadsheetService)
	}

	return spreadsheetService, nil
}

func (s *SpreadsheetService) GetValue(range_ string) (*sheets.ValueRange, error) {
	return s.SetGetter.Get(range_, s.SheetInfo.SheetId)
}

func (s *SpreadsheetService) SetValue(range_ string, valueRange *sheets.ValueRange) error {
	return s.SetGetter.Set(range_, s.SheetInfo.SheetId, valueRange)
}

func (s *SpreadsheetService) ProvideSheetID() string {
	return s.SheetInfo.SheetId
}


func GetDayExpense(getter GetSheetIDProvider, rp SheetRangeProvider) (*DayExpense, error) {
	vr, err := getter.Get(rp.ProvideAddress(), getter.ProvideSheetID())
	if err != nil {
		return nil, err
	}

	if vr.HTTPStatusCode != 200 {
		return nil, errors.New("sheets.v2: invalid response from the server")
	}

	dayExpense := new(DayExpense)
	dayExpense.parse(vr.Values)
	return dayExpense, nil
}

func SetDayExpense(setter SetSheetIDProvider, rp SheetRangeProvider, valueRange *sheets.ValueRange) error {
	return setter.Set(rp.ProvideAddress(), setter.ProvideSheetID(), valueRange)
}

type DayExpense struct {
	FoodExpense FoodExpense
	Transportation int
}

func (de *DayExpense) Total() int {
	 return de.FoodExpense.Total() + de.Transportation
}

func (de *DayExpense) String() string {
	strBuilder := &strings.Builder{}
	tabWriter := tabwriter.NewWriter(strBuilder, 0, 20, 1, ' ', tabwriter.TabIndent)
	fmt.Fprintln(tabWriter, "\tExpense\tAmount\tTotal")
	fmt.Fprintln(tabWriter, "Food\t--------------------")
	fmt.Fprintf(tabWriter, "\tBFast\tPhp%d\n", de.FoodExpense.Breakfast)
	fmt.Fprintf(tabWriter, "\tLunch\tPhp%d\n", de.FoodExpense.Lunch)
	fmt.Fprintf(tabWriter, "\tDinner\tPhp%d\n", de.FoodExpense.Dinner)
	fmt.Fprintln(tabWriter, "Transpo\t--------------------")
	fmt.Fprintf(tabWriter, "\tTranspo\tPhp%d\n", de.Transportation)
	fmt.Fprintf(tabWriter, "\t\t\tPhp%d", de.Total())
	tabWriter.Flush()
	return strBuilder.String()
}

func (de *DayExpense) parse(data [][]interface{}) error {
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

	de.FoodExpense.Breakfast = bfastVal
	de.FoodExpense.Lunch = lunchVal
	de.FoodExpense.Dinner = dinnerVal
	de.FoodExpense.Others = othersVal
	de.Transportation = transpoVal
	return nil
}

func NewValueRange(val int) *sheets.ValueRange {
	return &sheets.ValueRange{
		Values: [][]interface{}{{val}},
	}
}