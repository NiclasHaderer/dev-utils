package process

import (
	"errors"
	"os"
	"os/exec"
)

type ExecutionError interface {
	Error() string
	Command() []string
	Status() int
}

type ExecutionErrorImpl struct {
	Code int
	cmd  []string
	Err  error
}

func (e *ExecutionErrorImpl) Error() string {
	return e.Err.Error()
}

func (e *ExecutionErrorImpl) Command() []string {
	return e.cmd
}

func (e *ExecutionErrorImpl) Status() int {
	return e.Code
}

func Run(command []string, pipeStdout ...bool) ExecutionError {
	cmd := exec.Command(command[0], command[1:]...)
	if len(pipeStdout) > 0 && pipeStdout[0] || len(pipeStdout) == 0 {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}
	err := cmd.Run()
	if err != nil {
		var exitErr *exec.ExitError
		if errors.As(err, &exitErr) {
			code := exitErr.ExitCode()
			return &ExecutionErrorImpl{
				Code: code,
				Err:  err,
				cmd:  command,
			}
		} else {
			return &ExecutionErrorImpl{
				Code: 255,
				Err:  err,
				cmd:  command,
			}
		}
	}
	return nil
}
