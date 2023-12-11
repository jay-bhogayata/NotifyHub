package main

import (
	"log/slog"
	"os"
)

var logger *slog.Logger

func (app *application) LoggerInit() {
	logFile, err := os.OpenFile("server.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0777)
	if err != nil {
		panic("cannot open log file")
	}

	app.logFile = logFile

	logger = slog.New(slog.NewJSONHandler(logFile, nil))
}
