package main

import (
	"context"
	"log"

	awsconfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ses"
	"github.com/aws/aws-sdk-go-v2/service/ses/types"
)

func (app *application) sendEmail(dest string, subject string, body string) error {

	var destinations []string

	destinations = append(destinations, dest)

	cfg, err := awsconfig.LoadDefaultConfig(context.TODO())

	if err != nil {
		log.Fatal(err.Error())

	}

	service := ses.NewFromConfig(cfg)
	input := &ses.SendEmailInput{
		Destination: &types.Destination{
			ToAddresses: []string{
				dest,
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

	_, err = service.SendEmail(context.Background(), input)
	if err != nil {
		logger.Error("could not send email", "error", err)
	}

	return nil
}
