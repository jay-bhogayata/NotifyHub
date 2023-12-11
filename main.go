package main

import "os"

type config struct {
	port string

	sender_mail   string
	sns_topic_arn string
}

type application struct {
	config  config
	logFile *os.File
}

func main() {

	var cfg config

	app := &application{
		config: cfg,
	}

	app.LoggerInit()
	app.config.LoadConfig()
	app.ServerInit()
}
