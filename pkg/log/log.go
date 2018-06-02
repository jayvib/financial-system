package log

import (
	"log"
	"io"
)

var (
	Debug *log.Logger
	Info  *log.Logger
	Error *log.Logger
	InfoHandler io.Writer
)

func InitLog(
	traceFilename string,
	debugHandler io.Writer,
	infoHandler io.Writer,
	errorHandler io.Writer,
) {
	
}