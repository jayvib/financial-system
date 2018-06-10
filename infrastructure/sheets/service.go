package sheets

import (
	"financial-system/domain"
	"errors"
	ferror "financial-system/errors"
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
	rp := newDayExpenseRanger(strconv.Itoa(day))
	vr, err := s.getValue(rp.rangeAddress())
	if err != nil {
		ferror.Wrapf(err, "finsheet/sheets: fail getting the value")
	}
	de := new(dayExpense)

	err = de.parse(vr.Values)
	if err != nil {
		return nil, ferror.Wrapf(err, "finsheet/sheets: fail parsing the value range into day expense")
	}

	return de.toDayExpense(), nil
}

func (s *SheetService) SetDayExpense(exp *domain.Expense) error {
	vr := newValueRange(exp.GetValue())
	ranger, err := expenseRange(strconv.Itoa(exp.GetDay()), exp.GetName())
	if err != nil {
		return err
	}
	err = s.setValue(ranger.rangeAddress(), vr)
	if err != nil {
		return ferror.Wrapf(err, "finsheet/sheets: fail setting value to the google spreadsheet")
	}
	return nil
}

func (s *SheetService) GetFixedExpenses() (*domain.FixedExpenses, error) {
	return nil, notImplementedErr
}

func (s *SheetService) GetMiscExpenses() (*domain.MiscExpenses, error) {
	return nil, notImplementedErr
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

func getDayExpense(getter GetSheetIDProvider, rp sheetRanger) (*domain.DayExpense, error) {
	vr, err := getter.Get(rp.rangeAddress(), getter.ProvideSheetID())
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

func setDayExpense(setter SetSheetIDProvider, rp sheetRanger, valueRange *sheets.ValueRange) error {
	return setter.Set(rp.rangeAddress(), setter.ProvideSheetID(), valueRange)
}
