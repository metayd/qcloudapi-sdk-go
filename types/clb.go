package types

import (
	"net/url"
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
	return EncodeArgs(key, field, v)
}

type LoadBalancerVips []string

func (field LoadBalancerVips) EncodeValues(key string, v *url.Values) error {
	return EncodeArgs(key, field, v)
}

type BackendWanIps []string

func (field BackendWanIps) EncodeValues(key string, v *url.Values) error {
	return EncodeArgs(key, field, v)
}

type DescribeLoadBalancersResponse struct{}
