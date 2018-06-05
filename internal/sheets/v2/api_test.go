package sheets_test

import (
	"testing"
	. "financial-system/internal/sheets/v2"
)

func TestDefault(t *testing.T) {
	_, err := Default()
	if err != nil {
		t.Fatal(err)
	}
}
