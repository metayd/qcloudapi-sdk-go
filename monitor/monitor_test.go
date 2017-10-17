package monitor

import (
	"testing"

	"github.com/dbdd4us/qcloudapi-sdk-go/cvm"
)

const (
	CvmLanOutTraffic = "lan_outtraffic"
)

func TestGetMonitorData(t *testing.T) {
	client, err := NewClientFromEnv()
	if err != nil {
		t.Fatal(err)
	}

	cvmClient, err := cvm.NewClientFromEnv()
	if err != nil {
		t.Fatal(err)
	}

	cvmList, err := cvmClient.DescribeInstances(&cvm.DescribeInstancesArgs{
		Version: "2017-03-12",
	})
	if err != nil {
		t.Fatal(err)
	}

	if len(cvmList.InstanceSet) <= 0 {
		t.Fatal("no instance available")
	}

	getMonitorDataArgs := GetMonitorDataArgs{
		Namespace:  NameSpaceQceCvm,
		MetricName: CvmLanOutTraffic,
		Dimensions: []Dimension{
			{
				Name:  "unInstanceId",
				Value: cvmList.InstanceSet[0].InstanceID,
			},
		},
	}

	_, err = client.GetMonitorData(&getMonitorDataArgs)
	if err != nil {
		t.Fatal(err)
	}

	batchGetMonitorDataArgs := BatchGetMonitorDataArgs{
		Namespace:  NameSpaceQceCvm,
		MetricName: CvmLanOutTraffic,
		Batch: []Batch{
			{
				Dimensions: []Dimension{
					{
						Name:  "unInstanceId",
						Value: cvmList.InstanceSet[0].InstanceID,
					},
				},
			},
		},
	}

	_, err = client.BatchGetMonitorData(&batchGetMonitorDataArgs)
	if err != nil {
		t.Fatal(err)
	}
}
