package main

import (
	"log"
	"minlog/logger"
	"os"
)

func main() {
	logger.Info("default log")
	logger.SetOptions(logger.WithLevel(logger.DebugLevel))
	logger.Debug("change default log to debug level")

	fd, err := os.OpenFile("test.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalln("create file test.log failed")
	}
	defer fd.Close()

	l := logger.NewLog(logger.WithLevel(logger.InfoLevel),
		logger.WithOutput(fd),
		logger.WithFormatter(&logger.JsonFormatter{IgnoreBasicFields: false}),
	)
	l.Info("custom log with json formatter")
}
