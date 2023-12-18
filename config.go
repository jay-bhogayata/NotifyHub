package main

import (
	"os"
)

func (cfg *config) LoadConfig() {

	cfg.port = os.Getenv("PORT")
	if cfg.port == "" {
		logger.Error("PORT env variable is not set")
		panic("PORT env variable is not set")
	}

	cfg.sender_mail = os.Getenv("SENDER_EMAIL")
	if cfg.sender_mail == "" {
		logger.Error("SENDER_MAIL env variable is not set")
		panic("SENDER_MAIL env variable is not set")
	}

}
