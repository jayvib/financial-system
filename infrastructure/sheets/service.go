package sheets

import (
	"financial-system/domain"
	"errors"
	"google.golang.org/api/sheets/v4"
	"net/http"
	"financial-system/client"
	"strconv"
)

var (
	notImplementedErr = errors.New("not implemented yet")
)

type optionFunc func(sheetService *SheetService)

func New(sheetid string , cl *http.Client, opts ...optionFunc) (*SheetService, error) {
	srvc, err := sheets.New(cl)
	if err != nil {
		return nil, err
	}

	ss := new(SheetService)
	ss.sheetInfo = sheetInfo{ sheetId: sheetid }
	ss.service.srvc = srvc

	for _, opt := range opts {
		opt(ss)
	}

	return ss, nil

}

func Default(sheetId string) (*SheetService, error) {
	cl, err := client.DefaultClient()
	if err != nil {
		return nil, err
	}
	return New(sheetId, cl)
}

type SheetService struct {
	service
}

func (s *SheetService) GetDayExpense(day int) (*domain.DayExpense, error) {
	rp := DayExpenseRangeProvider(strconv.Itoa(day))
	return nil, notImplementedErr
}

func (s *SheetService) GetFixedExpenses() (*domain.FixedExpenses, error) {
	return nil, notImplementedErr
}

func (s *SheetService) GetMiscExpenses() (*domain.MiscExpenses, error) {
	return nil, notImplementedErr
}

func (s *SheetService) SetDayExpense(exp *domain.Expense) error {
	return notImplementedErr
}

type service struct {
	srvc *sheets.Service
	sheetInfo sheetInfo
}

func (s *service) set(range_ string, sheetId string, valueRange *sheets.ValueRange) error {
	sheetUpdateCall := s.srvc.Spreadsheets.Values.Update(sheetId, range_, valueRange).ValueInputOption(VALUE_OPTION_USER_ENTERED)
	_, err := sheetUpdateCall.Do()
	if err != nil {
		return err
	}
	return nil
}

func (s *service) get(range_ string, sheetId string) (*sheets.ValueRange, error) {
	return s.srvc.Spreadsheets.Values.Get(sheetId, range_).Do()
}

func (s *service) getValue(range_ string) (*sheets.ValueRange, error) {
	return s.get(range_, s.sheetInfo.sheetId)
}

func (s *service) setValue(range_ string, valueRange *sheets.ValueRange) error {
	return s.set(range_, s.sheetInfo.sheetId, valueRange)
}

func GetDayExpense(getter GetSheetIDProvider, rp SheetRangeProvider) (*domain.DayExpense, error) {
	vr, err := getter.Get(rp.ProvideAddress(), getter.ProvideSheetID())
	if err != nil {
		return nil, err
	}
	if vr.HTTPStatusCode != 200 {
		return nil, errors.New("sheets.v2: invalid response from the server")
	}
	dayExpense := new(dayExpense)
	dayExpense.parse(vr.Values)
	dde := dayExpense.toDayExpense()
	return dde, nil
}

func SetDayExpense(setter SetSheetIDProvider, rp SheetRangeProvider, valueRange *sheets.ValueRange) error {
	return setter.Set(rp.ProvideAddress(), setter.ProvideSheetID(), valueRange)
}
