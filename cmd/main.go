package cmd

import (
	"context"

	"log"

	"github.com/aws/aws-sdk-go-v2/service/sts"
	"github.com/grafana/grafana-aws-sdk/pkg/awsauth"
)

func Execute() {
	ctx := context.Context(context.Background())

	authConfig := awsauth.NewConfigProvider()

	authSettings := awsauth.Settings{}

	cfg, err := authConfig.GetConfig(ctx, authSettings)
	if err != nil {
		log.Fatalf("Failed to get AWS auth settings: %v", err)
		return
	}

	log.Printf("AWS auth succeeded. Region: %v", cfg.Region)

	stsClient := sts.NewFromConfig(cfg)

	callerId, err := stsClient.GetCallerIdentity(ctx, &sts.GetCallerIdentityInput{})
	if err != nil {
		log.Fatalf("Failed to get caller identity: %v", err)
		return
	}
	log.Printf("CallerId: %v", *callerId.Arn)
}
