package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	_ "github.com/jay-bhogayata/notifyHub/docs"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

func (app *application) setServer() *http.Server {

	r := chi.NewRouter()

	r.Use(Logging)
	r.Use(middleware.SetHeader("Access-Control-Allow-Origin", "*"))
	r.Use(middleware.SetHeader("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE"))
	r.Use(middleware.SetHeader("Access-Control-Allow-Headers", "Content-Type"))

	apiRouter := chi.NewRouter()
	r.Mount("/api/v1", apiRouter)

	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/doc.json"),
	))

	apiRouter.Get("/health", app.healthCheck)
	apiRouter.Post("/sendmail", app.sendMail)
	apiRouter.Post("/sendsms", app.sendSMS)

	server := &http.Server{
		Addr:         fmt.Sprintf(":%v", app.config.port),
		Handler:      r,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	return server
}

func (app *application) ServerInit() {
	server := app.setServer()

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	go func() {
		logger.Info("server is starting on", "port", app.config.port)
		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			logger.Error("could not listen on", "port", app.config.port, "error", err)
			os.Exit(1)
		}

	}()

	<-shutdown

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		logger.Error("could not gracefully shutdown the server", "error", err)
		os.Exit(1)
	}

	logger.Info("server stopped gracefully")
}
