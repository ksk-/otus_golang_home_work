package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRun_WithEnvironment(t *testing.T) {
	t.Run("simple case", func(t *testing.T) {
		res := runCmd(command{
			cmd:  "/bin/sh",
			args: []string{"-c", "echo $FOO $BAR"},
			env: Environment{
				"FOO": {Value: "foo", NeedRemove: false},
				"BAR": {Value: "bar", NeedRemove: false},
			},
		})
		require.Equal(t, 0, res.returnCode)
		require.Equal(t, "foo bar\n", res.stdout)

		t.Run("unset local vars after the running", func(t *testing.T) {
			for _, key := range []string{"FOO", "BAR"} {
				_, ok := os.LookupEnv(key)
				require.False(t, ok)
			}
		})
	})

	t.Run("replace and remove existent environment", func(t *testing.T) {
		setenv("FOO", "foo")
		setenv("BAR", "bar")
		setenv("BAZ", "baz")
		defer unsetenv("FOO", "BAR", "BAZ")

		res := runCmd(command{
			cmd:  "/bin/sh",
			args: []string{"-c", `bar="${BAR-removed}"; echo $FOO $bar $BAZ`},
			env: Environment{
				"FOO": {Value: "replaced", NeedRemove: false},
				"BAR": {NeedRemove: true},
			},
		})
		require.Equal(t, 0, res.returnCode)
		require.Equal(t, "replaced removed baz\n", res.stdout)

		t.Run("restore the original environment", func(t *testing.T) {
			v, ok := os.LookupEnv("FOO")
			require.True(t, ok)
			require.Equal(t, "foo", v)

			v, ok = os.LookupEnv("BAR")
			require.True(t, ok)
			require.Equal(t, "bar", v)

			v, ok = os.LookupEnv("BAZ")
			require.True(t, ok)
			require.Equal(t, "baz", v)
		})
	})
}

func TestRunCmd_KeepCmdBehaviour(t *testing.T) {
	t.Run("return codes forwarding", func(t *testing.T) {
		for _, code := range []int{0, 1, 2, 255} {
			res := runCmd(command{
				cmd:  "/bin/sh",
				args: []string{"-c", fmt.Sprintf("exit %d", code)},
			})
			require.Equal(t, code, res.returnCode)
		}
	})

	t.Run("standard files forwarding", func(t *testing.T) {
		t.Run("stdout and stderr", func(t *testing.T) {
			res := runCmd(command{
				cmd:  "/bin/sh",
				args: []string{"-c", "echo $FOO; echo $BAR 1>&2;"},
				env: Environment{
					"FOO": {Value: "foo", NeedRemove: false},
					"BAR": {Value: "bar", NeedRemove: false},
				},
			})
			require.Equal(t, 0, res.returnCode)
			require.Equal(t, "foo\n", res.stdout)
			require.Equal(t, "bar\n", res.stderr)
		})

		t.Run("stdin", func(t *testing.T) {
			res := runCmd(command{
				cmd:   "/bin/sh",
				args:  []string{"-c", "read input; echo $input"},
				stdin: "foo",
			})
			require.Equal(t, 0, res.returnCode)
			require.Equal(t, "foo\n", res.stdout)
		})
	})
}

type command struct {
	cmd   string
	args  []string
	env   Environment
	stdin string
}

type result struct {
	returnCode int
	stdout     string
	stderr     string
}

func runCmd(cmd command) result {
	res := result{}

	outPipe := newPipe()
	errPipe := newPipe()
	defer closeFiles(outPipe.reader, errPipe.reader)

	func() {
		inPipe := newPipe()
		defer closeFiles(inPipe.reader, outPipe.writer, errPipe.writer)

		origIn, origOut, origErr := os.Stdin, os.Stdout, os.Stderr
		defer func() {
			os.Stdin, os.Stdout, os.Stderr = origIn, origOut, origErr
		}()
		os.Stdin, os.Stdout, os.Stderr = inPipe.reader, outPipe.writer, errPipe.writer

		go func() {
			defer closeFiles(inPipe.writer)
			inPipe.write(cmd.stdin)
		}()

		res.returnCode = RunCmd(append([]string{cmd.cmd}, cmd.args...), cmd.env)
	}()
	res.stdout = outPipe.read()
	res.stderr = errPipe.read()

	return res
}

func closeFiles(files ...*os.File) {
	for _, f := range files {
		if err := f.Close(); err != nil {
			log.Fatalf("failed to close file: %v", err)
		}
	}
}

type pipe struct {
	reader *os.File
	writer *os.File
}

func (p *pipe) read() string {
	bytes, err := io.ReadAll(p.reader)
	if err != nil {
		log.Fatalf("failed to read from file: %v", err)
	}
	return string(bytes)
}

func (p *pipe) write(text string) {
	if _, err := p.writer.WriteString(text); err != nil {
		log.Fatalf("failed to write to file: %v", err)
	}
}

func newPipe() *pipe {
	r, w, err := os.Pipe()
	if err != nil {
		log.Fatalf("failed to create pipe: %v", err)
	}
	return &pipe{r, w}
}

func setenv(key, value string) {
	if err := os.Setenv(key, value); err != nil {
		log.Fatalf("failed to setenv: %v", err)
	}
}

func unsetenv(keys ...string) {
	for _, key := range keys {
		if err := os.Unsetenv(key); err != nil {
			log.Fatalf("failef to setenv: %v", err)
		}
	}
}
