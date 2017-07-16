package clb

import (
	"testing"
)

func TestDescribeLoadBalancers(t *testing.T) {
	client, err := NewClientFromEnv()
	if err != nil {
		t.Fatal(err)
	}

	args := DescribeLoadBalancersArgs{
		LoadBalancerType: 3,
		//LoadBalancerIds: LoadBalancerIds{"i-1", "i-2",},
	}

	loadBalancers, err := client.DescribeLoadBalancers(&args)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%++v", loadBalancers.LoadBalancerSet)
}

