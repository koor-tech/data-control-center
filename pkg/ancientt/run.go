package ancientt

import (
	"bytes"
	"io"
	"os/exec"
)

type IRun interface {
	Start() error
	IsRunning() bool
	Stop() error
	GetLogs() (string, error)
}

type Run struct {
	cmd *exec.Cmd

	buf  *bytes.Buffer
	logs string
}

func NewRun(workingDir string, command string, args []string) (IRun, error) {
	r := &Run{
		cmd:  exec.Command(command, args...),
		logs: "",
	}

	r.cmd.Dir = workingDir

	r.buf = new(bytes.Buffer)
	r.cmd.Stdout = r.buf
	r.cmd.Stderr = r.buf

	return r, nil
}

func (r *Run) Start() error {
	if err := r.cmd.Start(); err != nil {
		return err
	}

	return nil
}

func (r *Run) IsRunning() bool {
	return r.cmd != nil && ((r.cmd.ProcessState != nil && !r.cmd.ProcessState.Exited()) || r.cmd.Process != nil)
}

func (r *Run) Stop() error {
	if r.cmd == nil || r.cmd.Process == nil {
		return nil
	}

	return r.cmd.Process.Kill()
}

func (r *Run) GetLogs() (string, error) {
	out, err := io.ReadAll(r.buf)
	if err != nil {
		return "", err
	}

	r.logs = string(out)

	return r.logs, nil
}
