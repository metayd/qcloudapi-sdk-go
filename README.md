# qcloudapi-sdk-go

[![Go Report Card](https://goreportcard.com/badge/github.com/dbdd4us/qcloudapi-sdk-go)](https://goreportcard.com/report/github.com/dbdd4us/qcloudapi-sdk-go)
[![Build Status](https://travis-ci.org/dbdd4us/qcloudapi-sdk-go.svg?branch=master)](https://travis-ci.org/dbdd4us/qcloudapi-sdk-go)
[![codecov](https://codecov.io/gh/dbdd4us/qcloudapi-sdk-go/branch/master/graph/badge.svg)](https://codecov.io/gh/dbdd4us/qcloudapi-sdk-go)
[![GoDoc](https://godoc.org/github.com/dbdd4us/qcloudapi-sdk-go?status.svg)](http://godoc.org/github.com/dbdd4us/qcloudapi-sdk-go)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)

This is an unofficial Go SDK for QCloud Services. You are welcome for contribution.


## Usage

```go
package main

import (
	"log"

	"github.com/dbdd4us/qcloudapi-sdk-go/clb"
	"github.com/dbdd4us/qcloudapi-sdk-go/common"
)

func main() {
	credential := common.Credential{
		SecretId: "YOUR_SECRET_ID",
		SecretKey: "YOUR_SECRET_KEY",
	}
	opts := common.Opts{
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
