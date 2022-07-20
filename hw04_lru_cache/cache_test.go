package hw04lrucache

import (
	"math/rand"
	"strconv"
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCache(t *testing.T) {
	t.Run("empty cache", func(t *testing.T) {
		c := NewCache(10)

		_, ok := c.Get("aaa")
		require.False(t, ok)

		_, ok = c.Get("bbb")
		require.False(t, ok)
	})

	t.Run("simple", func(t *testing.T) {
		c := NewCache(5)

		wasInCache := c.Set("aaa", 100)
		require.False(t, wasInCache)

		wasInCache = c.Set("bbb", 200)
		require.False(t, wasInCache)

		val, ok := c.Get("aaa")
		require.True(t, ok)
		require.Equal(t, 100, val)

		val, ok = c.Get("bbb")
		require.True(t, ok)
		require.Equal(t, 200, val)

		wasInCache = c.Set("aaa", 300)
		require.True(t, wasInCache)

		val, ok = c.Get("aaa")
		require.True(t, ok)
		require.Equal(t, 300, val)

		val, ok = c.Get("ccc")
		require.False(t, ok)
		require.Nil(t, val)
	})

	t.Run("purge logic", func(t *testing.T) {
		t.Run("by filling cache", func(t *testing.T) {
			c := NewCache(3)
			c.Set("one", 1)   // ["one":1]
			c.Set("two", 2)   // ["two":2, "one":1]
			c.Set("three", 3) // ["three":3, "two":2, "one":1]
			c.Set("four", 4)  // ["four":4, "three":3, "two":2]
			c.Set("five", 5)  // ["five":5, "four":4, "three":3]

			for _, key := range []Key{"one", "two"} {
				value, ok := c.Get(key)
				require.False(t, ok)
				require.Nil(t, value)
			}

			value, ok := c.Get("three")
			require.True(t, ok)
			require.Equal(t, 3, value)

			value, ok = c.Get("four")
			require.True(t, ok)
			require.Equal(t, 4, value)

			value, ok = c.Get("five")
			require.True(t, ok)
			require.Equal(t, 5, value)
		})

		t.Run("for long unused items", func(t *testing.T) {
			c := NewCache(3)
			c.Set("one", 1)   // ["one":1]
			c.Set("two", 2)   // ["two":2, "one":1]
			c.Set("three", 3) // ["three":3, "two":2, "one":1]

			c.Set("one", -1) // ["one":-1, "three":3, "two":2]
			c.Set("four", 4) // ["four":4, "one":-1, "three":3]

			value, ok := c.Get("two")
			require.False(t, ok)
			require.Nil(t, value)

			c.Get("one")   // ["one":-1, "four":4, "three":3]
			c.Set("*", 0)  // ["*":0, "one":-1, "four":4]
			c.Set("**", 0) // ["**":0, "*":0, "one":-1]

			value, ok = c.Get("one") // ["one":-1, "**":0, "*":0]
			require.True(t, ok)
			require.Equal(t, -1, value)
		})
	})
}

func TestCacheMultithreading(t *testing.T) {
	c := NewCache(10)
	wg := &sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 0; i < 1_000_000; i++ {
			c.Set(Key(strconv.Itoa(i)), i)
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 1_000_000; i++ {
			c.Get(Key(strconv.Itoa(rand.Intn(1_000_000))))
		}
	}()

	wg.Wait()
}
