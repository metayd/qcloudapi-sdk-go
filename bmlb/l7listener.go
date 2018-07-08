package bmlb

const (
	SSL_MODE_NOSSL          = 0
	SSL_MODE_UNIDIRECTIONAL = 1
	SSL_MODE_BIDIRECTIONAL  = 2

	L7_LISTENER_TYPE = "L7Listener"

	LISTENER_PROTO_HTTP  = "http"
	LISTENER_PROTO_HTTPS = "https"

	BALANCE_MODE_WRR     = "wrr"
	BALANCE_MODE_IP_HASH = "ip_hash"
)

type NoData struct{}

type DescribeBmLbL7ListenerRequest struct {
	LbId         string   `qcloud_arg:"loadBalancerId"`
	ListenersIds []string `qcloud_arg:"listenerIds,omitempty"`
}

type L7Listener struct {
	ListenerId   string `json:"listenerId"`
	ListenerName string `json:"listenerName"`
	Protocol     string `json:"protocol"`
	LbPort       int    `json:"loadBalancerPort"`
	Bandwidth    int    `json:"bandwidth"`
	ListenerType string `json:"listenerType"`
	SSLMode      int    `json:"SSLMode"`
	CertId       string `json:"certId"`
	CertCaId     string `json:"certCaId"`
	Status       int    `json:"status"`
	AddTimestamp string `json:"addTimestamp"`

	//for interface DescribeBmL7Listeners, this field is nil !!!
	RuleSet []*ForwardRule `json:"ruleSet"`
}

type DescribeBmL7ListenersResponse struct {
	TotalCount  int           `json:"totalCount"`
	ListenerSet []*L7Listener `json:"listenerSet"`
}

// https://cloud.tencent.com/document/api/386/9283
func (client *Client) DescribeBmL7Listeners(req *DescribeBmLbL7ListenerRequest) (*DescribeBmL7ListenersResponse, error) {
	rsp := &DescribeBmL7ListenersResponse{}

	err := client.Invoke("DescribeBmForwardListeners", req, rsp)
	if err != nil {
		return nil, err
	}
	return rsp, nil
}

type L7ListenerTA struct {
	LbPort        int     `qcloud_arg:"loadBalancerPort"`
	Protocol      string  `qcloud_arg:"protocol"`
	ListenerName  string  `qcloud_arg:"listenerName"`
	SSLMode       int     `qcloud_arg:"SSLMode"`
	CertId        *string `qcloud_arg:"certId"`
	CertName      *string `qcloud_arg:"certName"`
	CertContent   *string `qcloud_arg:"certContent"`
	CertKey       *string `qcloud_arg:"certKey"`
	CertCaId      *string `qcloud_arg:"certCaId"`
	CertCaName    *string `qcloud_arg:"certCaName"`
	CertCaContent *string `qcloud_arg:"certCaContent"`
	BandWidth     *int    `qcloud_arg:"bandwidth"`
}

type CreateBmL7ListenerRequest struct {
	LbId      string          `qcloud_arg:"loadBalancerId"`
	Listeners []*L7ListenerTA `qcloud_arg:"listeners"`
}

type CreateBmL7ListenerResponse struct {
	ListenerIds []string `json:"listenerIds"`
}

//创建7层监听器，https://cloud.tencent.com/document/api/386/9277
func (client *Client) CreateBmL7Listeners(req *CreateBmL7ListenerRequest) ([]string, error) {
	rsp := &CreateBmL7ListenerResponse{}
	err := client.Invoke("CreateBmForwardListeners", req, rsp)
	if err != nil {
		return nil, err
	}
	return rsp.ListenerIds, nil
}

type DescribeBmForwardRulesRequest struct {
	LbId       string   `qcloud_arg:"loadBalancerId"`
	ListenerId string   `qcloud_arg:"listenerId"`
	DomainIds  []string `qcloud_arg:"domainIds"`
}

type ForwardLocation struct {
	Url             string `qcloud_arg:"url"`
	LocationId      string `json:"locationId"`
	SessionExpire   int    `json:"sessionExpire"`
	HealthSwitch    int    `json:"healthSwitch"`
	HttpCheckPath   string `json:"httpCheckPath"`
	HttpCheckDomain string `json:"httpCheckDomain"`
	IntervalTime    int    `json:"intervalTime"`
	HealthNum       int    `json:"healthNum"`
	UnhealthNum     int    `json:"unhealthNum"`
	HttpCode        int    `json:"httpCode"`
	BalanceMode     string `json:"balanceMode"`
	Status          int    `json:"status"`
	AddTimestamp    string `json:"addTimestamp"`

	//for interface DescribeBmForwardRules, this field is nil !!!
	RsList []*L7BackendRs `json:"rsList"`
}

type ForwardRule struct {
	Domain       string             `json:"domain"`
	DomainId     string             `json:"domainId"`
	Status       int                `json:"status"`
	AddTimestamp string             `json:"addTimestamp"`
	Locations    []*ForwardLocation `json:"locations"`
}

type DescribeBmForwardRulesResponse struct {
	TotalCount int            `json:"totalCount"`
	RuleSet    []*ForwardRule `json:"ruleSet"`
}

// https://cloud.tencent.com/document/api/386/9283
func (client *Client) DescribeBmForwardRules(req *DescribeBmForwardRulesRequest) (*DescribeBmForwardRulesResponse, error) {
	rsp := &DescribeBmForwardRulesResponse{}

	err := client.Invoke("DescribeBmForwardRules", req, rsp)
	if err != nil {
		return nil, err
	}
	return rsp, nil
}

type ForwardRulesTA struct {
	Domain          string  `qcloud_arg:"domain"`
	Url             string  `qcloud_arg:"url"`
	SessionExpire   *int    `qcloud_arg:"sessionExpire"`
	HealthSwitch    *int    `qcloud_arg:"healthSwitch"`
	IntervalTime    *int    `qcloud_arg:"intervalTime"`
	HealthNum       *int    `qcloud_arg:"healthNum"`
	UnhealthNum     *int    `qcloud_arg:"unhealthNum"`
	HttpCode        *int    `qcloud_arg:"httpCode"`
	HttpCheckPath   *string `qcloud_arg:"httpCheckPath"`
	HttpCheckDomain *string `qcloud_arg:"httpCheckDomain"`
	BalanceMode     *string `qcloud_arg:"balanceMode"`
}

type CreateBmForwardRulesRequest struct {
	LbId       string            `qcloud_arg:"loadBalancerId"`
	ListenerId string            `qcloud_arg:"listenerId"`
	Rules      []*ForwardRulesTA `qcloud_arg:"rules"`
}

//创建转发规则，https://cloud.tencent.com/document/api/386/9278
func (client *Client) CreateBmForwardRules(req *CreateBmForwardRulesRequest) (int, error) {
	rsp := &LbRequestIdResponse{}
	err := client.Invoke("CreateBmForwardRules", req, rsp)
	if err != nil {
		return 0, err
	}
	return rsp.RequestId, nil
}

type DeleteBmForwardDomainsRequest struct {
	LbId       string   `qcloud_arg:"loadBalancerId"`
	ListenerId string   `qcloud_arg:"listenerId"`
	DomainIds  []string `qcloud_arg:"domainIds"`
}

//删除转发域名，https://cloud.tencent.com/document/api/386/9279
func (client *Client) DeleteBmForwardDomains(req *DeleteBmForwardDomainsRequest) (int, error) {
	rsp := &LbRequestIdResponse{}
	err := client.Invoke("DeleteBmForwardDomains", req, rsp)
	if err != nil {
		return 0, err
	}
	return rsp.RequestId, nil
}

type DeleteBmForwardRulesRequest struct {
	LbId        string   `qcloud_arg:"loadBalancerId"`
	ListenerId  string   `qcloud_arg:"listenerId"`
	DomainId    string   `qcloud_arg:"domainId"`
	LocationIds []string `qcloud_arg:"locationIds"`
}

//删除转发路径，https://cloud.tencent.com/document/api/386/9280
func (client *Client) DeleteBmForwardRules(req *DeleteBmForwardRulesRequest) (int, error) {
	rsp := &LbRequestIdResponse{}
	err := client.Invoke("DeleteBmForwardRules", req, rsp)
	if err != nil {
		return 0, err
	}
	return rsp.RequestId, nil
}

type L7BackendRs struct {
	BindType int    `json:"bindType"`
	RsPort   int    `json:"rsPort"`
	Weight   int    `json:"weight"`
	Status   string `json:"status"`

	InstanceId string `json:"instanceId"`
	LanIp      string `json:"lanIp"` //当bindtype为1时表示黑石物理机的内网IP，当bindType为1时表示虚机IP
}

type L7BackendRsTA struct {
	Port       int    `qcloud_arg:"port"`
	InstanceId string `qcloud_arg:"instanceId"`
	Weight     *int   `qcloud_arg:"weight"`
}

type DescribeBmLocationBackendsRequest struct {
	LbId       string `qcloud_arg:"loadBalancerId"`
	ListenerId string `qcloud_arg:"listenerId"`
	DomainId   string `qcloud_arg:"domainId"`
	LocationId string `qcloud_arg:"locationId"`
}

//获取后端rs列表，https://cloud.tencent.com/document/api/386/9286
func (client *Client) DescribeBmLocationBackends(req *DescribeBmLocationBackendsRequest) ([]*L7BackendRs, error) {
	l7BackendRsList := make([]*L7BackendRs, 0)
	rsp := &BmlbResponse{
		Response: &l7BackendRsList,
	}

	err := client.Invoke("DescribeBmLocationBackends", req, rsp)
	if err != nil {
		return nil, err
	}
	return l7BackendRsList, nil
}

type BindBmLocationInstancesRequest struct {
	LbId       string           `qcloud_arg:"loadBalancerId"`
	ListenerId string           `qcloud_arg:"listenerId"`
	DomainId   string           `qcloud_arg:"domainId"`
	LocationId string           `qcloud_arg:"locationId"`
	Backends   []*L7BackendRsTA `qcloud_arg:"backends"`
}

//绑定物理服务器到七层转发路径，https://cloud.tencent.com/document/api/386/9281
func (client *Client) BindBmLocationInstances(req *BindBmLocationInstancesRequest) error {
	rsp := &NoData{}
	err := client.Invoke("BindBmLocationInstances", req, rsp)
	if err != nil {
		return err
	}
	return nil
}

type UnbindBmLocationInstancesRequest struct {
	LbId       string           `qcloud_arg:"loadBalancerId"`
	ListenerId string           `qcloud_arg:"listenerId"`
	DomainId   string           `qcloud_arg:"domainId"`
	LocationId string           `qcloud_arg:"locationId"`
	Backends   []*L7BackendRsTA `qcloud_arg:"backends"`
}

//解绑物理服务器到七层转发路径，https://cloud.tencent.com/document/api/386/9287
func (client *Client) UnbindBmLocationInstances(req *UnbindBmLocationInstancesRequest) error {
	rsp := &NoData{}
	err := client.Invoke("UnbindBmLocationInstances", req, rsp)
	if err != nil {
		return err
	}
	return nil
}

type DescribeBmForwardListenerInfoRequest struct {
	LbId        string  `qcloud_arg:"loadBalancerId"`
	SearchKey   *string `qcloud_arg:"searchKey"`
	IfGetRsInfo int     `qcloud_arg:"ifGetRsInfo"`
}

//获取负载均衡七层监听器详细信息，https://cloud.tencent.com/document/api/386/9284
func (client *Client) DescribeBmForwardListenerInfo(req *DescribeBmForwardListenerInfoRequest) (*DescribeBmL7ListenersResponse, error) {
	rsp := &DescribeBmL7ListenersResponse{}
	err := client.Invoke("DescribeBmForwardListenerInfo", req, rsp)
	if err != nil {
		return nil, err
	}
	return rsp, nil
}

type ModifyBmForwardListenerRequest struct {
	LbId          string  `qcloud_arg:"loadBalancerId"`
	ListenerId    string  `qcloud_arg:"listenerId"`
	ListenerName  *string `qcloud_arg:"listenerName"`
	SSLMode       *int    `qcloud_arg:"sslMode"`
	CertId        *string `qcloud_arg:"certId"`
	CertName      *string `qcloud_arg:"certName"`
	CertContent   *string `qcloud_arg:"certContent"`
	CertKey       *string `qcloud_arg:"certKey"`
	CertCaId      *string `qcloud_arg:"certCaId"`
	CertCaName    *string `qcloud_arg:"certCaName"`
	CertCaContent *string `qcloud_arg:"certCaContent"`
	Bandwidth     *int    `qcloud_arg:"bandwidth"`
}

//修改负载均衡七层监听器，https://cloud.tencent.com/document/api/386/9273
func (client *Client) ModifyBmForwardListener(req *ModifyBmForwardListenerRequest) (int, error) {
	rsp := &LbRequestIdResponse{}
	err := client.Invoke("ModifyBmForwardListenerRequest", req, rsp)
	if err != nil {
		return 0, err
	}
	return rsp.RequestId, nil
}
