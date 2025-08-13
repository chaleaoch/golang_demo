package handler

type SSHClient interface {
	ExecuteCommand(command string) (string, error)
	Close() error
}

func Cmd1Handler(sshClient SSHClient) string {
	out, _ := sshClient.ExecuteCommand("cmd1")
	// 干点别的...100行
	return out
}

func Cmd2Handler(sshClient SSHClient) string {
	out, _ := sshClient.ExecuteCommand("cmd1")
	// 干点别的...100行
	return out
}
