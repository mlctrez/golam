package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"os"
)

func main() {
	// https://docs.aws.amazon.com/lambda/latest/dg/configuration-envvars.html#configuration-envvars-runtime
	if os.Getenv("LAMBDA_TASK_ROOT") != "" {
		lambda.Start(Handle)
	} else {
		// do something else
	}
}

// Handle is a starting point for a go lambda function.
// Generic request and response can be replaced with specific event types from:
//
//	github.com/aws/aws-lambda-go/events
func Handle(ctx context.Context, request interface{}) (response interface{}, err error) {
	fmt.Println("handle", request)
	response = make(map[string]interface{})
	return response, err
}
