package main

import (
	"github.com/aws/aws-lambda-go/lambda"
)

type aggregate_reference struct {
	AggregatorId   string `json:"AggregatorId"`
	AggregatorName string `json:"AggregatorName"`
	SystemProvider string `json:"SystemProvider"`
	AggregatorAlt  string `json:"AggregatorAlt"`
	// really should add a created at timestamp to this too
}

// output function is the lambda handler
func output() (*aggregate_reference, error) {
	a, err := GetItem("1")
	if err != nil {
		return nil, err
	}

	return a, nil
}

func main() {
	lambda.Start(output)
}
