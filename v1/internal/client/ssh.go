package client

import (
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

type StandardSSHProvider struct {
}

func NewStandardSSHProvider() *StandardSSHProvider {
	return &StandardSSHProvider{}
}

// 问题: ssh.go 导入了task.SSHClient是否合理, 如果不合理, 这里应该怎么做?
func (h *StandardSSHProvider) GetFactory() func(username string, password string, host string) (task.SSHClient, error) {
	return func(username string, password string, host string) (task.SSHClient, error) {
		return NewStandardSSH(username, password, host)
	}
}
