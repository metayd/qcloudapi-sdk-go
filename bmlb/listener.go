package bmlb

type BmlbResponse struct {
	Response interface{} `json:"data"`
}

const (
	LISTENER_STATE_CREATING      = 0
	LISTENER_STATE_RUNNING       = 1
	LISTENER_STATE_CREATE_FAILED = 2
	LISTENER_STATE_DELETING      = 3
	LISTENER_STATE_DELETE_FAILED = 4
)

type DescribeBmLbListenerRequest struct {
	LbId         string    `qcloud_arg:"loadBalancerId"`
	ListenersIds *[]string `qcloud_arg:"listenerIds,omitempty"`
}

type LbListenerDetail struct {
	Id                 string `json:"listenerId"`
	Name               string `json:"listenerName"`
	Protocol           string `json:"protocol"`
	LbPort             int    `json:"loadBalancerPort"`
	Bandwidth          int    `json:"bandwidth"`
	ListenerType       string `json:"listenerType"`
	SessionExpire      int    `json:"sessionExpire"`
	HealthSwitch       int    `json:"healthSwitch"`
	TimeOut            int    `json:"timeOut"`
	IntervalTime       int    `json:"intervalTime"`
	HealthNum          int    `json:"healthNum"`
	UnhealthNum        int    `json:"unhealthNum"`
	CustomHealthSwitch int    `json:"customHealthSwitch"`
	InputType          string `json:"inputType"`
	LineSeparatorType  int    `json:"lineSeparatorType"`
	HealthRequest      string `json:"healthRequest"`
	HealthResponse     string `json:"healthResponse"`
	ToaFlag            int    `json:"toaFlag"`
	Status             int    `json:"status"`
	AddTimestamp       string `json:"AddTimestamp"`
}

type LbListenerSetResponse struct {
	TotalCount  int                `json:"totalCount"`
	ListenerSet []LbListenerDetail `json:"listenerSet"`
}

// https://cloud.tencent.com/document/product/386/9296
func (client *Client) DescribeBmListeners(req *DescribeBmLbListenerRequest) (*LbListenerSetResponse, error) {
	rsp := &LbListenerSetResponse{}

	err := client.Invoke("DescribeBmListeners", req, rsp)
	if err != nil {
		return nil, err
	}
	return rsp, nil
}

type BackendRs struct {
	Port       int    `qcloud_arg:"port"`
	ProbePort  *int   `qcloud_arg:"probePort,omitempty"`
	InstanceId string `qcloud_arg:"instanceId"`
	Weight     int    `qcloud_arg:"weight"`
}

type BmBindL4LisenerRsRequest struct {
	LbId       string      `qcloud_arg:"loadBalancerId"`
	ListenerId string      `qcloud_arg:"listenerId"`
	Backends   []BackendRs `qcloud_arg:"backends"`
}

//https://cloud.tencent.com/document/product/386/9294
func (client *Client) BindBmL4ListenerRs(req *BmBindL4LisenerRsRequest) (int, error) {
	rsp := &LbRequestIdResponse{}
	err := client.Invoke("BindBmL4ListenerRs", req, rsp)
	if err != nil {
		return 0, err
	}
	return rsp.RequestId, nil
}

type DescribeBm4BackendRsRequest struct {
	LbId       string `qcloud_arg:"loadBalancerId"`
	ListenerId string `qcloud_arg:"listenerId"`
}

type BmLbBackEndRsDetail struct {
	BindType int    `json:"bindType"`
	RsPort   int    `json:"rsPort"`
	Weight   int    `json:"weight"`
	Status   string `json:"status"`

	InstanceId string `json:"instanceId"`
	LanIp      string `json:"lanIp"` //当bindtype为1时表示黑石物理机的内网IP，当bindType为1时表示虚机IP
}

//获取后端rs列表，https://cloud.tencent.com/document/product/386/9297
func (client *Client) DescribeBmL4ListenerBackend(req *DescribeBm4BackendRsRequest) (*[]BmLbBackEndRsDetail, error) {
	bmLbBackendRsDetail := make([]BmLbBackEndRsDetail, 0)
	rsp := &BmlbResponse{
		Response: &bmLbBackendRsDetail,
	}

	err := client.Invoke("DescribeBmL4ListenerBackends", req, rsp)
	if err != nil {
		return nil, err
	}
	return &bmLbBackendRsDetail, nil
}

type BmUnbindRs struct {
	Port       int    `qcloud_arg:"port"`
	InstanceId string `qcloud_arg:"instanceId"`
}

type UnbindBmListenerRsRequest struct {
	LbId       string       `qcloud_arg:"loadBalancerId"`
	ListenerId string       `qcloud_arg:"listenerId"`
	Backends   []BmUnbindRs `qcloud_arg:"backends"`
}

//https://cloud.tencent.com/document/product/386/9299
func (client *Client) UnbindBmL4ListenerRs(req *UnbindBmListenerRsRequest) (int, error) {
	rsp := &LbRequestIdResponse{}
	err := client.Invoke("UnbindBmL4ListenerRs", req, rsp)
	if err != nil {
		return 0, err
	}
	return rsp.RequestId, nil
}

type BmListenerCreateDetail struct {
	LbPort             int     `qcloud_arg:"loadBalancerPort"`
	Protocol           string  `qcloud_arg:"protocol"`
	Name               string  `qcloud_arg:"listenerName,omitempty"`
	SessionExpire      *int    `qcloud_arg:"sessionExpire,omitempty"`
	HealthSwitch       *int    `qcloud_arg:"healthSwitch,omitempty"`
	Timeout            *int    `qcloud_arg:"timeOut,omitempty"`
	IntervalTime       *int    `qcloud_arg:"intervalTime,omitempty"`
	HealthNum          *int    `qcloud_arg:"healthNum,omitempty"`
	UnhealthNum        *int    `qcloud_arg:"unhealthNum,omitempty"`
	BandWidth          *int    `qcloud_arg:"bandwidth,omitempty"`
	CustomHealthSwitch *int    `qcloud_arg:"customHealthSwitch,omitempty"`
	InputType          *string `qcloud_arg:"inputType,omitempty"`
	LineSeparatorType  *int    `qcloud_arg:"lineSeparatorType,omitempty"`
	HealthRequest      *string `qcloud_arg:"healthRequest,omitempty"`
	HealthResponse     *string `qcloud_arg:"healthResponse,omitempty"`
	ToaFlag            *int    `qcloud_arg:"toaFlag,omitempty"`
}

type CreateBmListenerRequest struct {
	LbId      string                    `qcloud_arg:"loadBalancerId"`
	Listeners *[]BmListenerCreateDetail `qcloud_arg:"listeners"`
}

//创建4层监听器，https://cloud.tencent.com/document/product/386/9292
func (client *Client) CreateBmListeners(req *CreateBmListenerRequest) (int, error) {
	rsp := &LbRequestIdResponse{}
	err := client.Invoke("CreateBmListeners", req, rsp)
	if err != nil {
		return 0, err
	}
	return rsp.RequestId, nil
}

type DeleteBmListenersRequest struct {
	LbId        string    `qcloud_arg:"loadBalancerId"`
	ListenerIds *[]string `qcloud_arg:"listenerIds"`
}

//删除四层监听器，https://cloud.tencent.com/document/product/386/9293
func (client *Client) DeleteBmListeners(req *DeleteBmListenersRequest) (int, error) {
	rsp := &LbRequestIdResponse{}
	err := client.Invoke("DeleteBmListeners", req, rsp)
	if err != nil {
		return 0, err
	}
	return rsp.RequestId, nil
}
