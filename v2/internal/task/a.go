package task

import (
	"fmt"

	"chaleaoch.com/golang_demo/internal/entity"
	"chaleaoch.com/golang_demo/internal/handler"
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
	sshClientFactory := q.sshClientProvider.GetFactory()
	for _, mIp := range mIps {
		sshClient, _ := sshClientFactory(mIp.Username, mIp.Password, mIp.Ip+":"+mIp.Port)
		// 修改: 这里将所有的cmd命令都隔离出去了, 为了实现复用(如何复用这里没有展示).
		// 目前来看一切都还好. 但是我想把这些handler 注入进来... 参考v3
		// 问题:
		// 这种注入是有必要的吗? 哪种场景需要用接口做隔离? 我粗浅的理解, 需要单测mock的时候, 也就是IO的时候做隔离, 如果是类似v3的那种,
		// 相当于是业务层隔离, 就有点太复杂了... 没做过复杂的业务, 找不到答案.
		out1 := handler.Cmd1Handler(sshClient)
		out2 := handler.Cmd2Handler(sshClient)
		fmt.Println(out1, out2)

	}
}
