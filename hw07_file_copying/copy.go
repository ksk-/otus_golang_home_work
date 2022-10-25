package main

import (
	"errors"
	"io"
	"os"

	"github.com/cheggaaa/pb/v3"
)

const (
	defaultBufferSize = 4 * 1024
	pbTemplate        = `{{speed .}} {{ bar . "[" "#" "-" "-" "]"}} {{counters .}} {{percent .}}`
)

var (
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
)

func Copy(fromPath, toPath string, offset, limit int64) (err error) {
	in, size, err := openInputFileAt(fromPath, offset)
	if err != nil {
		return err
	}
	defer func(in io.ReadSeekCloser) {
		err = in.Close()
	}(in)

	if limit > 0 {
		size = min(limit, size)
	}

	bar := pb.ProgressBarTemplate(pbTemplate).Start64(size)
	defer bar.Finish()

	out, err := os.Create(toPath)
	if err != nil {
		return err
	}
	defer func(out *os.File) {
		err = out.Close()
	}(out)

	barWriter := bar.NewProxyWriter(out)
	bufferSize := min(defaultBufferSize, size)
	for totalWritten := int64(0); totalWritten < size; {
		written, err := io.CopyN(barWriter, in, bufferSize)
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return err
		}
		totalWritten += written
	}

	return err
}

func openInputFileAt(path string, offset int64) (io.ReadSeekCloser, int64, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, 0, err
	}

	stat, err := file.Stat()
	if err != nil {
		return nil, 0, err
	}

	if !stat.Mode().IsRegular() {
		return nil, 0, ErrUnsupportedFile
	}

	size := stat.Size()
	if offset > stat.Size() {
		return nil, 0, ErrOffsetExceedsFileSize
	}

	offset, err = file.Seek(offset, io.SeekStart)
	if err != nil {
		return nil, 0, err
	}

	return file, size - offset, nil
}

func min(x int64, y int64) int64 {
	if x < y {
		return x
	}

	return y
}
