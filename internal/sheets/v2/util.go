package sheets

import "strings"

func neutralizeString(prefix string, data interface{}) string {
	str := data.(string)
	return strings.Replace(str, prefix, "", -1)
}
