package clb

import (
	"context"
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestLoadBalancerListeners(t *testing.T) {
	client, err := NewClientFromEnv()
	if err != nil {
		t.Fatal(err)
	}

	createLoadBalancerArgs := CreateLoadBalancerArgs{
		LoadBalancerType: 3,
	}
	lb, err := client.CreateLoadBalancer(&createLoadBalancerArgs)
	if err != nil {
		t.Fatal(err)
	}

	dealId := lb.DealIds[0]
	lbId := lb.UnLoadBalancerIds[dealId][0]

	describeArgs := DescribeLoadBalancersArgs{
		LoadBalancerIds: &[]string{lbId},
	}

	for {
		time.Sleep(time.Second * 1)
		describeResponse, err := client.DescribeLoadBalancers(&describeArgs)
		if err != nil {
			t.Fatal(err)
		}
		if len(describeResponse.LoadBalancerSet) > 0 {
			break
		}
	}

	createListenerArgs := CreateLoadBalancerListenersArgs{
		LoadBalancerId: lbId,
		Listeners: []CreateListenerOpts{
			{
				LoadBalancerPort: 9000,
				InstancePort:     9000,
				Protocol:         3,
			},
		},
	}

	createListenerResponse, err := client.CreateLoadBalancerListeners(&createListenerArgs)
	if err != nil {
		t.Fatal(err)
	}

	task := NewTask(createListenerResponse.RequestId)
	task.WaitUntilDone(context.Background(), client)

	describeLoadBalancerListenersArgs := DescribeLoadBalancerListenersArgs{
		LoadBalancerId: lbId,
	}

	lbListeners, err := client.DescribeLoadBalancerListeners(&describeLoadBalancerListenersArgs)
	if err != nil {
		t.Fatal(err)
	}

	newName := fmt.Sprintf("lb-listener-v-%d", rand.Int())

	modifyListenerArgs := ModifyLoadBalancerListenerArgs{
		LoadBalancerId: lbId,
		ListenerId:     lbListeners.ListenerSet[0].UnListenerId,
		ListenerName:   &newName,
	}

	modifyListenerResponse, err := client.ModifyLoadBalancerListener(&modifyListenerArgs)
	if err != nil {
		t.Fatal(err)
	}

	task = NewTask(modifyListenerResponse.RequestId)
	task.WaitUntilDone(context.Background(), client)

	deleteArgs := DeleteLoadBalancerListenersArgs{
		LoadBalancerId: lbId,
		ListenerIds:    createListenerResponse.ListenerIds,
	}

	deleteListenerResponse, err := client.DeleteLoadBalancerListeners(&deleteArgs)
	if err != nil {
		t.Fatal(err)
	}

	task = NewTask(deleteListenerResponse.RequestId)
	_, err = task.WaitUntilDone(context.Background(), client)
	if err != nil {
		t.Fatal(err)
	}
}
