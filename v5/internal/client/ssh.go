package client

import (
	"chaleaoch.com/golang_demo/internal/entity"
	"chaleaoch.com/golang_demo/internal/task"
)

type StandardSSH struct {
}

func NewStandardSSH(username string, password string, host string) (*StandardSSH, error) {

	return &StandardSSH{}, nil
}
func (s *StandardSSH) Close() error {

	return nil
}

func (s *StandardSSH) ExecuteCommand(command string) (string, error) {
	return "demo out", nil

}

type MipRepo interface {
	GetMip() (*entity.MIp, error)
}

type StandardSSHProvider struct {
	mipRepo MipRepo
}

func NewStandardSSHProvider(mipRepo MipRepo) *StandardSSHProvider {
	return &StandardSSHProvider{
		mipRepo: mipRepo,
	}
}

// GetSSHClient 返回 task.SSHClient 接口, 避免上层依赖具体实现
func (h *StandardSSHProvider) GetSSHClient() (task.SSHClient, error) {
	return NewStandardSSH("username", "password", "host")
}
