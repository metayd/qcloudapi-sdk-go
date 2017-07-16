# qcloudapi-sdk-go

[![Go Report Card](https://goreportcard.com/badge/github.com/dbdd4us/qcloudapi-sdk-go)](https://goreportcard.com/report/github.com/dbdd4us/qcloudapi-sdk-go)

This is an unofficial Go SDK for QCloud Services. You are welcome for contribution.


## Usage

```go
package main

import (
	"log"

	"github.com/dbdd4us/qcloudapi-sdk-go/clb"
)

func main() {
	credential := clb.Credential{
		SecretId: "YOUR_SECRET_ID",
		SecretKey: "YOUR_SECRET_KEY",
	}
	opts := clb.Opts{
		Region: "gz",
	}
	client, err := clb.NewClient(credential, opts)
	if err != nil {
		log.Fatal(err)
	}
	args := clb.DescribeLoadBalancersArgs{}

	lbs, err := client.DescribeLoadBalancers(&args)
	if err != nil {
		log.Fatal(lbs)
	}
	log.Println(lbs.LoadBalancerSet)
}



```


## License

This library is distributed under the Apache License found in the [LICENSE](./LICENSE) file.
