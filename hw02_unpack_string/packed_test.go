package hw02unpackstring

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestResetToNext(t *testing.T) {
	escaped := []rune(`\`)[0]

	tests := []struct {
		name     string
		input    [3]rune
		expected [3]rune
	}{
		{name: "ab", input: [3]rune{'a', 'b', null}, expected: [3]rune{'b', null, null}},
		{name: "a3", input: [3]rune{'a', '3', null}, expected: [3]rune{null, null, null}},
		{name: `\1a`, input: [3]rune{escaped, '1', 'a'}, expected: [3]rune{'a', null, null}},
		{name: `\13`, input: [3]rune{escaped, '1', '3'}, expected: [3]rune{null, null, null}},
		{name: `\\3`, input: [3]rune{escaped, escaped, '3'}, expected: [3]rune{null, null, null}},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			p := packed{runes: tc.input}
			p.resetToNext()
			require.Equal(t, tc.expected, p.runes)
		})
	}
}

func TestPushBack(t *testing.T) {
	escaped := []rune(`\`)[0]

	t.Run("First rune", func(t *testing.T) {
		tests := []struct {
			name  string
			input rune
			err   error
		}{
			{name: "regular", input: 'a', err: nil},
			{name: "digit", input: '3', err: ErrInvalidString},
			{name: "escaped", input: escaped, err: nil},
		}

		for _, tc := range tests {
			p := packed{}

			t.Run(tc.name, func(t *testing.T) {
				err := p.pushBack(tc.input)

				require.Equal(t, tc.err, err)
				require.False(t, p.isFinished())
				require.Equal(t, [3]rune{tc.input}, p.runes)
			})
		}
	})

	t.Run("Second rune", func(t *testing.T) {
		tests := []struct {
			name   string
			input  rune
			string string
		}{
			{name: "regular", input: 'a', string: "a"},
			{name: "regular", input: 'b', string: "a"},
			{name: "digit", input: '3', string: "aaa"},
			{name: "escaped", input: escaped, string: "a"},
		}

		for _, tc := range tests {
			p := packed{[3]rune{'a'}}

			t.Run(tc.name, func(t *testing.T) {
				err := p.pushBack(tc.input)

				require.NoError(t, err)
				require.True(t, p.isFinished())
				require.Equal(t, [3]rune{'a', tc.input}, p.runes)
				require.Equal(t, tc.string, p.string())
			})
		}
	})

	t.Run("Third rune", func(t *testing.T) {
		tests := []struct {
			name   string
			input  rune
			string string
		}{
			{name: "regular", input: 'a', string: "1"},
			{name: "regular", input: 'b', string: "1"},
			{name: "digit", input: '3', string: "111"},
			{name: "escaped", input: escaped, string: "1"},
		}

		for _, tc := range tests {
			p := packed{[3]rune{escaped, '1'}}

			t.Run(tc.name, func(t *testing.T) {
				err := p.pushBack(tc.input)

				require.NoError(t, err)
				require.True(t, p.isFinished())
				require.Equal(t, [3]rune{escaped, '1', tc.input}, p.runes)
				require.Equal(t, tc.string, p.string())
			})
		}
	})
}
