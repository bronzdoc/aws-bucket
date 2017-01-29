package s3

import (
	"fmt"
	s3_sdk "github.com/aws/aws-sdk-go/service/s3"
	"github.com/bronzdoc/bucket/lib/aws"
	"log"
)

type S3 struct {
	sdk *s3_sdk.S3
}

func New() *S3 {
	aws := aws.New()
	s3 := s3_sdk.New(aws.Session())

	return &S3{
		sdk: s3,
	}
}

func (s3 *S3) ListBuckets() {
	var params *s3_sdk.ListBucketsInput
	resp, err := s3.sdk.ListBuckets(params)

	if err != nil {
		log.Println("ListBuckets() -", err.Error())
		return
	}

	fmt.Println(resp)
}
