package clb

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestLoadBalancer(t *testing.T) {
	client, err := NewClientFromEnv()
	if err != nil {
		t.Fatal(err)
	}

	createArgs := CreateLoadBalancerArgs{
		LoadBalancerType: 3,
	}

	createResponse, err := client.CreateLoadBalancer(&createArgs)
	if err != nil {
		t.Fatal(err)
	}

	dealId := createResponse.DealIds[0]
	lbId, ok := createResponse.UnLoadBalancerIds[dealId]
	if !ok {
		t.Fatalf("dealId %s not in unLoadBalancerIds", dealId)
	}

	describeArgs := DescribeLoadBalancersArgs{
		LoadBalancerIds: []string{lbId[0]},
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

	newName := fmt.Sprintf("test-lb-v-%d", rand.Int())

	modifyArgs := ModifyLoadBalancerAttributesArgs{
		LoadBalancerId:   lbId[0],
		LoadBalancerName: newName,
	}

	modifyResponse, err := client.ModifyLoadBalancerAttributes(&modifyArgs)
	if err != nil {
		t.Fatal(err)
	}

	requestId := modifyResponse.RequestId

	for {
		time.Sleep(time.Second * 1)
		describeTaskArgs := DescribeLoadBalancersTaskResultArgs{
			RequestId: requestId,
		}
		describeTaskResponse, err := client.DescribeLoadBalancersTaskResult(&describeTaskArgs)
		if err != nil {
			t.Fatal(err)
		}

		if describeTaskResponse.Data.Status == 0 {
			break
		}
	}

	deleteArgs := DeleteLoadBalancersArgs{
		LoadBalancerIds: []string{lbId[0]},
	}

	deleteResponse, err := client.DeleteLoadBalancers(&deleteArgs)
	if err != nil {
		t.Fatal(err)
	}

	deleteRequestId := deleteResponse.RequestId

	for {
		time.Sleep(time.Second * 1)
		describeTaskArgs := DescribeLoadBalancersTaskResultArgs{
			RequestId: deleteRequestId,
		}
		describeTaskResponse, err := client.DescribeLoadBalancersTaskResult(&describeTaskArgs)
		if err != nil {
			t.Fatal(err)
		}

		if describeTaskResponse.Data.Status == 0 {
			break
		}
	}

}
