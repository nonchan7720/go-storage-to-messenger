package aws

import (
	"context"

	awsConfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sns"
	"github.com/nonchan7720/go-storage-to-messenger/pkg/config"
	"github.com/nonchan7720/go-storage-to-messenger/pkg/interfaces/aws"
)

func NewSNSClient(ctx context.Context, conf *config.AWS) (aws.SNSClient, error) {
	endpoint := NewEndpoint(WithSNSEndpoint(conf.SNS.Endpoint))
	awsConfig, err := NewConfig(ctx,
		endpoint.EndpointResolver(),
		awsConfig.WithCredentialsProvider(conf.WithStatic()),
	)
	if err != nil {
		return nil, err
	}
	return sns.NewFromConfig(awsConfig), nil
}
