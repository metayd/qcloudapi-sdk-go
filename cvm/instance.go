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
	Version     *string   `qcloud_arg:"Version,required"`
	InstanceIds *[]string `qcloud_arg:"InstanceIds"`
	Filters     *[]Filter `qcloud_arg:"Filters"`
	Offset      *int      `qcloud_arg:"Offset"`
	Limit       *int      `qcloud_arg:"Limit"`
}

type Filter struct {
	Name   string        `qcloud_arg:"Name"`
	Values []interface{} `qcloud_arg:"Values"`
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
