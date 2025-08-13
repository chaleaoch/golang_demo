package main

import (
	"chaleaoch.com/golang_demo/internal/client"
	"chaleaoch.com/golang_demo/internal/repo"
	"chaleaoch.com/golang_demo/internal/task"
)

func main() {
	// 手动注入部分 略
	sshProvider := client.NewStandardSSHProvider()
	m2IPRepo := repo.NewMipRepo()
	qct := task.NewDemoTask(m2IPRepo, sshProvider)
	qct.Run()

}
