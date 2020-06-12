package services

import (
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)

// AWSSession builder
func AWSSession() (*session.Session, error) {
	local := os.Getenv("APP_USER") == "air"

	var sess *session.Session
	var err error

	if local {
		log.Println("Run localy")
		sess, err = session.NewSession(&aws.Config{
			Region:      aws.String("us-west-2"),
			Credentials: credentials.NewSharedCredentials("", os.Getenv("AWS_PROFILE")),
		})
	} else {
		sess, err = session.NewSession(&aws.Config{
			Region: aws.String("us-west-2"),
		})
	}

	if err != nil {
		log.Fatal(err.Error())
	}
	_, err = sess.Config.Credentials.Get()
	if err != nil {
		log.Fatal(err.Error())
	}

	return sess, err
}
