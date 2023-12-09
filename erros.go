package main

import "net/http"

func errResponse(w http.ResponseWriter, r *http.Request, status int, message string) {
	res := envelope{"error": message}
	err := writeJSON(w, status, res, nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func serverErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	logger.Error("internal server error", "message", err)
	message := "the server encountered a problem and could not process your request"
	errResponse(w, r, http.StatusInternalServerError, message)
}

func badRequestResponse(w http.ResponseWriter, r *http.Request, err error) {
	errResponse(w, r, http.StatusBadRequest, err.Error())
}
