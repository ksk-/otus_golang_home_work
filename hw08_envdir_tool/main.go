package main

import (
	"log"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		log.Fatalf("usage error: not enough args to run app")
	}

	dir := os.Args[1]
	env, err := ReadDir(dir)
	if err != nil {
		log.Fatalf("failed to read env dir: %v", err)
	}

	cmd := os.Args[2]
	args := os.Args[3:]
	os.Exit(RunCmd(append([]string{cmd}, args...), env))
}
