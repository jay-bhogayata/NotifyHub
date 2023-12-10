package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (app *application) setServer() *http.Server {

	r := chi.NewRouter()

	r.Use(middleware.Logger)

	apiRouter := chi.NewRouter()
	r.Mount("/api/v1", apiRouter)

	apiRouter.Get("/health", app.healthCheck)
	apiRouter.Post("/sendmail", app.sendMail)
	apiRouter.Post("/sendsms", app.sendSMS)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%v", app.config.port),
		Handler: r,
		// ReadTimeout:  5 * time.Second,
		// WriteTimeout: 5 * time.Second,
	}

	return server
}

func (app *application) ServerInit() {
	server := app.setServer()

	logger.Info("server is starting on", "port", app.config.port)
	err := server.ListenAndServe()
	if err != nil {
		logger.Error("error starting server", "error", err)
	}
}
