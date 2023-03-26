package logger

import (
	"log"
	"os"
	"path/filepath"
	"time"
)

func CreateLogger() (*log.Logger, error) {
	logDir := "./logs"
	if _, err := os.Stat(logDir); os.IsNotExist(err) {
		if err := os.Mkdir(logDir, 0755); err != nil {
			return nil, err
		}
	}

	filename := time.Now().Format("20060102") + ".log"
	logfile, err := os.OpenFile(filepath.Join(logDir, filename), os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}

	logger := log.New(logfile, "", log.LstdFlags)
	return logger, nil
}

func Write(message string) error {
	logger, err := CreateLogger()
	if err != nil {
		return err
	}
	defer logger.Writer().(*os.File).Close()

	logger.Println(message)
	return nil
}
