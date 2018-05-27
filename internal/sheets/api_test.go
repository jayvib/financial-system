package sheets

import (
	"financial-system/config"
	"fmt"
	"testing"
)

func setup() (*Service, error) {
	title := "february"
	id := "1LBUAV99tHgD7RpdMekiDwVRE5A29B7L49wpHdJAmPSA"

	client, err := config.DefaultClient()
	if err != nil {
		return nil, err
	}

	srv, err := New(client, title, id)
	if err != nil {
		return nil, err
	}

	return srv, nil
}

func TestGetDailyExpense(t *testing.T) {
	srv, err := setup()
	if err != nil {
		t.Fatal(err)
	}
	exp, err := srv.GetDayExpense(FirstDay)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Print(exp)
}
