package main

import (
	"net/http"
)

func healthCheck(w http.ResponseWriter, r *http.Request) {

	res := envelope{"status": "ok"}

	err := writeJSON(w, http.StatusOK, res, nil)
	if err != nil {
		serverErrorResponse(w, r, err)
	}

}
