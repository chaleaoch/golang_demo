package client

type StandardSSH struct {
}

func NewStandardSSH() (*StandardSSH, error) {

	return &StandardSSH{}, nil
}
func (s *StandardSSH) Close() error {

	return nil
}

func (s *StandardSSH) ExecuteCommand(command string) (string, error) {
	return "demo out", nil
}

func (s *StandardSSH) Conenct(username string, password string, host string) error {
	return nil
}

// type StandardSSHProvider struct {
// }

// func NewStandardSSHProvider() *StandardSSHProvider {
// 	return &StandardSSHProvider{}
// }

// func (h *StandardSSHProvider) GetFactory() func(username string, password string, host string) (task.SSHClient, error) {
// 	return func(username string, password string, host string) (task.SSHClient, error) {
// 		return NewStandardSSH(username, password, host)
// 	}
// }
