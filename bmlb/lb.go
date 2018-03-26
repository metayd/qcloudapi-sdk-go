package bmlb

import (
	"errors"
	"time"
)

const (
	LoadBalancerNameMaxLenth = 20

	TASK_STATE_SUCCESS = 0
	TASK_STATE_FAILED  = 1
	TASK_STATE_DOING   = 2
	TASK_STATE_UNKNOWN = 3
	TASK_STATE_TIMEOUT = 4

	TASK_QUERY_INTERVAL = 1
)

type LbRequestIdResponse struct {
	RequestId int `json:"requestId"`
}

type LbQueryTaskResultRequest struct {
	RequestId int `qcloud_arg:"requestId"`
}

type LbQueryTaskStatus struct {
	Status int `json:"status"`
}

type BmLbResponse struct {
	Response interface{} `json:"data"`
}

//https://cloud.tencent.com/document/product/386/9308
func (client *Client) QueryLbTaskResult(requestId int) (int, error) {
	req := &LbQueryTaskResultRequest{
		RequestId: requestId,
	}

	taskStatus := &LbQueryTaskStatus{}
	rsp := &BmLbResponse{
		Response: taskStatus,
	}

	err := client.Invoke("DescribeBmLoadBalancersTaskResult", req, rsp)
	if err != nil {
		return TASK_STATE_UNKNOWN, err
	}

	return taskStatus.Status, nil
}

type DescribeBmLbReqest struct {
	Name  *string   `qcloud_arg:"loadBalancerName"`
	LbIds *[]string `qcloud_arg:"loadBalancerIds"`
}

type BmLoadBalancerDetail struct {
	LbId       string   `json:"loadBalancerId"`
	LbName     string   `json:"loadBalancerName"`
	LbType     string   `json:"loadBalancerType"`
	LbVips     []string `json:"loadBalancerVips"`
	UnSubnetId string   `json:"unSubnetId"`
	ProjectId  int      `json:"projectId"`
	PayMode    string   `json:"payMode"`
	TgwSetType string   `json:"tgwSetType"`
	Status     int      `json:"status"` //0创建中，1表示正常运行
}

type BmLoadBalancerSetResonse struct {
	TotalCount int                    `json:"totalCount"`
	LbSet      []BmLoadBalancerDetail `json:"loadBalancerSet"`
}

//https://cloud.tencent.com/document/product/386/9306
func (client *Client) DescribeBmLoadBalancers(req *DescribeBmLbReqest) (*[]BmLoadBalancerDetail, error) {
	rsp := &BmLoadBalancerSetResonse{}

	err := client.Invoke("DescribeBmLoadBalancers", req, rsp)
	if err != nil {
		return nil, err
	}

	return &rsp.LbSet, nil
}

const (
	BM_LOAD_BALANCER_TYPE_INTERNAL = "internal"
	BM_LOAD_BALANCER_TYPE_OPEN     = "open"

	BM_PAY_MODE_FLOW      = "flow"      //流量模式
	BM_PAY_MODE_BANDWIDTH = "bindwidth" //带宽模式

	BM_LB_STATE_CREATING = 0
	BM_LB_STATE_RUNNING  = 1
)

type CreateBmLbReq struct {
	UnVpcId    string  `qcloud_arg:"unVpcId"`
	LbType     string  `qcloud_arg:"loadBalancerType"`
	UnSubnetId *string `qcloud_arg:"unSubnetId,omitempty"`
	ProjectId  *int    `qcloud_arg:"projectId,omitempty"`
	GoodsNum   *int    `qcloud_arg:"goodsNum,omitempty"`
	PayMode    *string `qcloud_arg:"payMode,omitempty"`
	TgwSetType *string `qcloud_arg:"tgwSetType,omitempty"`
}

type CreateBmLbResponse struct {
	LoadBalancerIds []string `json:"loadBalancerIds"`
}

//创建LB，https://cloud.tencent.com/document/product/386/9303
func (client *Client) CreateBmLoadBalancer(req *CreateBmLbReq) (*[]string, error) {
	rsp := &CreateBmLbResponse{}
	err := client.Invoke("CreateBmLoadBalancer", req, rsp)
	if err != nil {
		return nil, err
	}
	return &rsp.LoadBalancerIds, nil
}

type DeleteBmLbRequest struct {
	LbId string `qcloud_arg:"loadBalancerId"`
}

//删除LB：https://cloud.tencent.com/document/product/386/9304
func (client *Client) DeleteBmLoadBalancers(req *DeleteBmLbRequest) (int, error) {
	rsp := &LbRequestIdResponse{}
	err := client.Invoke("DeleteBmLoadBalancers", req, rsp)
	if err != nil {
		return 0, err
	}
	return rsp.RequestId, nil
}

type BmModifyLbAttributeRequest struct {
	LbId           string  `qcloud_arg:"loadBalancerId"`
	LbName         *string `qcloud_arg:"loadBalancerName,omitempty"`
	LbDomainPrefix *string `qcloud_arg:"domainPrefix,omitempty"`
}

//修改LB属性接口:https://cloud.tencent.com/document/product/386/9302
func (client *Client) ModifyBmLoadBalancerAttributes(req *BmModifyLbAttributeRequest) (int, error) {
	rsp := &LbRequestIdResponse{}
	err := client.Invoke("ModifyBmLoadBalancerAttributes", req, rsp)
	if err != nil {
		return 0, err
	}
	return rsp.RequestId, nil
}

func (client *Client) WaitUntiTaskDone(taskId int, timeout int) error {
	count := 0
	for {
		time.Sleep(TASK_QUERY_INTERVAL * time.Second)
		count++

		state, err := client.QueryLbTaskResult(taskId)
		if err != nil {
			return err
		}
		if state == TASK_STATE_SUCCESS {
			return nil
		} else if state == TASK_STATE_FAILED {
			return errors.New("bmLb waitUntilTaskDone task failed")
		}

		if count*TASK_QUERY_INTERVAL < timeout {
			continue
		} else {
			return errors.New("bmLb waitUntilTaskDone task timeout")
		}
	}
}
