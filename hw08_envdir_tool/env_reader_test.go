package main

import (
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

const testDir = "testdata/env"

func TestReadDir(t *testing.T) {
	t.Run("regular", func(t *testing.T) {
		expected := Environment{
			"BAR":   {Value: "bar", NeedRemove: false},
			"FOO":   {Value: "   foo\nwith new line", NeedRemove: false},
			"EMPTY": {Value: "", NeedRemove: false},
			"HELLO": {Value: `"hello"`, NeedRemove: false},
			"UNSET": {Value: "", NeedRemove: true},
		}

		env, err := ReadDir(testDir)
		require.NoError(t, err)
		require.Equal(t, expected, env)
	})

	t.Run("skip files with invalid names", func(t *testing.T) {
		dir := makeInvalidEnvDir()
		env, err := ReadDir(dir)
		defer func(path string) {
			if err := os.RemoveAll(path); err != nil {
				log.Fatalf("failed to remove directory: %v", err)
			}
		}(dir)

		require.Empty(t, env)
		require.NoError(t, err)
	})
}

func makeInvalidEnvDir() string {
	dir, err := os.MkdirTemp(os.TempDir(), "invalid_env")
	if err != nil {
		log.Fatalf("failed to create directory: %v", err)
	}

	f, err := os.CreateTemp(dir, "invalid_name=")
	if err != nil {
		log.Fatalf("failed to create file: %v", err)
	}
	defer func(f *os.File) {
		if err := f.Close(); err != nil {
			log.Fatalf("failed to close file: %v", err)
		}
	}(f)

	return dir
}
