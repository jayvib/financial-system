package sheets

import (
	"google.golang.org/api/sheets/v4"
	"github.com/pkg/errors"
)

var notImplementedErr = errors.New("not implemented yet")

type Setter interface {
	Set(range_ string, sheetId string, valueRange *sheets.ValueRange) error
}

type Getter interface {
	Get(range_ string, sheetId string) (*sheets.ValueRange, error)
}

type GetSheetIDProvider interface {
	SheetIDProvider
	Getter
}

type SetSheetIDProvider interface {
	SheetIDProvider
	Setter
}

type SetGetter interface {
	Setter
	Getter
}

type SheetIDProvider interface {
	ProvideSheetID() string
}

type SheetRangeProvider interface {
	ProvideAddress() string
}

type neutralizer interface {
	Neutralize(prefix string, data interface{})
}