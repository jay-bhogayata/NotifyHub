package main

import (
	"context"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsconfig "github.com/aws/aws-sdk-go-v2/config"
)

func (cfg *config) LoadConfig() {

	cfg.port = os.Getenv("PORT")
	if cfg.port == "" {
		logger.Error("PORT env variable is not set")
		panic("PORT env variable is not set")
	}

	cfg.sender_email = os.Getenv("SENDER_EMAIL")
	if cfg.sender_email == "" {
		logger.Error("SENDER_MAIL env variable is not set")
		panic("SENDER_MAIL env variable is not set")
	}

	cfg.env = os.Getenv("ENV")
	if cfg.env == "" {
		logger.Error("ENV env variable is not set")
		panic("ENV env variable is not set")
	}

}

func (cfg *config) ConfigLocalAws() (aws.Config, error) {
	awsEndpoint := "http://localhost:4566"
	awsRegion := "us-east-1"

	customResolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		if awsEndpoint != "" {
			return aws.Endpoint{
				PartitionID:   "aws",
				URL:           awsEndpoint,
				SigningRegion: awsRegion,
			}, nil
		}

		return aws.Endpoint{}, &aws.EndpointNotFoundError{}
	})
	awsCfg, err := awsconfig.LoadDefaultConfig(context.TODO(),
		awsconfig.WithRegion(awsRegion),
		awsconfig.WithEndpointResolverWithOptions(customResolver),
	)
	if err != nil {
		logger.Error("Cannot load the AWS configs: %s", err)
		return awsCfg, err
	}
	return awsCfg, nil

}

func (c *config) ConfigAws() {
	var cfg aws.Config
	var err error

	if c.env == "test" {
		cfg, err = c.ConfigLocalAws()
		if err != nil {
			log.Fatal(err.Error())

		}
	} else {
		cfg, err = awsconfig.LoadDefaultConfig(context.TODO())
		if err != nil {
			log.Fatal(err.Error())

		}
	}

	c.awsConfig = cfg

}
