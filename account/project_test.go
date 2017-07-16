package account

import (
	"testing"
)

func TestDescribeProject(t *testing.T) {
	client, err := NewClientFromEnv()
	if err != nil {
		t.Fatal(err)
	}
	args := DescribeProjectArgs{}

	response, err := client.DescribeProject(&args)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%++v", response.Data)

}
