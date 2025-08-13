package main

import (
	"chaleaoch.com/golang_demo/internal/client"
	"chaleaoch.com/golang_demo/internal/handler"
	"chaleaoch.com/golang_demo/internal/repo"
	"chaleaoch.com/golang_demo/internal/task"
)

func main() {
	// 手动注入部分 略
	sshProvider := client.NewStandardSSHProvider()
	m2IPRepo := repo.NewMipRepo()
	// 修改: 这里做了适配, 否则, 会出现循环导入问题或者是类型不匹配问题
	// 问题, 这么做合理吗? 有其他办法吗? 用接口怎么做, 怎么做是符合最佳实践的? 写到这里脑子已经炸了ww
	cmd1Handler := func(sshClient task.SSHClient) string {
		return handler.Cmd1Handler(sshClient)
	}
	cmd2Handler := func(sshClient task.SSHClient) string {
		return handler.Cmd2Handler(sshClient)
	}
	qct := task.NewDemoTask(m2IPRepo, sshProvider, cmd1Handler, cmd2Handler)
	qct.Run()

}
