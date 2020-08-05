package sns

import (
	"github.com/VoodooTeam/GP-Go-Utilities/logger"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sns"
)

type SmsSender struct {
	svc *SNS
}

func NewSmsSender(svc *SNS) *SmsSender {
	return &SmsSender{
		svc: svc,
	}
}

func (s *SmsSender) SendSMS(message string, phoneNumber string) error {
	params := &sns.PublishInput{
		Message:     aws.String(message),
		PhoneNumber: aws.String(phoneNumber),
	}
	_, err := s.svc.client.Publish(params)

	if err != nil {
		logger.Errorf("Error on SendSMS: %s;", err.Error())
		return err
	}

	return nil
}
