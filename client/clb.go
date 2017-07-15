package client

import (
	"github.com/dbdd4us/qcloudapi-sdk-go/types"
)

const (
	ClbHost = "lb.api.qcloud.com"
)

func (cli *Client) DescribeLoadBalancers(args types.DescribeLoadBalancersArgs) (*types.DescribeLoadBalancersResponse, error) {
	resp := &types.DescribeLoadBalancersResponse{}
	err := cli.Invoke(ClbHost, "DescribeLoadBalancers", args, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}