package cbs

import (
	"context"
	"time"
)

func (client *Client) CreateCbsStorageTask(args *CreateCbsStorageArgs) ([]string, error) {
	storageIds, err := client.CreateCbsStorage(args)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*180)
	defer cancel()
	ticker := time.NewTicker(TaskCheckInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-ticker.C:
			args := DescribeCbsStorageArgs{
				StorageIds: &storageIds,
			}
			response, err := client.DescribeCbsStorage(&args)
			if err == nil && len(response.StorageSet) == len(storageIds) {
				return storageIds, nil
			}
		}
	}
}

func (client *Client) AttachCbsStorageTask(storageId string, uInstanceId string) error {

	return WaitUntilDone(
		func() (string, error) {
			_, err := client.AttachCbsStorage([]string{storageId}, uInstanceId)
			if err != nil {
				return "", err
			}
			return storageId, nil
		},
		func(info *StorageSet) bool {
			return info.Attached == 1
		},
		client,
	)
}

func (client *Client) DetachCbsStorageTask(storageId string) error {

	return WaitUntilDone(
		func() (string, error) {
			_, err := client.DetachCbsStorage([]string{storageId})
			if err != nil {
				return "", err
			}
			return storageId, nil
		},
		func(info *StorageSet) bool {
			return info.Attached == 0
		},
		client,
	)
}
