package bmvpc

import (
	"errors"
	"github.com/dbdd4us/qcloudapi-sdk-go/common"
	"time"
)

const (
	TASK_STATE_SUCCESS = "success"
	TASK_STATE_FAILED  = "fail"
	TASK_STATE_DOING   = "doing"
	TASK_STATE_UNKNOWN = "unknown"
	TASK_STATE_TIMEOUT = "timeout"

	TASK_QUERY_INTERVAL = 1
)

type BmVpcResponse struct {
	Response interface{} `json:"data"`
}

type DescribeBmSubnetRequest struct {
	UnVpcId    *string `qcloud_arg:"unVpcId,omitempty"`
	UnSubnetId *string `qcloud_arg:"unSubnetId,omitempty"`
	SubnetName *string `qcloud_arg:"subnetName,omitempty"`
	VlanId     *int    `qcloud_arg:"vlanId,omitempty"`
	Limit      *int    `qcloud_arg:"limit,omitempty"`
	Offset     *int    `qcloud_arg:"offset,omitempty"`
}

type BmSubnetDetail struct {
	VpcId            int    `json:"vpcId"`
	UnVpcId          string `json:"unVpcId"`
	SubnetId         int    `json:"subnetId"`
	UnSubnetId       string `json:"unSubnetId"`
	SubnetName       string `json:"subnetName"`
	CidrBlock        string `json:"cidrBlock"`
	ZoneId           int    `json:"zoneId"`
	VlanId           int    `json:"vlanId"`
	DhcpEnable       int    `json:"dhcpEnable"`
	IpReserved       int    `json:"ipReserve"`
	DistributeedFlag int    `json:"distributedFlag"`
}

type BmDescribeSubnetResponse struct {
	TotalCount    int              `json:"totalCount"`
	SubnetDetails []BmSubnetDetail `json:"data"`
}

//查询子网列表：https://cloud.tencent.com/document/product/386/6648
func (client *Client) DescribeBmSubnetEx(req *DescribeBmSubnetRequest) (*BmDescribeSubnetResponse, error) {
	rsp := &BmDescribeSubnetResponse{}
	err := client.Invoke("DescribeBmSubnetEx", req, rsp)
	if err != nil {
		return nil, err
	}
	return rsp, nil
}

type DescribeBmCpmRequest struct {
	UnVpcId    string `qcloud_arg:"vpcId"`
	UnSubnetId string `qcloud_arg:"subnetId"`
}

type CpmInfo struct {
	InstanceId string `json:"instanceId"`
}

type DescribeBmCpmResponse struct {
	CpmSet []CpmInfo `json:"cpmSet"`
}

//根据子网subnetID和vpcId，查询子网中的主机instanceID
//https://cloud.tencent.com/document/product/386/9319
func (client *Client) DescribeBmCpmBySubnetId(req *DescribeBmCpmRequest) (*[]CpmInfo, error) {
	rsp := &DescribeBmCpmResponse{}
	err := client.Invoke("DescribeBmCpmBySubnetId", req, rsp)
	if err != nil {
		return nil, err
	}
	return &rsp.CpmSet, nil
}

type BmSubnetCreateParam struct {
	SubnetName      string `qcloud_arg:"subnetName"`
	CidrBlock       string `qcloud_arg:"cidrBlock"`
	DistributedFlag *int   `qcloud_arg:"distributedFlag"`
}

type CreateBmSubnetRequest struct {
	UnVpcId   string                `qcloud_arg:"unVpcId"`
	VLanId    *int                  `qcloud_arg:"vlanId"`
	SubnetSet []BmSubnetCreateParam `qcloud_arg:"subnetSet"`
}

type BmSubnetInfo struct {
	SubnetId   int    `json:"subnetId"`
	UnSubnetId string `json:"unSubnetId"`
	SubnetName string `json:"subnetName"`
	CidrBlock  string `json:"cidrBlock"`
}

type CreateBmSubnetResponse struct {
	SubnetSet []BmSubnetInfo `json:"subnetSet"`
}

//创建子网：https://cloud.tencent.com/document/product/386/9263
func (client *Client) CreateBmSubnet(req *CreateBmSubnetRequest) (*[]BmSubnetInfo, error) {
	rsp := &CreateBmSubnetResponse{}
	err := client.Invoke("CreateBmSubnet", req, rsp)
	if err != nil {
		return nil, err
	}
	return &rsp.SubnetSet, nil
}

//物理机加入和移除的时候，都使用下面这两个数据结构
type CreateBmInterfaceRequest struct {
	UnVpcId     string   `qcloud_arg:"unVpcId"`
	UnSubnetId  string   `qcloud_arg:"unSubnetId"`
	InstanceIds []string `qcloud_arg:"instanceIds"`
}

type BmVpcTask struct {
	TaskId      int      `json:"taskId"`
	ResourceIds []string `json:"instanceIds"`
}

//将物理机添加到子网：https://cloud.tencent.com/document/product/386/9265
func (client *Client) CreateBmInterface(req *CreateBmInterfaceRequest) (int, error) {
	bmVpcTask := &BmVpcTask{}
	rsp := &BmVpcResponse{
		Response: bmVpcTask,
	}

	err := client.Invoke("CreateBmInterface", req, rsp)
	if err != nil {
		return 0, err
	}

	return bmVpcTask.TaskId, nil
}

type DelBmInterfaceRequest CreateBmInterfaceRequest

//物理机中移除子网：https://cloud.tencent.com/document/product/386/9266
func (client *Client) DelBmInterface(req *DelBmInterfaceRequest) (int, error) {
	bmVpcTask := &BmVpcTask{}
	rsp := &BmVpcResponse{
		Response: bmVpcTask,
	}

	err := client.Invoke("DelBmInterface", req, rsp)
	if err != nil {
		return 0, err
	}

	return bmVpcTask.TaskId, nil
}

type DeleteBmSubnetRequest struct {
	UnVpcId    string `qcloud_arg:"unVpcId"`
	UnSubnetId string `qcloud_arg:"unSubnetId"`
}

//https://cloud.tencent.com/document/product/386/9264
func (client *Client) DeleteBmSubnet(req *DeleteBmSubnetRequest) error {
	rsp := &common.LegacyAPIError{}
	err := client.Invoke("DeleteBmSubnet", req, rsp)
	if err != nil {
		return err
	}
	if rsp.Code == 0 {
		return nil
	} else {
		return errors.New(rsp.Message)
	}
}

type BmVpcQueryTaskRequest struct {
	TaskId int `qcloud_arg:"taskId"`
}

type BmVpcTaskStatus struct {
	Status map[string]string `json:"data"`
}

//https://cloud.tencent.com/document/product/386/9267
func (client *Client) QueryBmTaskResult(taskId int) (string, error) {
	req := BmVpcQueryTaskRequest{
		TaskId: taskId,
	}

	rsp := &BmVpcTaskStatus{}

	err := client.Invoke("QueryBmTaskResult", req, rsp)
	if err != nil || len(rsp.Status) != 1 {
		return TASK_STATE_UNKNOWN, err
	}

	for _, val := range rsp.Status {
		return val, nil
	}
	return TASK_STATE_UNKNOWN, errors.New("QueryBmTaskResult can't go here")
}

func (client *Client) WaitUntiTaskDone(taskId int, timeout int) error {
	count := 0
	for {
		time.Sleep(TASK_QUERY_INTERVAL * time.Second)
		count++

		state, err := client.QueryBmTaskResult(taskId)
		if err != nil {
			return err
		}
		if state == TASK_STATE_SUCCESS {
			return nil
		} else if state == TASK_STATE_FAILED {
			return errors.New("bmVpc waitUntilTaskDone task failed")
		} else if state == TASK_STATE_UNKNOWN {
			return errors.New("bmVpc waitUntilTaskDone task state unknown")
		}

		if count*TASK_QUERY_INTERVAL < timeout {
			continue
		} else {
			return errors.New("bmVpc waitUntilTaskDone task timeout")
		}
	}
}

type BmCreateContainerSubnetReq struct {
	UnVpcId   string                `qcloud_arg:"unVpcId"`
	VLanId    *int                  `qcloud_arg:"vlanId"`
	SubnetType int                  `qcloud_arg:"subnetType"`      // 6:ccs; 7:docker, 默认为docker子网
	SubnetSet []BmSubnetCreateParam `qcloud_arg:"subnetSet"`       //只要填subnetName和cidrBlock"即可

}


func (client *Client) CreateBmContainerSubnet(req *BmCreateContainerSubnetReq)(*[]BmSubnetInfo, error) {
	rsp := &CreateBmSubnetResponse{}
	err := client.Invoke("CreateBmContainerSubnet", req, rsp)
	if err != nil {
		return nil, err
	}
	return &rsp.SubnetSet, nil
}


func (client *Client) DeleteBmContainerSubnet (req *DeleteBmSubnetRequest)( error){
	rsp := &common.LegacyAPIError{}
	err := client.Invoke("DeleteBmContainerSubnet", req, rsp)
	if err != nil {
		return err
	}
	if rsp.Code == 0 {
		return nil
	} else {
		return errors.New(rsp.Message)
	}

}


type RegisterBatchIpRequest struct {
	UnVpcId    string   `qcloud_arg:"unVpcId"`
	UnSubnetId string   `qcloud_arg:"unSubnetId"`
	IPList     []string `qcloud_arg:"ipList"`
}

//将物理机添加到子网：https://cloud.tencent.com/document/product/386/9265
func (client *Client) RegisterBatchIp(req *RegisterBatchIpRequest) error {
	rsp := &common.LegacyAPIError{}

	err := client.Invoke("RegisterBatchIp", req, rsp)
	if err != nil {
		return err
	}
	if rsp.Code == 0 {
		return nil
	} else {
		return errors.New(rsp.Message)
	}
}

type ReturnIpsRequest struct {
	UnVpcId string   `qcloud_arg:"unVpcId"`
	IPList  []string `qcloud_arg:"ips"`
}

//将物理机添加到子网：https://cloud.tencent.com/document/product/386/9265
func (client *Client) ReturnIps(req *ReturnIpsRequest) error {
	rsp := &common.LegacyAPIError{}

	err := client.Invoke("ReturnIps", req, rsp)
	if err != nil {
		return err
	}
	if rsp.Code == 0 {
		return nil
	} else {
		return errors.New(rsp.Message)
	}
}

type DescribeBmSubnetIpsRequest struct {
	UnVpcId    string `qcloud_arg:"unVpcId"`
	UnSubnetId string `qcloud_arg:"unSubnetId"`
}

type BmSubnetIp struct {
	CpmSet []string `json:"cpmSet"`
	TgSet  []string `json:"tgSet"`
	VmSet  []string `json:"vmSet"`
}

type DescribeBmSubnetIpsResponse struct {
	SubnetIp BmSubnetIp `json:"data"`
}

func (client *Client) DescribeBmSubnetIps(req *DescribeBmSubnetIpsRequest) ([]string, error) {
	rsp := &DescribeBmSubnetIpsResponse{}

	err := client.Invoke("DescribeBmSubnetIps", req, rsp)
	if err != nil {
		return []string{}, err
	}
	return rsp.SubnetIp.VmSet, nil
}