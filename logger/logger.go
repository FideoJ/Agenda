package logger

import (
	"log"
	"os"
)

var (
	infoLogger *log.Logger
	errLogger  *log.Logger
)

func init() {
	infoLogger = log.New(os.Stdout, "[INFO] ", log.Ldate|log.Ltime|log.Lshortfile)
	errLogger = log.New(os.Stderr, "[ERROR] ", log.Ldate|log.Ltime|log.Lshortfile)
}

func Info(format string, v ...interface{}) {
	infoLogger.Printf(format, v...)
	len := len(format)
	if len > 0 && format[len-1] != '\n' {
		infoLogger.Print("\n")
	}
}

func ErrIf(err error) {
	if err != nil {
		errLogger.Println(err)
	}
}

func FatalIf(err error) {
	if err != nil {
		errLogger.Fatalln(err)
	}
}
