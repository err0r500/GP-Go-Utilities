package sqs

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/sqs"
)

func getQueueURLOrCreate(svc *sqs.SQS, queueName *string) (*string, error) {
	resultURL, err := svc.GetQueueUrl(&sqs.GetQueueUrlInput{
		QueueName: queueName,
	})

	if err != nil {
		if aerr, ok := err.(awserr.Error); ok && aerr.Code() == sqs.ErrCodeQueueDoesNotExist {
			return createQueue(svc, queueName)
		}
		return nil, err
	}

	return resultURL.QueueUrl, nil
}

func createQueue(svc *sqs.SQS, queueName *string) (*string, error) {
	resultURL, err := svc.CreateQueue(&sqs.CreateQueueInput{
		QueueName: queueName,
		Attributes: map[string]*string{
			"DelaySeconds":           aws.String("0"),
			"MessageRetentionPeriod": aws.String("259200"), // 3 days
		},
	})

	return resultURL.QueueUrl, err
}
