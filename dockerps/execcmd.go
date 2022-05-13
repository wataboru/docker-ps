package dockerps

import (
	"os/exec"
)

var (
	execCommand   = exec.Command
	commandOutput = func(c *exec.Cmd) ([]byte, error) {
		return c.Output()
	}
)

func _dockerPs() (string, error) {
	c := execCommand("docker", "ps", "-a")
	out, err := commandOutput(c)
	if err != nil {
		return "", err
	}
	return string(out), nil
}
