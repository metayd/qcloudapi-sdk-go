package cbs

import (
	"testing"
)

func TestCreateStorage(t *testing.T) {
	client, err := NewClientFromEnv()
	if err != nil {
		t.Fatal(err)
	}

	args := &CreateCbsStorageArgs{
		StorageType: StorageTypeCloudBasic,
		StorageSize: 10,
		PayMode:     PayModePrePay,
		Period:      1,
		GoodsNum:    1,
		Zone:        "ap-guangzhou-2",
	}

	storageIds, err := client.CreateCbsStorageTask(args)
	if err != nil {
		t.Fatal(err)
	}
	if len(storageIds) != 1 {
		t.Fatalf("err:len(storageIds)=%d", len(storageIds))
	}
	t.Logf("%++v", storageIds)
}

func TestAttachDetachStorage(t *testing.T) {
	client, err := NewClientFromEnv()
	if err != nil {
		t.Fatal(err)
	}

	err = client.AttachCbsStorageTask("disk-am1x7yys", "ins-psfsw27e")
	if err != nil {
		t.Fatal(err)
	}

	err = client.DetachCbsStorageTask("disk-am1x7yys")
	if err != nil {
		t.Fatal(err)
	}

}
