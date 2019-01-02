package bmvpc

import (
	"github.com/dbdd4us/qcloudapi-sdk-go/common"
	"os"
)

const (
	BmVpcHost = "bmvpc.api.qcloud.com"
	BmVpcPath = "/v2/index.php"
)

type Client struct {
	*common.Client
}

func NewClient(credential common.CredentialInterface, opts common.Opts) (*Client, error) {
	if opts.Host == "" {
		opts.Host = BmVpcHost
	}
	if opts.Path == "" {
		opts.Path = BmVpcPath
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
	region := os.Getenv("QCloudBmAPIRegion")
	host := os.Getenv("QCloudBmVpcLbAPIHost")
	path := os.Getenv("QCloudBmVpcAPIPath")

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
