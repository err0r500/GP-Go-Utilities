package sqs

import (
	"github.com/aws/aws-sdk-go/aws/session"
	SDK "github.com/aws/aws-sdk-go/service/sqs"
)

const (
	serviceName = "SQS"
)

// SQS has SQS client and Queue list.
type SQS struct {
	client *SDK.SQS
}

// New returns initialized *SQS.
func New(sess *session.Session) *SQS {
	return &SQS{
		client: SDK.New(sess),
	}
}
