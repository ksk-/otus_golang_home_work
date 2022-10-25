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

	os.Exit(RunCmd(os.Args[2:], env))
}
