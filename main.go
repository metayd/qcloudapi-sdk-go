package main

import (
	"fmt"
	"log"
	"os"

	"github.com/dbdd4us/qcloudapi-sdk-go/client"
	"github.com/dbdd4us/qcloudapi-sdk-go/types"
)

func main() {
	secretId := os.Getenv("QCloudAPISecretId")
	secretKey := os.Getenv("QcloudAPISecretKey")
	credential := client.NewCredential(
		[]byte(secretId),
		[]byte(secretKey),
	)
	cli, err := client.NewClient(credential, client.Opts{"gz"})
	if err != nil {
		log.Fatal(err)
	}
	args := types.DescribeLoadBalancersArgs{
		//LoadBalancerIds: []string{"id-1", "id-2"},
		LoadBalancerType: 1,
	}

	resp, err := cli.DescribeLoadBalancers(args)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp)
}
