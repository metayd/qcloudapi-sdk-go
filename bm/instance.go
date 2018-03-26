package bm

type DescribeDeviceArgs struct {
	VpcId       *int      `qcloud_arg:"vpcId"`
	SubnetId    *int      `qcloud_arg:"subnetId"`
	LanIps      *[]string `qcloud_arg:"lanIps"`
	InstanceIds *[]string `qcloud_arg:"instanceIds"`
}

type BmDeviceDetail struct {
	InstanceId   string `json:"instanceId"`
	SubnetId     string `json:"subnetId"`
	VpcId        string `json:"vpcId"`
	LanIp        string `json:"lanIp"`
	DeviceStatus string `json:"deviceStatus"`
	ZoneId       string `json:"zoneId"`
	WanIp        string `json:"wanIp"`
	UnVpcId      string `json:"unVpcId"`
	UnSubnetId   string `json:"unSubnetId"`
}

type BmDeviceInfo struct {
	TotalNum   int              `json:"totalNum"`
	DeviceList []BmDeviceDetail `json:"deviceList"`
}

type BmResponse struct {
	Response interface{} `json:"data"`
}

//https://cloud.tencent.com/document/product/386/6728
func (bmClient *Client) DescribeDevice(args *DescribeDeviceArgs) (*BmDeviceInfo, error) {
	bmInfo := &BmDeviceInfo{}
	rsp := &BmResponse{
		Response: bmInfo,
	}

	err := bmClient.Invoke("DescribeDevice", args, rsp)
	if err != nil {
		return nil, err
	}

	return bmInfo, nil
}


