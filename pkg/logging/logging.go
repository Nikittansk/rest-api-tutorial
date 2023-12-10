package logging

import (
	"log"
	"os"
)

type Logging struct {
	InfoLog  *log.Logger
	ErrorLog *log.Logger
}

func NewLogging(filePath string) (*Logging, error) {

	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return nil, err
	}

	return &Logging{
		InfoLog:  log.New(file, "INFO\t", log.Ldate|log.Ltime),
		ErrorLog: log.New(file, "ERROR\t", log.Ldate|log.Ltime|log.Llongfile),
	}, nil
}
