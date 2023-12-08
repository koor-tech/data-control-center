package ancientt

import (
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

	stdout io.ReadCloser
	stderr io.ReadCloser
}

func NewRun(workingDir string, command string, args []string) (IRun, error) {
	r := &Run{
		cmd: exec.Command(command, args...),
	}

	r.cmd.Dir = workingDir

	stdout, err := r.cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}
	r.stdout = stdout

	stderr, err := r.cmd.StderrPipe()
	if err != nil {
		return nil, err
	}
	r.stderr = stderr

	return r, nil
}

func (r *Run) Start() error {
	if err := r.cmd.Start(); err != nil {
		return err
	}

	return nil
}

func (r *Run) IsRunning() bool {
	return !(r.cmd == nil || (r.cmd.ProcessState != nil && r.cmd.ProcessState.Exited()) || r.cmd.Process == nil)
}

func (r *Run) Stop() error {
	if r.cmd == nil || r.cmd.Process == nil {
		return nil
	}

	return r.cmd.Process.Kill()
}

func (r *Run) GetLogs() (string, error) {
	multi := io.MultiReader(r.stdout, r.stderr)
	out, err := io.ReadAll(multi)
	if err != nil {
		return "", err
	}

	return string(out), nil
}
