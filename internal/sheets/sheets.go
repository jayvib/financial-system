package sheets

import (
	"errors"
	"sync"
	"os"
	"fmt"
	"encoding/csv"
	"io"
)

const sheetIDPath = "./sheetids.csv" // current directory

var (
	notImplementedError = errors.New("not implemented yet")
)

type monthSheetId map[string]string

func (id monthSheetId) set(month, sheetId string) {
	id[month] = sheetId
}

func (id monthSheetId) get(filename string) (sheetId string, err error) {
	sheetId, ok := id[filename]
	if !ok {
		return "", errors.New("not found filename's spreadsheet ID")
	}
	return sheetId, nil
}

type MonthSheetID struct {
	mutex sync.Mutex
	monthSheetId monthSheetId
	sheetIdPath string
}

func (s *MonthSheetID) SetSheetId(mo, id string) error {
	if s.monthSheetId == nil {
		s.monthSheetId = make(monthSheetId)
	}

	err := s.storeSheetId(mo, id)
	if err != nil {
		return err
	}

	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.monthSheetId.set(mo, id)

	return nil
}

func (s *MonthSheetID) GetSheetId(filename string) string {
	id, _ := s.monthSheetId.get(filename)
	return id
}

func (s *MonthSheetID) loadSheetIds() error {
	file, err := os.OpenFile(sheetIDPath, os.O_APPEND|os.O_RDWR, 0664)
	if err != nil {
		if os.IsNotExist(err) {
			return errors.New(fmt.Sprintf("sheet: %s is not yet created", s.sheetIdPath))
		}
		return err
	}
	defer file.Close()

	s.monthSheetId = make(monthSheetId)

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1

	monthIndex := 0
	sheetIdIndex := 1

	for {
		rows, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}

		s.monthSheetId[rows[monthIndex]] = rows[sheetIdIndex]
	}

	return nil
}

func (s *MonthSheetID) storeSheetId(mo, id string) error {

	file, err := os.OpenFile(s.sheetIdPath, os.O_APPEND|os.O_RDWR, 0664)
	if err != nil {
		if os.IsNotExist(err) {
			file, err = os.OpenFile(s.sheetIdPath, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0600)
			if err != nil {
				return err
			}
		}
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	err = writer.Write([]string{mo, id})
	if err != nil {
		return err
	}
	writer.Flush()

	return nil
}

