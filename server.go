package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

var port int = 8080

func setServer() *http.Server {

	r := chi.NewRouter()

	r.Use(middleware.Logger)

	apiRouter := chi.NewRouter()
	r.Mount("/api/v1", apiRouter)

	apiRouter.Get("/health", healthCheck)

	server := &http.Server{
		Addr:         fmt.Sprintf(":%v", port),
		Handler:      r,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	return server
}

func ServerInit() {
	server := setServer()

	logger.Info("server is starting on", "port", port)
	err := server.ListenAndServe()
	if err != nil {
		logger.Error("error starting server", "error", err)
	}
}
