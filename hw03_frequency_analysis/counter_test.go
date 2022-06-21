package hw03frequencyanalysis

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestInsert(t *testing.T) {
	value := "value"
	c := NewCounter[string]()
	for i := 1; i < 10; i++ {
		c.Insert(value)
		require.Equal(t, i, c.data[value])
	}
}

func TestTop(t *testing.T) {
	people := []string{"Alice", "Bob", "Dave", "Alice", "Dave", "Bob", "Carlos", "Bob"}

	c := NewCounter[string]()
	for _, person := range people {
		c.Insert(person)
	}

	require.Equal(t, []string{"Bob"}, c.Top(1))
	require.Equal(t, []string{"Bob", "Alice"}, c.Top(2))
	require.Equal(t, []string{"Bob", "Alice", "Dave"}, c.Top(3))
	require.Equal(t, []string{"Bob", "Alice", "Dave", "Carlos"}, c.Top(4))
	require.Equal(t, []string{"Bob", "Alice", "Dave", "Carlos"}, c.Top(5))
}
