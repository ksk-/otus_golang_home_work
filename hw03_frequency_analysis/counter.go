package hw03frequencyanalysis

import (
	"sort"

	"golang.org/x/exp/constraints"
)

type Counter[T constraints.Ordered] struct {
	data map[T]int
}

func (c *Counter[T]) Insert(value T) {
	if count, ok := c.data[value]; ok {
		c.data[value] = count + 1
	} else {
		c.data[value] = 1
	}
}

func (c *Counter[T]) Top(n int) []T {
	s := c.toSlice()
	sort.Slice(s, func(i, j int) bool {
		return s[i].value >= s[j].value
	})
	sort.SliceStable(s, func(i, j int) bool {
		return s[i].count >= s[j].count
	})

	count := min(n, len(s))
	res := make([]T, 0, count)
	for i := 0; i < count; i++ {
		res = append(res, s[i].value)
	}

	return res
}

func NewCounter[T constraints.Ordered]() *Counter[T] {
	return &Counter[T]{data: make(map[T]int)}
}

type item[T constraints.Ordered] struct {
	value T
	count int
}

func (c *Counter[T]) toSlice() []item[T] {
	s := make([]item[T], 0, len(c.data))
	for k, v := range c.data {
		s = append(s, item[T]{value: k, count: v})
	}

	return s
}

func min[T constraints.Ordered](x T, y T) T {
	if x < y {
		return x
	}

	return y
}
