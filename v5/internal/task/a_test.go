package task

import (
	"testing"

	"chaleaoch.com/golang_demo/internal/entity"
)

// fakeSSH implements SSHClient for tests
type fakeSSH struct {
	cmds []string
	out  string
	err  error
}

func (f *fakeSSH) ExecuteCommand(c string) (string, error) {
	f.cmds = append(f.cmds, c)
	return f.out, f.err
}
func (f *fakeSSH) Close() error { return nil }

type fakeProvider struct{ c SSHClient }

func (p *fakeProvider) GetSSHClient() (SSHClient, error) { return p.c, nil }

// minimal stub for MIPRepo
type fakeRepo struct{}

func (r *fakeRepo) GetByType(t string) ([]*entity.MIp, error) { return nil, nil }

// compile-time interface assertions
var _ SSHClient = (*fakeSSH)(nil)
var _ SSHClientProvider = (*fakeProvider)(nil)

func TestDemoTask_Run(t *testing.T) {
	fssh := &fakeSSH{out: "ok"}
	provider := &fakeProvider{c: fssh}
	repo := &fakeRepo{}

	cmd1 := func(c SSHClient) string { o, _ := c.ExecuteCommand("cmd1"); return o }
	cmd2 := func(c SSHClient) string { o, _ := c.ExecuteCommand("cmd2"); return o }

	task := NewDemoTask(repo, provider, cmd1, cmd2)
	task.Run()

	if len(fssh.cmds) != 2 {
		t.Fatalf("expected 2 cmds, got %d", len(fssh.cmds))
	}
}
