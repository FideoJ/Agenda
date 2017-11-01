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
	infoLogger = log.New(os.Stdout, "[INFO] ", log.Ldate|log.Ltime)
	errLogger = log.New(os.Stderr, "[ERROR] ", log.Ldate|log.Ltime)
}

func Info(format string, v ...interface{}) {
	infoLogger.Printf(format, v...)
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
