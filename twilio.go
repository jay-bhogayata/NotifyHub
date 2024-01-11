package main

import (
	"encoding/json"
	"fmt"

	"github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/api/v2010"
)

func (cfg *config) SendSMSUsingTwilio(body string, recipient string) error {
	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: cfg.tw_acc_sid,
		Password: cfg.tw_auth_token,
	})

	params := &twilioApi.CreateMessageParams{}
	params.SetTo(recipient)
	params.SetFrom(cfg.tw_phone_number)
	params.SetBody(body)

	res, err := client.Api.CreateMessage(params)
	if err != nil {
		logger.Error(err.Error())
		return err
	} else {
		response, _ := json.Marshal(*res)
		logger.Info(fmt.Sprintf("Twilio response: %s", string(response)))
	}
	return nil
}
