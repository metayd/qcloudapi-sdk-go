package cis

import (
	"os"

	"github.com/dbdd4us/qcloudapi-sdk-go/common"
)

const (
	CisHost = "cis.tencentcloudapi.com"
	CisPath = "/"
)

type Client struct {
	*common.Client
}

func NewClient(credential common.CredentialInterface, opts common.Opts) (*Client, error) {
	if opts.Host == "" {
		opts.Host = CisHost
	}
	if opts.Path == "" {
		opts.Path = CisPath
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
	region := os.Getenv("QCloudCisAPIRegion")
	host := os.Getenv("QCloudCisAPIHost")
	path := os.Getenv("QCloudCisAPIPath")

	return NewClient(
		common.Credential{
			secretId,
			secretKey,
		},
		common.Opts{
			Region: region,
			Host:   host,
			Path:   path,
		},
	)
}
