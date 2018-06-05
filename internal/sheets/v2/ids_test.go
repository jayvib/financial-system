package sheets_test

import (
	"testing"
	. "financial-system/internal/sheets/v2"
	"log"
	"fmt"
)

func TestLoadSheetIDs(t *testing.T) {
	ids, err := LoadSheetIDs()
	if err != nil {
		log.Fatal(err)
	}

	if ids == nil {
		t.Error("sheet ids must not be empty")
	}

	for month, id := range ids {
		fmt.Printf("%s %s\n", month, id)
	}
}
