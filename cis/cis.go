package cis

import ()

const ()

type CisResponse struct {
	Response interface{} `json:"Response"`
}

type CreateContainerRequest struct {
	Version       string      `qcloud_arg:"Version"`
	Zone          string      `qcloud_arg:"Zone"`
	VpcId         string      `qcloud_arg:"VpcId"`
	SubnetId      string      `qcloud_arg:"SubnetId"`
	InstanceName  string      `qcloud_arg:"InstanceName"`
	RestartPolicy string      `qcloud_arg:"RestartPolicy"`
	Containers    []Container `qcloud_arg:"Containers"`
}

type Container struct {
	Name            string        `qcloud_arg:"Name"`
	Command         string        `qcloud_arg:"Command"`
	Args            []string      `qcloud_arg:"Args"`
	EnvironmentVars []Environment `qcloud_arg:"EnvironmentVars"`
	Image           string        `qcloud_arg:"Image"`
	Cpu             float64       `qcloud_arg:"Cpu"`
	Memory          float64       `qcloud_arg:"Memory"`
}

type Environment struct {
	Name  string `qcloud_arg:"Name"`
	Value string `qcloud_arg:"Value"`
}

type CreateContainerResponse struct {
	RequestId   string `json:"RequestID"`
	ContainerID string `json:"InstanceID"`
}

func (client *Client) CreateContainer(args *CreateContainerRequest) (*CisResponse, error) {
	realRsp := &CreateContainerResponse{}
	cisResponse := &CisResponse{
		Response: realRsp,
	}
	err := client.Invoke("CreateContainerInstance", args, cisResponse)
	if err != nil {
		return &CisResponse{}, err
	}
	return cisResponse, nil
}

type GetContainerInstanceRequest struct {
	Version      string `qcloud_arg:"Version"`
	InstanceName string `qcloud_arg:"InstanceName"`
}

type GetContainerInstanceResponse struct {
	RequestId         string            `json:"RequestID"`
	ContainerInstance ContainerInstance `json:"ContainerInstance"`
}

type ContainerInstance struct {
	InstanceId    string          `json:"InstanceId"`
	InstanceName  string          `json:"InstanceName"`
	VpcId         string          `json:"VpcId"`
	SubnetId      string          `json:"SubnetId"`
	CreateTime    string          `json:"CreateTime"`
	RestartPolicy string          `json:"RestartPolicy"`
	LanIp         string          `json:"LanIp"`
	State         string          `json:"State"`
	StartTime     string          `json:"StartTime"`
	Containers    []ContainerInfo `json:"Containers"`
}

type ContainerInfo struct {
	Name            string          `json:"Name"`
	Command         string          `json:"Command"`
	Args            []string        `json:"Args"`
	EnvironmentVars []Environment   `json:"EnvironmentVars"`
	Image           string          `json:"Image"`
	Cpu             float64         `json:"Cpu"`
	Memory          float64         `json:"Memory"`
	RestartCount    uint64          `json:"RestartCount"`
	CurrentState    ContainerStatus `json:"CurrentState"`
	PreviousState   ContainerStatus `json:"PreviousState"`
}

type ContainerStatus struct {
	Reason     string
	State      string
	StartTime  string
	FinishTime string
	ExitCode   int32
}

func (client *Client) GetContainerInstance(args *GetContainerInstanceRequest) (*CisResponse, error) {
	realRsp := &GetContainerInstanceResponse{}
	cisResponse := &CisResponse{
		Response: realRsp,
	}
	err := client.Invoke("DescribeContainerInstance", args, cisResponse)
	if err != nil {
		return &CisResponse{}, err
	}
	return cisResponse, nil
}

type ListContainerInstanceRequest struct {
	Version string `qcloud_arg:"Version"`
}

type ListContainerInstanceResponse struct {
	RequestId             string              `json:"RequestID"`
	ContainerInstanceList []ContainerInstance `json:"ContainerInstanceList"`
}

func (client *Client) ListContainerInstance(args *ListContainerInstanceRequest) (*CisResponse, error) {
	realRsp := &ListContainerInstanceResponse{}
	cisResponse := &CisResponse{
		Response: realRsp,
	}
	err := client.Invoke("DescribeContainerInstances", args, cisResponse)
	if err != nil {
		return &CisResponse{}, err
	}
	return cisResponse, nil
}

type GetContainerLogRequest struct {
	Version       string `qcloud_arg:"Version"`
	InstanceName  string `qcloud_arg:"InstanceName"`
	ContainerName string `qcloud_arg:"ContainerName"`
	Tail          uint64 `qcloud_arg:"Tail"`
}

type GetContainerLogResponse struct {
	RequestId        string         `json:"RequestID"`
	ContainerLogList []ContainerLog `json:"ContainerLogList"`
}

type ContainerLog struct {
	Name string `json:"Name"`
	Log  string `json:"Log"`
	Time string `json:"Time"`
}

func (client *Client) GetContainerLog(args *GetContainerLogRequest) (*CisResponse, error) {
	realRsp := &GetContainerLogResponse{}
	cisResponse := &CisResponse{
		Response: realRsp,
	}
	err := client.Invoke("DescribeContainerLog", args, cisResponse)
	if err != nil {
		return &CisResponse{}, err
	}
	return cisResponse, nil
}

type DeleteContainerInstanceRequest struct {
	Version      string `qcloud_arg:"Version"`
	InstanceName string `qcloud_arg:"InstanceName"`
}

type DeleteContainerInstanceResponse struct {
	RequestId string `json:"RequestID"`
	Msg       string `json:"Msg"`
}

func (client *Client) DeleteContainerInstance(args *DeleteContainerInstanceRequest) (*CisResponse, error) {
	realRsp := &DeleteContainerInstanceResponse{}
	cisResponse := &CisResponse{
		Response: realRsp,
	}
	err := client.Invoke("DeleteContainerInstance", args, cisResponse)
	if err != nil {
		return &CisResponse{}, err
	}
	return cisResponse, nil
}
