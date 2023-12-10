package main

type config struct {
	port string

	sender_mail string
}

type application struct {
	config config
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
