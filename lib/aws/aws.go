package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"log"
)

type AWS struct {
	session *session.Session
}

func New() *AWS {
	sess, err := session.NewSession()
	if err != nil {
		log.Fatal("failed to create session,", err)
	}

	return &AWS{
		session: sess,
	}
}

func (aws *AWS) Session() *session.Session {
	return aws.session
}
