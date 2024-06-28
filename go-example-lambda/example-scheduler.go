package main

import (
	"context"
	"log"

	"github.com/aws/aws-lambda-go/lambda"
)

func handler(ctx context.Context) error {
	log.Println("Scheduled task executed")
	return nil
}

func main() {
	lambda.Start(handler)
}
