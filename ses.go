package main

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/ses"
	"github.com/aws/aws-sdk-go-v2/service/ses/types"
)

func (app *application) sendEmail(dest string, subject string, body string) error {

	var destinations []string

	destinations = append(destinations, dest)

	service := ses.NewFromConfig(app.config.awsConfig)
	input := &ses.SendEmailInput{
		Destination: &types.Destination{
			ToAddresses: []string{
				destinations[0],
			},
		},
		Message: &types.Message{
			Body: &types.Body{
				Text: &types.Content{
					Data: &body,
				},
			},
			Subject: &types.Content{
				Data: &subject,
			},
		},
		Source: &app.config.sender_email,
	}

	_, err := service.SendEmail(context.Background(), input)
	if err != nil {
		logger.Error("could not send email", "error", err)
	}

	return nil
}
