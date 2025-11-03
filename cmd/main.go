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
		log.Fatalf("Failed to get AWS config: %v", err)
		return
	}

	creds, err := cfg.Credentials.Retrieve(ctx)
	log.Printf("AWS config worked. Region: %v. AccessKey: %v. CredsSource: %v", cfg.Region, creds.AccessKeyID, creds.Source)

	stsClient := sts.NewFromConfig(cfg)

	callerId, err := stsClient.GetCallerIdentity(ctx, &sts.GetCallerIdentityInput{})
	if err != nil {
		log.Fatalf("Failed to get caller identity: %v", err)
		return
	}
	log.Printf("CallerId: %v", *callerId.Arn)
}
