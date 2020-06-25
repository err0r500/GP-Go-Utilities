package sns

import (
	"fmt"
	"strings"

	"github.com/VoodooTeam/GP-Go-Utilities/logger"
	"github.com/aws/aws-sdk-go/service/sns"
)

func getTopicOrCreate(svc *sns.SNS, topicName *string) (*string, error) {
	result, err := svc.ListTopics(nil)
	if err != nil {
		logger.Error(fmt.Sprintf("error on `getTopicOrCreate`; error=%s;", err.Error()))
		return nil, err
	}

	var arnTopic *string
	for _, t := range result.Topics {
		if strings.Contains(*t.TopicArn, *topicName) {
			arnTopic = t.TopicArn
		}
	}

	if arnTopic == nil {
		return createTopic(svc, topicName)
	}

	return arnTopic, nil
}

func createTopic(svc *sns.SNS, topicName *string) (*string, error) {
	resultURL, err := svc.CreateTopic(&sns.CreateTopicInput{
		Name: topicName,
	})

	return resultURL.TopicArn, err
}
