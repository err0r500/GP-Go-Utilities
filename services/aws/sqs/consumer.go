package sqs

import (
	"fmt"

	"github.com/VoodooTeam/GP-Go-Utilities/logger"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
)

type Consumer struct {
	svc       *SQS
	queueName string
	queueURL  *string
	messages  chan<- *Message

	autoDel bool
}

const (
	sqsMaxMessages     = int64(1)
	sqsPollWaitSeconds = int64(10)
)

// NewConsumer create a new sqs consumer
//
// Sample of how to use it:
//
// chnMessages := make(chan *sqs.Message, sqsMaxMessages)
//
// consumer := NewConsumer(..., chnMessages)
// go consumer.Consume()
//
// for message := range chnMessages {
// 		handleMessage(message)
//		consumer.Delete(message)
// }
//
func NewConsumer(svc *SQS, queueName string, messages chan<- *Message) *Consumer {
	url, err := getQueueURLOrCreate(svc.client, &queueName)
	if err != nil {
		logger.Error("NewConsumer - ", err.Error())
		return nil
	}

	return &Consumer{
		svc:       svc,
		queueName: queueName,
		queueURL:  url,
		messages:  messages,
		autoDel:   false,
	}
}

func (c *Consumer) Consume() {
	logger.Debug("New consumer start working")
	c.worker()
}

func (c *Consumer) worker() {
	for {
		output, err := retrieveSQSMessages(c.svc, c.queueURL)
		if err != nil {
			continue
		}

		for _, message := range output.Messages {
			msg := NewMessage(message)
			c.messages <- msg
			if c.autoDel {
				c.Delete(msg)
			}
		}
	}
}

func retrieveSQSMessages(svc *SQS, queueURL *string) (*sqs.ReceiveMessageOutput, error) {
	return svc.client.ReceiveMessage(&sqs.ReceiveMessageInput{
		QueueUrl:            queueURL,
		MaxNumberOfMessages: aws.Int64(sqsMaxMessages),
		WaitTimeSeconds:     aws.Int64(sqsPollWaitSeconds),
	})
}

func (c *Consumer) Delete(msg *Message) error {
	_, err := c.svc.client.DeleteMessage(&sqs.DeleteMessageInput{
		QueueUrl:      c.queueURL,
		ReceiptHandle: msg.message.ReceiptHandle,
	})
	if err != nil {
		logger.Error(fmt.Sprintf("error on `Delete(msg *Message)`; queue=%s; error=%s;", c.queueName, err.Error()))
	}
	return err
}
