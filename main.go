package main

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/sesv2"
)

func main() { // Using the SDK's default configuration, loading additional config
	// and credentials values from the environment variables, shared
	// credentials, and shared configuration files
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-east-1"))
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	s3Client := s3.NewFromConfig(cfg)
	sesClient := sesv2.NewFromConfig(cfg)

	// Using the Config value, create the DynamoDB client
	// svc := s3
	// svc := dynamodb.NewFromConfig(cfg)

	// Build the request with its input parameters
	// resp, err := svc.ListTables(context.TODO(), &dynamodb.ListTablesInput{
	// 	Limit: aws.Int32(5),
	// })
	// if err != nil {
	// 	log.Fatalf("failed to list tables, %v", err)
	// }

	// fmt.Println("Tables:")
	// for _, tableName := range resp.TableNames {
	// 	fmt.Println(tableName)
	// }
}
