package main

import (
	"chaleaoch.com/golang_demo/internal/client"
	"chaleaoch.com/golang_demo/internal/repo"
	"chaleaoch.com/golang_demo/internal/task"
)

func main() {
	// sshProvider := client.NewStandardSSHProvider()
	m2IPRepo := repo.NewMipRepo()
	sshClient, _ := client.NewStandardSSH()
	qct := task.NewDemoTask(m2IPRepo, sshClient)
	qct.Run()

}
