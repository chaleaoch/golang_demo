package task

import (
	"fmt"

	"chaleaoch.com/golang_demo/internal/entity"
)

type MIPRepo interface {
	GetByType(mType string) ([]*entity.MIp, error)
	// GetByTypeDc(mType string, dcId int64) ([]*entity.MIp, error)
}

type HealthCheckAlertRepo interface {
	Create(alert *entity.Alert) error
}

type SSHClientProvider interface {
	GetFactory() func(username string, password string, host string) (SSHClient, error)
}

type SSHClient interface {
	ExecuteCommand(command string) (string, error)
	Close() error
}

type HttpClientProvider interface {
	GetClient() httpClient
}

type httpClient interface {
	Get(url string) (string, error)
	Post(url string, body map[string]interface{}) (string, error)
}

type DemoTask struct {
	mIPRepo           MIPRepo
	sshClientProvider SSHClientProvider
}

func NewDemoTask(managementIPRepo MIPRepo, sshClientProvider SSHClientProvider) *DemoTask {
	return &DemoTask{
		mIPRepo:           managementIPRepo,
		sshClientProvider: sshClientProvider,
	}
}

func (q *DemoTask) Run() {
	mIps, _ := q.mIPRepo.GetByType("exampleType")
	// 问题, 这种依赖动态数据的初始化, 要如何做? 有更好的办法吗?
	sshClientFactory := q.sshClientProvider.GetFactory()
	for _, mIp := range mIps {
		sshClient, _ := sshClientFactory(mIp.Username, mIp.Password, mIp.Ip+":"+mIp.Port)
		out, _ := sshClient.ExecuteCommand("exampleCommand")
		fmt.Println(out)
	}
}
