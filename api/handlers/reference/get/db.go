package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

// Declare a new DynamoDB instance. Note that this is safe for concurrent
// use.
var db = dynamodb.New(session.New(), aws.NewConfig().WithRegion("us-east-1"))

func GetItem(a string) (*aggregate_reference, error) {
	// Prepare the input for the query.
	input := &dynamodb.GetItemInput{
		TableName: aws.String("AggregateReference"),
		Key: map[string]*dynamodb.AttributeValue{
			"AggregatorId": {
				S: aws.String(a),
			},
		},
	}

	// return item from dynamodb and if no matching items is found
	// then return nil
	result, err := db.GetItem(input)
	if err != nil {
		return nil, err
	}
	if result.Item == nil {
		return nil, nil
	}

	// the retuning object has an underlying data type of a mapped attribute string
	// The UnMarshalMap helper can be used to parse the fields straight into fields of
	// a struct.

	agg := new(aggregate_reference)
	err = dynamodbattribute.UnmarshalMap(result.Item, agg)
	if err != nil {
		return nil, err
	}

	return agg, nil
}
