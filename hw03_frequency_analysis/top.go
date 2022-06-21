package hw03frequencyanalysis

import (
	"regexp"
	"strings"
	"unicode"
)

var re = regexp.MustCompile(`\S+`)

func Top10(text string) []string {
	c := NewCounter()
	for _, word := range re.FindAllString(text, -1) {
		if word != "-" {
			word = strings.TrimLeftFunc(word, unicode.IsPunct)
			word = strings.TrimRightFunc(word, unicode.IsPunct)
			word = strings.ToLower(word)
			c.Insert(word)
		}
	}

	return c.Top(10)
}
