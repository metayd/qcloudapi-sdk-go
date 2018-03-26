package bmlb

import (
	"errors"
	"github.com/dbdd4us/qcloudapi-sdk-go/common"
	"strings"
	"testing"
	"time"
)

func getLbInfoByLbId(lbId string, client *Client, t *testing.T) int {
	lbIds := []string{lbId}
	req := &DescribeBmLbReqest{
		LbIds: &lbIds,
	}

	lbSet, err := client.DescribeBmLoadBalancers(req)
	if err != nil {
		t.Error(err.Error())
		return -1
	}

	t.Logf("lbSet=%v\n", lbSet)
	return (*lbSet)[0].Status
}

func waitLbCreateOk(client *Client, t *testing.T, lbId string) {
	for i := 0; i < 60; i++ {
		status := getLbInfoByLbId(lbId, client, t)
		if status == BM_LB_STATE_RUNNING {
			return
		}
		time.Sleep(time.Second * 5)
	}
}

func testBackEndRs(lbId, listenerId string, client *Client, t *testing.T) {
	//bind后端rs
	instanceIds := []string{"cpm-an5a9wv4"}
	backends := []BackendRs{}
	for _, instance := range instanceIds {
		backend := BackendRs{
			Port:       9090,
			InstanceId: instance,
			Weight:     10,
		}
		backends = append(backends, backend)
	}

	req := &BmBindL4LisenerRsRequest{
		LbId:       lbId,
		ListenerId: listenerId,
		Backends:   backends,
	}

	taskId, err := client.BindBmL4ListenerRs(req)
	if err != nil {
		t.Error(err.Error())
		return
	}

	err = client.WaitUntiTaskDone(taskId, 60)
	if err != nil {
		t.Error(err.Error())
		return
	}

	//获取后端rs
	descReq := &DescribeBm4BackendRsRequest{
		LbId:       lbId,
		ListenerId: listenerId,
	}

	backendsList, err := client.DescribeBmL4ListenerBackend(descReq)
	if err != nil {
		t.Error(err.Error())
		return
	} else {
		t.Logf("DescribeBmL4ListenerBackend backends=%v", backendsList)
	}

	//删除后端rs
	unbindBackends := []BmUnbindRs{}
	for _, instanceId := range instanceIds {
		backend := BmUnbindRs{
			Port:       9090,
			InstanceId: instanceId,
		}
		unbindBackends = append(unbindBackends, backend)
	}

	unbindReq := &UnbindBmListenerRsRequest{
		LbId:       lbId,
		ListenerId: listenerId,
		Backends:   unbindBackends,
	}

	taskId, err = client.UnbindBmL4ListenerRs(unbindReq)
	if err != nil {
		t.Error(err.Error())
		return
	}

	err = client.WaitUntiTaskDone(taskId, 60)
	if err != nil {
		t.Error(err.Error())
		return
	}

}

func convertToLbProtocol(protocol string) string {
	upperProtocol := strings.ToUpper(protocol)
	switch upperProtocol {
	case "TCP":
		return "tcp"
	case "UDP":
		return "udp"
	default:
		return "unknown"
	}
}

func getListenerId(client *Client, t *testing.T, lbId string, protocol string, port int) (string, error) {
	listenerreq := &DescribeBmLbListenerRequest{
		LbId: lbId,
	}

	rsp, err := client.DescribeBmListeners(listenerreq)
	if err != nil {
		return "", err
	}

	for _, listener := range rsp.ListenerSet {
		if convertToLbProtocol(listener.Protocol) == convertToLbProtocol(protocol) && port == listener.LbPort {
			return listener.Id, nil
		}
	}

	return "", errors.New("not found listenerId")
}

func testListener(lbId string, client *Client, t *testing.T) {
	//创建listener
	createListener := BmListenerCreateDetail{
		LbPort:   8890,
		Protocol: "tcp",
		Name:     "listnenerBrank",
	}

	listeners := []BmListenerCreateDetail{createListener}
	req := &CreateBmListenerRequest{
		LbId:      lbId,
		Listeners: &listeners,
	}

	taskId, err := client.CreateBmListeners(req)
	if err != nil {
		t.Error(err.Error())
		return
	}
	if err = client.WaitUntiTaskDone(taskId, 60); err != nil {
		t.Error(err.Error())
		return
	}

	//获取listener
	listnerId, err := getListenerId(client, t, lbId, "tcp", 8890)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	testBackEndRs(lbId, listnerId, client, t)

	//删除listener
	delListenerIds := []string{listnerId}
	delReq := &DeleteBmListenersRequest{
		LbId:        lbId,
		ListenerIds: &delListenerIds,
	}

	taskId, err = client.DeleteBmListeners(delReq)
	if err != nil {
		t.Error(err.Error())
		return
	}

	if err = client.WaitUntiTaskDone(taskId, 60); err != nil {
		t.Error(err.Error())
		return
	}

}

func TestLb(t *testing.T) {
	unVpcId := "vpc-9jso7qmq"
	unSubnetId := "subnet-dv07fj1p"
	client, _ := NewClientFromEnv()
	goodsNum := 1

	req := &CreateBmLbReq{
		UnVpcId:    unVpcId,
		LbType:     "internal",
		UnSubnetId: &unSubnetId,
		GoodsNum:   &goodsNum,
	}

	lbIds, err := client.CreateBmLoadBalancer(req)
	if err != nil {
		t.Errorf("createBmLoadBalancer failed err=%v", err)
		return
	} else {
		t.Logf("createBmLoadBalancer ok,lbIds=%v", lbIds)
	}

	waitLbCreateOk(client, t, (*lbIds)[0])

	//修改LB
	lbName := "ccs_lb_brankbao"
	modifyReq := &BmModifyLbAttributeRequest{
		LbId:   (*lbIds)[0],
		LbName: &lbName,
	}
	taskId, err := client.ModifyBmLoadBalancerAttributes(modifyReq)
	if err != nil {
		t.Errorf("ModifyBmLb failed err=%v", err)
	}
	err = client.WaitUntiTaskDone(taskId, 60)
	if err != nil {
		t.Errorf("ModifyBmLb failed err=%v", err)
	}

	getLbInfoByLbId((*lbIds)[0], client, t)

	testListener((*lbIds)[0], client, t)

	//删除LB
	delLbReq := &DeleteBmLbRequest{
		LbId: (*lbIds)[0],
	}

	taskId, err = client.DeleteBmLoadBalancers(delLbReq)
	if err != nil {
		t.Errorf("DeleteBmLb failed err=%v", err)
	}
	err = client.WaitUntiTaskDone(taskId, 60)
	if err != nil {
		t.Errorf("DeleteBmLb failed err=%v", err)
	}

}
