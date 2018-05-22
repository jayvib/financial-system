package convert

import "errors"

func InterfaceToString(data interface{}) (string, error) {
	v, ok := data.(string)
	if !ok {
		return "", errors.New("cannot convert non-string type in interface value")
	}
	return v, nil
}
