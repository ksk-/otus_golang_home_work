package main

import (
	"errors"
	"log"
	"os"
	"os/exec"
	"strings"
)

// RunCmd runs a command + arguments (cmd) with environment variables from env.
func RunCmd(cmd []string, env Environment) (returnCode int) {
	if len(cmd) < 1 {
		log.Fatalf("failed to run command: the command isn't set")
	}

	name := cmd[0]
	args := cmd[1:]

	command := exec.Command(name, args...)
	command.Stdin = os.Stdin
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	command.Env = prepareCmdEnv(env)

	if err := command.Run(); err != nil {
		var exitError *exec.ExitError
		if errors.As(err, &exitError) {
			returnCode = exitError.ExitCode()
		} else {
			log.Fatalf("failed to run command: %v", err)
		}
	}

	return returnCode
}

func prepareCmdEnv(env Environment) []string {
	var toSet, toUnset []string
	for k, v := range env {
		if v.NeedRemove {
			toUnset = append(toUnset, k)
		} else {
			toSet = append(toSet, k+"="+v.Value)
		}
	}

	e := os.Environ()
	for i, v := range e {
		for _, key := range toUnset {
			if strings.Contains(v, key+"=") {
				e = append(e[:i], e[i+1:]...)
				break
			}
		}
	}

	return append(e, toSet...)
}
