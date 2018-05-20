package sheets

import (
	"net/http"
	"google.golang.org/api/sheets/v4"
	"github.com/pkg/errors"
)


type Service struct {
	service       *sheets.Service
	title, sheetId string
}

func New(client *http.Client, title, sheetId string) (service *Service, err error) {
	servc, err := sheets.New(client)
	if err != nil {
		return nil, err
	}

	service =  &Service{
		service: servc,
		title: title,
		sheetId: sheetId,
	}
	return service, nil
}

func (s *Service) getValues(_range string) ([][]interface{}, error) {
	resp, err := s.service.Spreadsheets.Values.Get(s.sheetId, _range).Do()
	if err != nil {
		return nil, err
	}

	if resp.HTTPStatusCode != http.StatusOK {
		return nil, errors.New("sheets: error when communicating the google api")
	}

	if len(resp.Values) == 0 {
		return nil, errors.New("sheets: no data found")
	}

	return resp.Values, nil
}

func (s *Service) GetValues(range_ string) ([][]interface{}, error) {
	return s.getValues(range_)
}

func (s *Service) SetValues(range_ string, vr *sheets.ValueRange) error {
	sheetUpdateCall := s.service.Spreadsheets.Values.Update(s.sheetId, range_, vr)
	sheetUpdateCall = sheetUpdateCall.ValueInputOption("USER_ENTERED")
	_, err := sheetUpdateCall.Do()
	if err != nil {
		return err
	}
	return nil
}

