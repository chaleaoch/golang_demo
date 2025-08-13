package main

import (
	"chaleaoch.com/golang_demo/internal/client"
	"chaleaoch.com/golang_demo/internal/repo"
	"chaleaoch.com/golang_demo/internal/task"
)

func main() {
	// 手动注入
	sshProvider := client.NewStandardSSHProvider()
	m2IPRepo := repo.NewMipRepo()
	t := task.NewDemoTask(m2IPRepo, sshProvider)
	t.Run()

}
