package cvm

const (
	FilterNameZone               = "zone"
	FilterNameProjectId          = "project-id"
	FilterNameHostId             = "host-id"
	FilterNameInstanceId         = "instance-id"
	FilterNameInstanceName       = "instance-name"
	FilterNameInstanceChargeType = "instance-charge-type"
	FilterNamePrivateIpAddress   = "private-ip-address"
	FilterNamePublicIpAddress    = "public-ip-address"
)

type DescribeInstancesArgs struct {
	Version     string   `url:"Version"`
	InstanceIds []string `url:"InstanceIds"`
	Filters     []Filter `url:"Filters"`
	Offset      int      `url:"Offset"`
	Limit       int      `url:"Limit"`
}

type Filter struct {
	Name   string        `url:"Name"`
	Values []interface{} `url:"Values"`
}

type DescribeInstancesResponse struct {
	Response struct {
		TotalCount  int        `json:"TotalCount"`
		InstanceSet []Instance `json:"InstanceSet"`
	}
	RequestId string `json:"RequestId"`
}

type Instance struct {
	InstanceId string `json:"InstanceId"`
}

func (client *Client) DescribeInstances(args *DescribeInstancesArgs) (*DescribeInstancesResponse, error) {
	response := &DescribeInstancesResponse{}
	err := client.Invoke("DescribeInstances", args, response)
	if err != nil {
		return &DescribeInstancesResponse{}, err
	}
	return response, nil
}
