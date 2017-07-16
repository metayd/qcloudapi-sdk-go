package cvm

import (
	"testing"
)

func TestDescribeProject(t *testing.T) {
	client, err := NewClientFromEnv()
	if err != nil {
		t.Fatal(err)
	}
	args := DescribeInstancesArgs{
		Version:"2017-03-12",
	}

	response, err := client.DescribeInstances(&args)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%++v", response.Response.InstanceSet)

}
