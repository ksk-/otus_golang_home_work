package hw02unpackstring

import (
	"strconv"
	"strings"
	"unicode"
)

const null = 0

type packed struct {
	runes [3]rune
}

func (p *packed) isEscaped() bool {
	return string(p.runes[0]) == `\`
}

func (p *packed) isValid() bool {
	if unicode.IsDigit(p.runes[0]) {
		return false
	}

	if p.isEscaped() {
		second := p.runes[1]
		return unicode.IsDigit(second) || string(second) == `\` || second == null
	}

	return true
}

func (p *packed) isFinished() bool {
	if p.isEscaped() {
		return p.runes[2] != null
	}

	return p.runes[1] != null
}

func (p *packed) string() string {
	if p.runes[0] == null {
		return ""
	}

	second := string(p.runes[1])

	if p.isEscaped() {
		if count, err := strconv.Atoi(string(p.runes[2])); err == nil {
			return strings.Repeat(second, count)
		}

		return second
	}

	first := string(p.runes[0])

	if count, err := strconv.Atoi(second); err == nil {
		return strings.Repeat(first, count)
	}

	return first
}

func (p *packed) resetToNext() {
	if p.isEscaped() {
		p.runes = [3]rune{p.runes[2], null, null}
	} else {
		p.runes = [3]rune{p.runes[1], null, null}
	}

	if unicode.IsDigit(p.runes[0]) {
		p.runes[0] = null
	}
}

func (p *packed) pushBack(r rune) error {
	switch {
	case p.runes[0] == null:
		p.runes[0] = r

	case p.runes[1] == null:
		p.runes[1] = r

	case p.runes[2] == null:
		if p.isEscaped() {
			p.runes[2] = r
		} else {
			p.runes[0] = p.runes[1]
			p.runes[1] = r
		}
	}

	if !p.isValid() {
		return ErrInvalidString
	}

	return nil
}
