package main

import (
	"context"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/config"
)

// Handler handles the CloudWatch Event
func Handler(ctx context.Context, event events.CloudWatchEvent) {

	//log.Println(string(event.Detail))

	_, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		log.Fatalf("Unable to load SDK config, %v", err)
	}

	log.Println("Hello World!")

}

func main() {
	lambda.Start(Handler)
}
