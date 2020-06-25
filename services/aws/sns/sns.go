package sns

import (
	"github.com/aws/aws-sdk-go/aws/session"
	SDK "github.com/aws/aws-sdk-go/service/sns"
)

const (
	serviceName = "SNS"
)

// SQS has SQS client and Queue list.
type SNS struct {
	client *SDK.SNS
}

// New returns initialized *SQS.
func New(sess *session.Session) *SNS {
	return &SNS{
		client: SDK.New(sess),
	}
}
