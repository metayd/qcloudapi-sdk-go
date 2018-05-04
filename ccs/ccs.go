package ccs

const (
	//CvmType
	PayByHour = "PayByHour"
	PayByMonth = "PayByMonth"


	//storageType
	LOCAL_BASIC = "LOCAL_BASIC"
	LOCAL_SSD = "LOCAL_SSD"
	CLOUD_BASIC = "CLOUD_BASIC"
	CLOUD_PREMIUM = "CLOUD_PREMIUM"
	CLOUD_SSD = "CLOUD_SSD"

	//bandwidthType
	PayByTraffic = "PayByTraffic"
	PayByBandwidth = "PayByBandwidth"

	//ZoneId
	ApShanghai_1 = "ap-shanghai-1"
	ApShanghai_2 = "ap-shanghai-2"
	ApShanghai_3 = "ap-shanghai-3"
	ApShanghai_4 = "ap-shanghai-4"
)

type DescribeClusterArgs struct {
	ClusterIds  []*string `qcloud_arg:"clusterIds"`
	ClusterName *string   `qcloud_arg:"clusterName"`
	Status      *string   `qcloud_arg:"status"`
	OrderField  *string   `qcloud_arg:"orderField"`
	OrderType   *string   `qcloud_arg:"orderType"`
	Offset      *int      `qcloud_arg:"offset"`
	Limit       *int      `qcloud_arg:"limit"`
}

type Cluster struct {
	ClusterCIDR             *string `json:"clusterCIDR"`
	ClusterExternalEndpoint *string `json:"clusterExternalEndpoint"`
	ClusterId               *string `json:"clusterId"`
	ClusterName             *string `json:"clusterName"`
	CreatedAt               *string `json:"createdAt"`
	Description             *string `json:"description"`
	K8sVersion              *string `json:"k8sVersion"`
	MasterLbSubnetId        *string `json:"masterLbSubnetId"`
	NodeNum                 *int    `json:"nodeNum"`
	NodeStatus              *string `json:"nodeStatus"`
	OpenHttps               *int    `json:"openHttps"`
	OS                      *string `json:"os"`
	ProjectId               *int    `json:"projectId"`
	Region                  *string `json:"region"`
	RegionId                *int    `json:"regionId"`
	Status                  *string `json:"status"`
	TotalCPU                *int    `json:"totalCpu"`
	TotalMem                *int    `json:"totalMem"`
	UnVpcId                 *string `json:"unVpcId"`
	UpdatedAt               *string `json:"updatedAt"`
	VpcId                   *int    `json:"vpcId"`
}

type DescribeClusterResponse struct {
	Response
	Data struct {
			 TotalCount int       `json:"totalCount"`
			 Clusters   []Cluster `json:"clusters"`
		 } `json:"data"`
}

type DescribeClusterInstancesRequest struct {
	ClusterId *string `qcloud_arg:"clusterId"`
	Offset    *int    `qcloud_arg:"offset"`
	Limit     *int    `qcloud_arg:"limit"`
}

type ClusterInstance struct {
	AbnormalReason       *string            `json:"abnormalReason"`
	AutoScalingGroupId   *string            `json:"autoScalingGroupId"`
	CPU                  *int               `json:"cpu"`
	CreatedAt            *string            `json:"createdAt"`
	CvmPayMode           *int               `json:"cvmPayMode"`
	CvmState             *int               `json:"cvmState"`
	InstanceCreateTime   *string            `json:"instanceCreateTime"`
	InstanceDeadlineTime *string            `json:"instanceDeadlineTime"`
	InstanceId           *string            `json:"instanceId"`
	InstanceName         *string            `json:"instanceName"`
	InstanceType         *string            `json:"instanceType"`
	IsNormal             *int               `json:"isNormal"`
	KernelVersion        *string            `json:"kernelVersion"`
	//Labels               *map[string]string `json:"labels,omitempty"`
	LanIp                *string            `json:"lanIp"`
	Mem                  *int               `json:"mem"`
	NetworkPayMode       *int               `json:"networkPayMode"`
	OSImage              *string            `json:"osImage"`
	PodCidr              *string            `json:"podCidr"`
	Unschedulable        *bool              `json:"unschedulable"`
	WanIp                *string            `json:"wanIp"`
	Zone                 *string            `json:"zone"`
	ZoneId               *int               `json:"zoneId"`
}

type DescribeClusterInstancesResponse struct {
	Response
	Data *struct {
		TotalCount *int               `json:"totalCount"`
		Nodes      []*ClusterInstance `json:"nodes"`
	} `json:"data"`
}

type AddClusterInstancesArgs struct {
	ClusterId       string `qcloud_arg:"clusterId"`
	ZoneId          string `qcloud_arg:"zoneId"`
	Cpu             int    `qcloud_arg:"cpu"`
	Mem             int    `qcloud_arg:"mem"`
	InstanceType    string `qcloud_arg:"instanceType"`
	CvmType         string `qcloud_arg:"cvmType"`
	BandwidthType   string `qcloud_arg:"bandwidthType"`
	Bandwidth       int    `qcloud_arg:"bandwidth"`
	IsVpcGateway    int    `qcloud_arg:"isVpcGateway"`
	RootType        string `qcloud_arg:"rootType"`
	StorageType     string `qcloud_arg:"storageType"`

	WanIp           int    `qcloud_arg:"wanIp"`
	SubnetId        string `qcloud_arg:"subnetId"`
	StorageSize     int    `qcloud_arg:"storageSize"`
	RootSize        int    `qcloud_arg:"rootSize"`
	GoodsNum        int    `qcloud_arg:"goodsNum"`
	Password        *string `qcloud_arg:"password"`
	KeyId           *string `qcloud_arg:"keyId"`
	Period          int    `qcloud_arg:"period"`
	SgId            string `qcloud_arg:"sgId"`
	MountTarget     string `qcloud_arg:"mountTarget"`
	DockerGraphPath string `qcloud_arg:"dockerGraphPath"`
	UserScript       string `qcloud_arg:"userScript"`
}

type AddClusterInstancesResponse struct {
	Response
	Data *struct {
		InstanceIds []*string `json:"instanceIds"`
		RequestId   *int      `json:"requestId"`
	}
}

type DescribeClusterTaskResultRequest struct {
	RequestId *int `qcloud_arg:"requestId"`
}

type Response struct {
	Code     *int    `json:"code"`
	Message  string `json:"message"`
	CodeDesc string `json:"codeDesc"`
}

type DescribeClusterTaskResultResponse struct {
	Response
	Data *struct {
		Status *string `json:"status"`
	} `json:"data"`
}

func (client *Client) DescribeCluster(args *DescribeClusterArgs) (*DescribeClusterResponse, error) {
	response := &DescribeClusterResponse{}
	err := client.Invoke("DescribeCluster", args, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (client *Client) AddClusterInstances(args *AddClusterInstancesArgs) (*AddClusterInstancesResponse, error) {
	response := &AddClusterInstancesResponse{}
	err := client.Invoke("AddClusterInstances", args, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}