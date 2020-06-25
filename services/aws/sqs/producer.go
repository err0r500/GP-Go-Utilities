package sqs

import (
	"encoding/json"
	"fmt"

	"github.com/VoodooTeam/GP-Go-Utilities/logger"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
)

type Producer struct {
	svc       *SQS
	queueName string
	queueURL  *string
}

// NewProducer create a new sqs producer
func NewProducer(svc *SQS, queueName string) *Producer {
	url, err := getQueueURLOrCreate(svc.client, &queueName)
	if err != nil {
		logger.Error("NewConsumer - ", err.Error())
		return nil
	}

	return &Producer{
		svc:       svc,
		queueName: queueName,
		queueURL:  url,
	}
}

func (p *Producer) SendMessageJSONMarshal(msgType string, msgContent interface{}) error {
	msg, err := json.Marshal(msgContent)
	if err != nil {
		logger.Error(fmt.Sprintf("error on Producer.SendMessageJSONMarshal `json.Marshal` message=%s; error=%s;", fmt.Sprint(msg), err.Error()))
		return err
	}

	p.SendMessage(string(msg), msgType)
	return nil
}

func (p *Producer) SendMessage(message string, msgType string) error {
	_, err := p.svc.client.SendMessage(&sqs.SendMessageInput{
		QueueUrl:    p.queueURL,
		MessageBody: aws.String(message),
		MessageAttributes: map[string]*sqs.MessageAttributeValue{
			"type": &sqs.MessageAttributeValue{
				DataType:    aws.String("String"),
				StringValue: aws.String(msgType),
			},
		},
	})

	if err != nil {
		logger.Error(fmt.Sprintf("error on Producer.SendMessage queueName=%s; error=%s;", p.queueName, err.Error()))
		return err
	}

	return nil
}
