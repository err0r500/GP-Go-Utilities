package sns

import (
	"encoding/json"
	"fmt"

	"github.com/VoodooTeam/GP-Go-Utilities/logger"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sns"
)

type Producer struct {
	svc       *SNS
	topicName string
	topicARN  *string
}

// NewProducer create a new sqs producer
func NewProducer(svc *SNS, topicName string) *Producer {
	topicARN, err := getTopicOrCreate(svc.client, &topicName)
	if err != nil {
		logger.Error("NewProducer - ", err.Error())
		return nil
	}

	return &Producer{
		svc:       svc,
		topicName: topicName,
		topicARN:  topicARN,
	}
}

func (p *Producer) SendMessageJSONMarshal(msgType string, msgContent interface{}) error {
	msg, err := json.Marshal(msgContent)
	if err != nil {
		logger.Error(fmt.Sprintf("error on Producer.SendMessageJSONMarshal `json.Marshal` message=%s; error=%s;", fmt.Sprint(msg), err.Error()))
		return err
	}

	return p.SendMessage(msgType, string(msg))
}

func (p *Producer) SendMessage(msgType string, message string) error {

	_, err := p.svc.client.Publish(&sns.PublishInput{
		TopicArn: p.topicARN,
		Message:  aws.String(message),
		MessageAttributes: map[string]*sns.MessageAttributeValue{
			"type": &sns.MessageAttributeValue{
				DataType:    aws.String("String"),
				StringValue: aws.String(msgType),
			},
		},
	})

	if err != nil {
		logger.Error(fmt.Sprintf("error on SendMessage topic=%s; error=%s;", p.topicName, err.Error()))
		return err
	}

	return nil
}
