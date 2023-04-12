package main

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"log"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func handler(ctx context.Context) (string, error) {
	// Initialize a session that the SDK will use to load
	// credentials from the shared credentials file ~/.aws/credentials
	// and region from the shared configuration file ~/.aws/config.
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-west-1")},
	)

	// Create DynamoDB client
	svc := dynamodb.New(sess)

	// Define the table name
	tableName := "organizations_dev"

	// Define the scan input
	input := &dynamodb.ScanInput{
		TableName: aws.String(tableName),
	}
	log.Println("TEST SUCCESS! :D")

	// Retrieve the items from the table
	result, err := svc.Scan(input)
	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}

	// Print the items
	fmt.Println("Items:")
	for _, i := range result.Items {
		fmt.Println("  ", i)
	}
	return "success", nil
}

func main() {
	lambda.Start(handler)
}
