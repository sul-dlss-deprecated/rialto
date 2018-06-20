package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// Handler is the Lambda function handler
func Handler(ctx context.Context, snsEvent events.SNSEvent) {
	// stdout and stderr are sent to AWS CloudWatch Logs
	log.Printf("Running Lambda function\n")

	host := os.Getenv("ES_HOST")
	url := fmt.Sprintf("http://%s:4571/records/item", host)

	for _, record := range snsEvent.Records {
		// snsRecord := record.SNS
		// str = fmt.Sprintf("[%s %s] Message = %s \n", record.EventSource, snsRecord.Timestamp, snsRecord.Message)
		message := fmt.Sprintf("{\"foo\": \"bar%s\"}", record.SNS.Message)
		reader := strings.NewReader(message)
		resp, err := http.Post(url, "application/json", reader)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()
	}
}

func main() {
	lambda.Start(Handler)
}
