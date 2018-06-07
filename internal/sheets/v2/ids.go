package sheets

import (
	"github.com/pkg/errors"
	"financial-system/config"
	"os"
	"isell/encoding/csv"
	"io"
	"strings"
	"sync"
)

var (
	notImplementedError = errors.New("not implemented yet")
)

type sheetInfo struct {
	title string
	id string
}

type SheetInfoCache struct {
	sheetinfos map[string]sheetInfo
}

type sheetInfoLoader struct {
	load func() map[string]string
	sync.Once
}


var sheetIDs SheetIDs

type SheetIDs map[string]string

func (id SheetIDs) set(month, sheetId string) {
	id[month] = sheetId
}

func (id SheetIDs) get(mo string) (sheetId string, err error) {
	sheetId, ok := id[mo]
	if !ok {
		return "", errors.New("month's spreadsheet ID not exist.")
	}
	return sheetId, nil
}

func (id SheetIDs) Get(month string) (string, error) {
	return id.get(month)
}

func LoadSheetIDs() (SheetIDs, error) {
	if sheetIDs != nil {
		return sheetIDs, nil
	}

	conf, err := config.DefaultConfig()
	if err != nil {
		return nil, err
	}

	file, err := os.Open(conf.Path.SheetIDs)
	if err != nil {
		return nil, err
	}

	csvReader := csv.NewReader(file)
	csvReader.FieldsPerRecord = -1
	ids := make(SheetIDs)

	for {
		line, err := csvReader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			continue
		}
		ids[strings.ToLower(line[0])] = line[1]
	}

	sheetIDs = ids

	return ids, nil
}
