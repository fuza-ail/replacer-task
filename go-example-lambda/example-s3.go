package main

import (
	"context"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(ctx context.Context, s3Event events.S3Event) error {
	for _, record := range s3Event.Records {
		s3 := record.S3
		log.Printf("Bucket = %s, Key = %s", s3.Bucket.Name, s3.Object.Key)
	}
	return nil
}

func main() {
	lambda.Start(handler)
}
