package clb

type RegisterInstancesWithLoadBalancerArgs struct {
	LoadBalancerId string        `qcloud_arg:"loadBalancerId,required"`
	Backends       []backendOpts `qcloud_arg:"backends,required"`
}

type backendOpts struct {
	InstanceId string `qcloud_arg:"instanceId,required"`
	Weight     *int   `qcloud_arg:"weight"`
}

type RegisterInstancesWithLoadBalancerResponse struct {
	Response
	RequestId int `json:"requestId"`
}

func (client *Client) RegisterInstancesWithLoadBalancer(args *RegisterInstancesWithLoadBalancerArgs) (
	*RegisterInstancesWithLoadBalancerResponse,
	error,
) {
	response := &RegisterInstancesWithLoadBalancerResponse{}
	err := client.Invoke("RegisterInstancesWithLoadBalancer", args, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

type DescribeLoadBalancerBackendsArgs struct {
	LoadBalancerId string `qcloud_arg:"loadBalancerId,required"`
	Offset         int    `qcloud_arg:"offset"`
	Limit          int    `qcloud_arg:"limit"`
}

type DescribeLoadBalancerBackendsResponse struct {
	Response
	TotalCount int       `json:"totalCount"`
	BackendSet []backend `json:"backendSet"`
}

type backend struct {
	InstanceId     string   `json:"instanceId"`
	UnInstanceId   string   `json:"unInstanceId"`
	Weight         int      `json:"weight"`
	InstanceName   string   `json:"instanceName"`
	LanIp          string   `json:"lanIp"`
	WanIpSet       []string `json:"wanIpSet"`
	InstanceStatus int      `json:"instanceStatus"`
}

func (client *Client) DescribeLoadBalancerBackends(args *DescribeLoadBalancerBackendsArgs) (
	*DescribeLoadBalancerBackendsResponse,
	error,
) {
	response := &DescribeLoadBalancerBackendsResponse{}
	err := client.Invoke("DescribeLoadBalancerBackends", args, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

type ModifyLoadBalancerBackendsArgs struct {
	LoadBalancerId string              `qcloud_arg:"loadBalancerId,required"`
	Backends       []modifyBackendOpts `qcloud_arg:"backends,required"`
}

type modifyBackendOpts struct {
	InstanceId string `qcloud_arg:"instanceId,required"`
	Weight     int    `qcloud_arg:"weight,required"`
}

type ModifyLoadBalancerBackendsResponse struct {
	Response
	RequestId int `json:"requestId"`
}

func (client *Client) ModifyLoadBalancerBackends(args *ModifyLoadBalancerBackendsArgs) (
	*ModifyLoadBalancerBackendsResponse,
	error,
) {
	response := &ModifyLoadBalancerBackendsResponse{}
	err := client.Invoke("ModifyLoadBalancerBackends", args, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

type DeregisterInstancesFromLoadBalancerArgs struct {
	LoadBalancerId string              `qcloud_arg:"loadBalancerId,required"`
	Backends       []deRegisterBackend `qcloud_arg:"backends,required"`
}

type deRegisterBackend struct {
	InstanceId string `qcloud_arg:"instanceId"`
}

type DeregisterInstancesFromLoadBalancerResponse struct {
	Response
	RequestId int `json:"requestId"`
}

func (client *Client) DeregisterInstancesFromLoadBalancer(args *DeregisterInstancesFromLoadBalancerArgs) (
	*DeregisterInstancesFromLoadBalancerResponse,
	error,
) {
	response := &DeregisterInstancesFromLoadBalancerResponse{}
	err := client.Invoke("DeregisterInstancesFromLoadBalancer", args, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}
