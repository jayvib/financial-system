package sheets

import "fmt"

func parseToRange(sheetName, addressRange string) string {
	return fmt.Sprintf("%s!%s", sheetName, addressRange)
}
