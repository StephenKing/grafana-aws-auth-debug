package cmd

import (
	"context"

	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	"github.com/grafana/grafana-aws-sdk/pkg/awsauth"
)

func Execute() {
	ctx := context.Context(context.Background())

	log.Printf("Trying regular AWS SDK...")

	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("eu-west-1"))
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	if err = tryStsGetCallerIdentity(ctx, cfg); err != nil {
		log.Fatalf("Failed to get sts caller identity: %v", err)
	}

	log.Printf("Trying Grafana...")

	authConfig := awsauth.NewConfigProvider()

	authSettings := awsauth.Settings{
		Region:             "eu-west-1",
		CredentialsProfile: "default",
	}

	cfg, err = authConfig.GetConfig(ctx, authSettings)
	if err != nil {
		log.Fatalf("Failed to get AWS config: %v", err)
	}

	creds, err := cfg.Credentials.Retrieve(ctx)
	log.Printf("AWS config worked. Region: %v. AccessKey: %v. CredsSource: %v", cfg.Region, creds.AccessKeyID, creds.Source)

	if err = tryStsGetCallerIdentity(ctx, cfg); err != nil {
		log.Fatalf("Failed to get sts caller identity: %v", err)
	}

}

func tryStsGetCallerIdentity(ctx context.Context, cfg aws.Config) error {
	stsClient := sts.NewFromConfig(cfg)
	callerId, err := stsClient.GetCallerIdentity(ctx, &sts.GetCallerIdentityInput{})
	if err != nil {
		log.Fatalf("Failed to get caller identity: %v", err)
		return err
	}
	log.Printf("CallerId: %v", *callerId.Arn)
	return nil
}
