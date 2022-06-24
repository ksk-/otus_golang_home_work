package hw02unpackstring

import (
	"errors"
	"strings"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(line string) (string, error) {
	if len(line) == 0 {
		return "", nil
	}

	curr := packed{}
	var sb strings.Builder

	for _, r := range line {
		err := curr.pushBack(r)
		if err != nil {
			return "", err
		}

		if curr.isFinished() {
			sb.WriteString(curr.string())
			curr.resetToNext()
		}
	}

	sb.WriteString(curr.string())

	return sb.String(), nil
}
