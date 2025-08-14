package main

import (
	"chaleaoch.com/golang_demo/internal/client"
	"chaleaoch.com/golang_demo/internal/handler"
	"chaleaoch.com/golang_demo/internal/repo"
	"chaleaoch.com/golang_demo/internal/task"
)

func main() {
	m2IPRepo := repo.NewMipRepo()
	sshProvider := client.NewStandardSSHProvider(m2IPRepo)

	cmd1Handler := func(sshClient task.SSHClient) string {
		return handler.Cmd1Handler(sshClient)
	}
	cmd2Handler := func(sshClient task.SSHClient) string {
		return handler.Cmd2Handler(sshClient)
	}
	qct := task.NewDemoTask(m2IPRepo, sshProvider, cmd1Handler, cmd2Handler)
	qct.Run()

}
