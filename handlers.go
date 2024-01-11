package main

import (
	"net/http"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

type response struct {
	Message string `json:"message"`
}

// @Summary      Health Check
// @Description  This endpoint is used to check the health of the service
// @Produce      json
// @Success      200 {object} response
// @Failure      400
// @Failure      404
// @Failure      500
// @Router       /health [get]
func (app *application) healthCheck(w http.ResponseWriter, r *http.Request) {

	res := envelope{"status": "ok"}

	err := writeJSON(w, http.StatusOK, res, nil)
	if err != nil {
		serverErrorResponse(w, r, err)
	}

}

type mail_input struct {
	DestinationEmail string `json:"destinationEmail" validate:"required,email"`
	Subject          string `json:"subject" validate:"required"`
	Body             string `json:"body" validate:"required"`
}

// @Summary      Send Email
// @Description  This endpoint is used to send email
// @Accept       json
// @Produce      json
// @Param        body body mail_input true "body"
// @Success      200 {object} response
// @Failure      400
// @Failure      404
// @Failure      500
// @Router       /sendmail [post]
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

type sms_input struct {
	Recipient string `json:"recipient" validate:"required,e164"`
	Message   string `json:"message" validate:"required"`
}

// @Summary      Send SMS
// @Description  This endpoint is used to send sms
// @Accept       json
// @Produce      json
// @Param        body body sms_input true "body"
// @Success      200 {object} response
// @Failure      400
// @Failure      404
// @Failure      500
// @Router       /sendsms [post]
func (app *application) sendSMS(w http.ResponseWriter, r *http.Request) {

	var input struct {
		Recipient string `json:"recipient" validate:"required,e164"`
		Message   string `json:"message" validate:"required"`
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

	recipient := input.Recipient
	message := input.Message

	// err = app.sendSms(recipient, message)
	err = app.config.SendSMSUsingTwilio(message, recipient)
	if err != nil {
		badRequestResponse(w, r, err)
		return
	}

	data := envelope{
		"success": true,
		"message": "sms send successfully",
	}

	err = writeJSON(w, http.StatusOK, data, nil)
	if err != nil {
		serverErrorResponse(w, r, err)
	}
}
