package main

import (
	"net/http"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func (app *application) healthCheck(w http.ResponseWriter, r *http.Request) {

	res := envelope{"status": "ok"}

	err := writeJSON(w, http.StatusOK, res, nil)
	if err != nil {
		serverErrorResponse(w, r, err)
	}

}

func (app *application) sendMail(w http.ResponseWriter, r *http.Request) {

	var input struct {
		DestinationEmail string `json:"destinationEmail" validate:"required,email"`
		Subject          string `json:"subject" validate:"required"`
		Body             string `json:"body" validate:"required"`
	}

	err := readJSON(w, r, &input)
	if err != nil {
		badRequestResponse(w, r, err)
		return
	}

	validator := validator.New()

	err = validator.Struct(input)
	if err != nil {
		badRequestResponse(w, r, err)
		return
	}

	destMail := input.DestinationEmail
	subject := input.Subject
	body := input.Body

	err = app.sendEmail(destMail, subject, body)
	if err != nil {
		badRequestResponse(w, r, err)
		return
	}

	data := envelope{
		"success": true,
		"message": "email send successfully",
	}

	err = writeJSON(w, http.StatusOK, data, nil)
	if err != nil {
		serverErrorResponse(w, r, err)
	}
}
