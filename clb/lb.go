package clb

import (
	"net/url"

	"github.com/dbdd4us/qcloudapi-sdk-go/common"
)

type DescribeLoadBalancersArgs struct {
	LoadBalancerIds  LoadBalancerIds  `url:"loadBalancerIds,omitempty"`
	LoadBalancerType int              `url:"loadBalancerType,omitempty"`
	LoadBalancerName string           `url:"loadBalancerName,omitempty"`
	Domain           string           `url:"domain,omitempty"`
	LoadBalancerVips LoadBalancerVips `url:"loadBalancerVips,omitempty"`
	BackendWanIps    BackendWanIps    `url:"backendWanIps,omitempty"`
	Offset           int              `url:"offset,omitempty"`
	Limit            int              `url:"limit,omitempty"`
	OrderBy          string           `url:"orderBy,omitempty"`
	OrderType        int              `url:"orderType,omitempty"`
	SearchKey        string           `url:"searchKey,omitempty"`
	ProjectId        int              `url:"projectId,omitempty"`
	Forward          int              `url:"forward,omitempty"`
	WithRs           int              `url:"withRs,omitempty"`
}

type LoadBalancerIds []string

func (field LoadBalancerIds) EncodeValues(key string, v *url.Values) error {
	return common.EncodeArgs(key, field, v)
}

type LoadBalancerVips []string

func (field LoadBalancerVips) EncodeValues(key string, v *url.Values) error {
	return common.EncodeArgs(key, field, v)
}

type BackendWanIps []string

func (field BackendWanIps) EncodeValues(key string, v *url.Values) error {
	return common.EncodeArgs(key, field, v)
}

type DescribeLoadBalancersResponse struct {
	Code            int            `json:"code"`
	Message         string         `json:"message"`
	TotalCount      int            `json:"totalCount"`
	LoadBalancerSet []LoadBalancer `json:"loadBalancerSet"`
	CodeDesc        string         `json:"codeDesc"`
}

type LoadBalancer struct {
	LoadBalancerId   string           `json:"loadBalancerId"`
	UnLoadBalancerId string           `json:"unLoadBalancerId"`
	LoadBalancerName string           `json:"loadBalancerName"`
	LoadBalancerType int              `json:"loadBalancerType"`
	Domain           string           `json:"domain"`
	LoadBalancerVips LoadBalancerVips `json:"loadBalancerVips"`
	Status           int              `json:"status"`
	CreateTime       string           `json:"createTime"`
	StatusTime       string           `json:"statusTime"`
	ProjectId        int              `json:"projectId"`
	VpcId            int              `json:"vpcId"`
	SubnetId         int              `json:"subnetId"`
}

func (client *Client) DescribeLoadBalancers(args *DescribeLoadBalancersArgs) (*DescribeLoadBalancersResponse, error) {
	response := &DescribeLoadBalancersResponse{}
	err := client.Invoke("DescribeLoadBalancers", args, response)
	if err != nil {
		return &DescribeLoadBalancersResponse{}, err
	}
	return response, nil
}
