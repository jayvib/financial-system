package errors

import "runtime/debug"

type FinsheetError struct {
	Orig error
	Message string
	StackTrace string
	Args map[string]interface{}
}

func (err FinsheetError) Error() string {
	return err.Message
}

func Wrapf(err error, messagf string, msgArgs ...interface{}) FinsheetError {
	return FinsheetError{
		Orig: err,
		Message: messagf,
		StackTrace: string(debug.Stack()),
		Args: make(map[string]interface{}),
	}
}
