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
	mIps, _ := q.mIPRepo.GetByType("exampleType")
	sshClientFactory := q.sshClientProvider.GetFactory()
	for _, mIp := range mIps {
		sshClient, _ := sshClientFactory(mIp.Username, mIp.Password, mIp.Ip+":"+mIp.Port)
		out1 := q.cmd1Handler(sshClient)
		out2 := q.cmd2Handler(sshClient)
		fmt.Println(out1, out2)
	}
}
