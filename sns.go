package main

import (
	"context"
	"fmt"
	"log"
	"time"

	awsconfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sns"
)

func isInList(numberToCheck string, arrayOfNumbers []string) error {
	for _, number := range arrayOfNumbers {
		if number == numberToCheck {
			return nil
		}
	}
	return fmt.Errorf("number not in allowed list")
}

var totalNumberIWant int32 = 50

func (app *application) sendSms(recipient string, message string) error {

	cfg, err := awsconfig.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatal(err.Error())

	}

	service := sns.NewFromConfig(cfg)

	input := &sns.ListSMSSandboxPhoneNumbersInput{
		MaxResults: &totalNumberIWant,
	}

	res, err := service.ListSMSSandboxPhoneNumbers(context.TODO(), input)
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	var numbers []string

	for _, number := range res.PhoneNumbers {
		numbers = append(numbers, *number.PhoneNumber)
	}

	err = isInList(recipient, numbers)
	if err != nil {

		logger.Error("number not in list")
		return err
	}

	publishInput := &sns.PublishInput{
		Message:     &message,
		PhoneNumber: &recipient,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := service.Publish(ctx, publishInput)
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	if result.MessageId != nil {
		logger.Info("Message sent successfully.", " Message ID:", *result.MessageId)
	} else {
		logger.Info("err in sending message", "error ", err.Error())
		return err
	}

	return nil
}
