package handler

import "chaleaoch.com/golang_demo/internal/task"

// 问题,
func Cmd1Handler(sshClient task.SSHClient) string {
	out, _ := sshClient.ExecuteCommand("cmd1")
	// 干点别的...100行
	return out
}

func Cmd2Handler(sshClient task.SSHClient) string {
	out, _ := sshClient.ExecuteCommand("cmd1")
	// 干点别的...100行
	return out
}
