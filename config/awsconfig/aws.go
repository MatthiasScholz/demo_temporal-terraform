package awsconfig

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
)

// FIXME NOT WORKING NOTE The way this function is defined it relies on environment variables to provide the AWS credentials.
func LoadConfig(profile string) aws.Config {
	// FIXME passing an explicit profile as an argument,
	//       furthermore the function does not check if the profile exists.
	awsConfig, err := config.LoadDefaultConfig(context.Background(), config.WithSharedConfigProfile(profile))
	if err != nil {
		log.Fatalf("unable to load aws config: %e", err)
	}

	// Check for valid credentials
	creds, err := awsConfig.Credentials.Retrieve(context.Background())
	if err != nil {
		log.Fatalf("unable to retrieve aws credentials, %e", err)
	}
	if !creds.HasKeys() {
		log.Fatal("no aws credentials present")
	}
	if creds.Expired() {
		log.Fatal("given aws credentials expired")
	}

	return awsConfig
}
