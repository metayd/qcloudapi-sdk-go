package vpc

import (
	"testing"
)

func TestDescribeVpc(t *testing.T) {
	client, err := NewClientFromEnv()
	if err != nil {
		t.Fatal(err)
	}
	args := DescribeVpcExArgs{}

	response, err := client.DescribeVpcEx(&args)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%++v", response.Data)

}
