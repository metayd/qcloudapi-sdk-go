package ccs

import (
	"testing"
)

func TestDescribeCluster(t *testing.T) {
	client, err := NewClientFromEnv()
	if err != nil {
		t.Fatal(err)
	}

	args := DescribeClusterArgs{
	}

	response, err := client.DescribeCluster(&args)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%++v", response)

}

func StringPtr(s string) *string {
	return &s
}

func TestAddClusterInstances(t *testing.T) {
	client, err := NewClientFromEnv()
	if err != nil {
		t.Fatal(err)
	}

	args := AddClusterInstancesArgs{
		ClusterId:"cls-18zs6ca9",
		InstanceType:"S2.SMALL1",
		Bandwidth:0,
		BandwidthType:PayByTraffic,
		SubnetId:"subnet-s1jz1ycx",
		StorageSize:50,
		RootSize:50,
		CvmType:PayByMonth,
		Period:1,
		ZoneId:ApShanghai_1,
		GoodsNum:1,
		RootType:LOCAL_BASIC,
		StorageType:LOCAL_BASIC,
		SgId:"sg-ni5ob4js",
		//Password:StringPtr("aa123456"),
		KeyId:StringPtr("skey-44213ztl"),
		MountTarget:"/data",
		DockerGraphPath:"/data/docker",
		UserScript:"IyEvYmluL2Jhc2gKZWNobyBmdWNreW91Cg==",


	}

	response, err := client.AddClusterInstances(&args)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%++v", response)

}
