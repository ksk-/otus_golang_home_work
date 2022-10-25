package main

import (
	"log"
	"os"
	"path"
	"testing"

	"github.com/stretchr/testify/suite"
)

const (
	inputFile     = "testdata/input.txt"
	inputFileSize = 6617
)

type CopyTestSuite struct {
	suite.Suite
	testDir    string
	outputFile string
}

func (s *CopyTestSuite) SetupTest() {
	dir, err := os.MkdirTemp(os.TempDir(), "hw07_file_copying")
	if err != nil {
		log.Fatalf("failed to create directory: %v", err)
	}

	f, err := os.CreateTemp(dir, "output")
	if err != nil {
		log.Fatalf("failed to create file: %v", err)
	}
	defer func(f *os.File) {
		if err := f.Close(); err != nil {
			log.Fatalf("failed to close file: %v", err)
		}
	}(f)

	s.testDir = dir
	s.outputFile = f.Name()
}

func (s *CopyTestSuite) TearDownTest() {
	if err := os.RemoveAll(s.testDir); err != nil {
		log.Fatalf("failed to remove directory: %v", err)
	}
}

func (s *CopyTestSuite) TestCopy() {
	s.Run("whole file", func() {
		err := Copy(inputFile, s.outputFile, 0, 0)
		s.Require().NoError(err)
		s.Require().Equal(readFile(inputFile), s.readOutputFile())
	})

	s.Run("some parts", func() {
		tests := []struct {
			offset   int64
			limit    int64
			expected string
		}{
			{offset: 0, limit: 2, expected: "Go"},
			{offset: 0, limit: 10, expected: "Go\nDocumen"},
			{offset: 57, limit: 15, expected: "Getting Started"},
			{offset: 6597, limit: inputFileSize, expected: "Supported by Google\n"},
		}

		for _, test := range tests {
			s.Run("", func() {
				err := Copy(inputFile, s.outputFile, test.offset, test.limit)
				s.Require().NoError(err)
				s.Require().Equal(test.expected, s.readOutputFile())
			})
		}
	})

	s.Run("offset exceeds file size", func() {
		err := Copy(inputFile, s.outputFile, 6597, inputFileSize*2)
		s.Require().NoError(err)
		s.Require().Equal("Supported by Google\n", s.readOutputFile())
	})
}

func (s *CopyTestSuite) TestInvalidInput() {
	s.Run("offset exceeds file size", func() {
		err := Copy(inputFile, s.outputFile, inputFileSize+1, 0)
		s.Require().ErrorIs(ErrOffsetExceedsFileSize, err)
	})

	s.Run("unsupported file", func() {
		err := Copy("/dev/urandom", s.outputFile, 0, 0)
		s.Require().ErrorIs(ErrUnsupportedFile, err)
	})

	s.Run("not existent file", func() {
		err := Copy("not_existent_file", s.outputFile, 0, 0)
		s.Require().ErrorContains(err, "no such file or directory")
	})

	s.Run("file without read permissions", func() {
		err := Copy(s.createWriteOnlyFile("write_only_file"), s.outputFile, 0, 0)
		s.Require().ErrorContains(err, "permission denied")
	})
}

func (s *CopyTestSuite) readOutputFile() string {
	return readFile(s.outputFile)
}

func (s *CopyTestSuite) createWriteOnlyFile(name string) string {
	f, err := os.OpenFile(path.Join(s.testDir, name), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.ModeAppend)
	if err != nil {
		log.Fatalf("failed to create file: %v", err)
	}
	defer func(f *os.File) {
		if err := f.Close(); err != nil {
			log.Fatalf("failed to close file: %v", err)
		}
	}(f)
	return f.Name()
}

func readFile(path string) string {
	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("failed to read file: %v", err)
	}
	return string(data)
}

func TestCopy(t *testing.T) {
	suite.Run(t, new(CopyTestSuite))
}
