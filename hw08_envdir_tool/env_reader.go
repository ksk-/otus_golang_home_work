package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"strings"
)

type Environment map[string]EnvValue

// EnvValue helps to distinguish between empty files and files with the first empty line.
type EnvValue struct {
	Value      string
	NeedRemove bool
}

// ReadDir reads a specified directory and returns map of env variables.
// Variables represented as files where filename is name of variable, file first line is a value.
func ReadDir(dir string) (Environment, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	env := make(Environment)
	for _, entry := range entries {
		info, err := entry.Info()
		if err != nil {
			return nil, err
		}

		name := entry.Name()
		if strings.Contains(name, "=") {
			return nil, fmt.Errorf("invalid variable name: %s", name)
		}

		if info.Size() == 0 {
			env[name] = EnvValue{Value: "", NeedRemove: true}
		} else {
			v, err := readValue(path.Join(dir, name))
			if err != nil {
				return nil, err
			}
			env[name] = EnvValue{Value: v, NeedRemove: false}
		}
	}

	return env, nil
}

func readValue(name string) (string, error) {
	f, err := os.Open(name)
	if err != nil {
		return "", err
	}
	defer func(f *os.File) {
		if err := f.Close(); err != nil {
			log.Printf("failed to close file: %v", err)
		}
	}(f)

	r := bufio.NewReader(f)
	v, err := r.ReadBytes('\n')
	if err != nil && !errors.Is(err, io.EOF) {
		return "", err
	}
	v = bytes.ReplaceAll(v, []byte{0x00}, []byte{'\n'})

	return strings.TrimRight(string(v), " \t\n"), nil
}
