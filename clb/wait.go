package clb

import (
	"context"
	"time"
)

const (
	TaskCheckInterval = time.Second * 1

	TaskSuccceed = 0
	TaskFailed = 1
	TaskRunning = 2

	TaskStatusUnknown = 9
)

type Task struct {
	requestId int
}

func NewTask(requestId int) Task {
	return Task{requestId: requestId}
}

func (task *Task) WaitUntilDone(ctx context.Context, client *Client) (int, error) {
	ticker := time.NewTicker(TaskCheckInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return TaskStatusUnknown, ctx.Err()
		case <- ticker.C:
			args := DescribeLoadBalancersTaskResultArgs{
				RequestId: task.requestId,
			}
			response, err := client.DescribeLoadBalancersTaskResult(&args)
			if err != nil {
				return TaskStatusUnknown, err
			}
			if response.Data.Status != TaskRunning {
				return response.Data.Status, nil
			}
		}
	}
}
