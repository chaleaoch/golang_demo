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

// SSHClientProvider 由使用方(task) 定义, 返回最小能力接口, 实现方(client) 依赖它实现依赖倒置
type SSHClientProvider interface {
	GetSSHClient() (SSHClient, error)
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
type CmdHander func(SSHClient) string

type DemoTask struct {
	mIPRepo           MIPRepo
	sshClientProvider SSHClientProvider
	cmd1Handler       CmdHander
	cmd2Handler       CmdHander
}

func NewDemoTask(managementIPRepo MIPRepo, sshClientProvider SSHClientProvider, cmd1Handler, cmd2Handler CmdHander) *DemoTask {
	return &DemoTask{
		mIPRepo:           managementIPRepo,
		sshClientProvider: sshClientProvider,
		cmd1Handler:       cmd1Handler,
		cmd2Handler:       cmd2Handler,
	}
}

func (q *DemoTask) Run() {
	// mIps, _ := q.mIPRepo.GetByType("exampleType")
	sshClient, _ := q.sshClientProvider.GetSSHClient()

	out1 := q.cmd1Handler(sshClient)
	out2 := q.cmd2Handler(sshClient)
	fmt.Println(out1, out2)

}
