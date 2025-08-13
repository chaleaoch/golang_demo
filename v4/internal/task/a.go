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
	Conenct(username string, password string, host string) error
}

type HttpClientProvider interface {
	GetClient() httpClient
}

type httpClient interface {
	Get(url string) (string, error)
	Post(url string, body map[string]interface{}) (string, error)
}

type DemoTask struct {
	mIPRepo   MIPRepo
	sshClient SSHClient
}

func NewDemoTask(managementIPRepo MIPRepo, sshClient SSHClient) *DemoTask {
	return &DemoTask{
		mIPRepo:   managementIPRepo,
		sshClient: sshClient,
	}
}

func (q *DemoTask) Run() {
	mIps, _ := q.mIPRepo.GetByType("exampleType")

	for _, mIp := range mIps {
		_ = q.sshClient.Conenct(mIp.Username, mIp.Password, mIp.Ip+":"+mIp.Port)
		out1 := handler.Cmd1Handler(q.sshClient)
		out2 := handler.Cmd2Handler(q.sshClient)
		fmt.Println(out1, out2)

	}
}
