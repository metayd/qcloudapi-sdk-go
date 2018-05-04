package ccs

import (
	"os"

	"github.com/dbdd4us/qcloudapi-sdk-go/common"
)

const (
	CCSHost = "ccs.api.qcloud.com"
	CCSPath = "/v2/index.php"
)

type Client struct {
	*common.Client
}

func NewClient(credential common.CredentialInterface, opts common.Opts) (*Client, error) {
	if opts.Host == "" {
		opts.Host = CCSHost
	}
	if opts.Path == "" {
		opts.Path = CCSPath
	}

	client, err := common.NewClient(credential, opts)
	if err != nil {
		return &Client{}, err
	}
	return &Client{client}, nil
}

func NewClientFromEnv() (*Client, error) {

	secretId := os.Getenv("QCloudSecretId")
	secretKey := os.Getenv("QCloudSecretKey")
	region := os.Getenv("QCloudCCSAPIRegion")

	return NewClient(
		common.Credential{
			secretId,
			secretKey,
		},
		common.Opts{
			Region: region,

		},
	)
}
