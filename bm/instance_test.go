package bm

import (
	"github.com/dbdd4us/qcloudapi-sdk-go/common"
	"testing"
)

func TestDescribeDevice(t *testing.T) {

	client, _ := NewClientFromEnv()

	lanIps := []string{"10.0.0.4"}
	req := DescribeDeviceArgs{
		LanIps: &lanIps,
	}

	if devInfo, err := client.DescribeDevice(&req); err != nil {
		t.Error(err.Error())
	} else {
		t.Logf("DescribeDevice Pass devInfo=%v", devInfo)
	}

}


func TestSubnetIp(t *testing.T) {
	client, _ := NewClientFromEnv()
    taskId,err := client.RegisterContainerSubnetIp("vpc-muinpf9p","subnet-c6bzyq4a")
	if err  != nil {
		t.Error(err.Error())
		return 
	}
	err = client.WaitUntiTaskDone(taskId)
	if err  != nil {
        t.Error(err.Error())
		return 
	}

	taskId,err := client.ReleaseContainerSubnetIp("vpc-muinpf9p","subnet-c6bzyq4a")
	if err  != nil {
		t.Error(err.Error())
		return 
	}

}
