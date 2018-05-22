package sheets

import (
	"errors"
	"sync"
	"os"
	"fmt"
	"encoding/csv"
	"io"
	"google.golang.org/api/sheets/v4"
)

const sheetIDPath = "sheetids.csv" // current directory

var (
	notImplementedError = errors.New("not implemented yet")
)

type cacheSheetIDs map[string]string

func (id cacheSheetIDs) set(month, sheetId string) {
	id[month] = sheetId
}

func (id cacheSheetIDs) get(filename string) (sheetId string, err error) {
	sheetId, ok := id[filename]
	if !ok {
		return "", errors.New("not found filename's spreadsheet ID")
	}
	return sheetId, nil
}

type CacheSheetIDs struct {
	mutex         sync.Mutex
	cacheSheetIds cacheSheetIDs
	sheetIdPath   string
}

func (s *CacheSheetIDs) SetSheetId(mo, id string) error {
	if s.cacheSheetIds == nil {
		s.cacheSheetIds = make(cacheSheetIDs)
	}

	err := s.storeSheetId(mo, id)
	if err != nil {
		return err
	}

	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.cacheSheetIds.set(mo, id)

	return nil
}

func (s *CacheSheetIDs) GetSheetId(filename string) string {
	id, _ := s.cacheSheetIds.get(filename)
	return id
}

func (s *CacheSheetIDs) loadSheetIds() error {
	file, err := os.OpenFile(sheetIDPath, os.O_APPEND|os.O_RDWR, 0664)
	if err != nil {
		if os.IsNotExist(err) {
			return errors.New(fmt.Sprintf("sheet: %s is not yet created", s.sheetIdPath))
		}
		return err
	}
	defer file.Close()

	s.cacheSheetIds = make(cacheSheetIDs)

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

		s.cacheSheetIds[rows[monthIndex]] = rows[sheetIdIndex]
	}

	return nil
}

func (s *CacheSheetIDs) storeSheetId(mo, id string) error {

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

func (s *CacheSheetIDs) GetAllSheetIDs() map[string]string {
	return s.cacheSheetIds
}

type SpreadSheet struct {
	spreadSheet *sheets.Spreadsheet
}

func LoadSheetIDs() (*CacheSheetIDs, error) {
	sheetIds := &CacheSheetIDs{
		cacheSheetIds: (cacheSheetIDs)(nil),
		sheetIdPath: sheetIDPath,
	}

	err := sheetIds.loadSheetIds()
	if err != nil {
		return nil, err
	}

	return sheetIds, nil
}