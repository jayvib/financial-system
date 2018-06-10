package sheets

import (
	"google.golang.org/api/sheets/v4"
)


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

type sheetRanger interface {
	rangeAddress() string
}

type neutralizer interface {
	Neutralize(prefix string, data interface{})
}