package main

import (
	"os"
	"os/exec"
)

// RunCmd runs a command + arguments (cmd) with environment variables from env.
func RunCmd(cmd []string, env Environment) (returnCode int) {
	c := exec.Command(cmd[0], cmd[1:]...) //nolint:gosec
	c.Stdin = os.Stdin
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr

	for name, envValue := range env {
		if envValue.NeedRemove {
			os.Unsetenv(name)
			continue
		}

		os.Setenv(name, envValue.Value)
	}
	c.Env = os.Environ()

	if err := c.Run(); err != nil {
		return c.ProcessState.ExitCode()
	}

	return 0
}
