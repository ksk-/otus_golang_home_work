package hw03frequencyanalysis

import (
	"sort"
)

type Counter struct {
	data map[string]int
}

func (c *Counter) Insert(value string) {
	if count, ok := c.data[value]; ok {
		c.data[value] = count + 1
	} else {
		c.data[value] = 1
	}
}

func (c *Counter) Top(n int) []string {
	s := c.toSlice()
	sort.Slice(s, func(i, j int) bool {
		return s[i].value >= s[j].value
	})
	sort.SliceStable(s, func(i, j int) bool {
		return s[i].count >= s[j].count
	})

	count := min(n, len(s))
	res := make([]string, 0, count)
	for i := 0; i < count; i++ {
		res = append(res, s[i].value)
	}

	return res
}

func NewCounter() *Counter {
	return &Counter{data: make(map[string]int)}
}

type item struct {
	value string
	count int
}

func (c *Counter) toSlice() []item {
	s := make([]item, 0, len(c.data))
	for k, v := range c.data {
		s = append(s, item{value: k, count: v})
	}

	return s
}

func min(x int, y int) int {
	if x < y {
		return x
	}

	return y
}
