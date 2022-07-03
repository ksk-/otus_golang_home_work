package hw04lrucache

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestList(t *testing.T) {
	t.Run("empty list", func(t *testing.T) {
		l := NewList()

		require.Equal(t, 0, l.Len())
		require.Nil(t, l.Front())
		require.Nil(t, l.Back())
	})

	t.Run("complex", func(t *testing.T) {
		l := NewList()

		l.PushFront(10) // [10]
		l.PushBack(20)  // [10, 20]
		l.PushBack(30)  // [10, 20, 30]
		require.Equal(t, 3, l.Len())

		middle := l.Front().Next // 20
		l.Remove(middle)         // [10, 30]
		require.Equal(t, 2, l.Len())

		for i, v := range [...]int{40, 50, 60, 70, 80} {
			if i%2 == 0 {
				l.PushFront(v)
			} else {
				l.PushBack(v)
			}
		} // [80, 60, 40, 10, 30, 50, 70]

		require.Equal(t, 7, l.Len())
		require.Equal(t, 80, l.Front().Value)
		require.Equal(t, 70, l.Back().Value)

		l.MoveToFront(l.Front()) // [80, 60, 40, 10, 30, 50, 70]
		l.MoveToFront(l.Back())  // [70, 80, 60, 40, 10, 30, 50]

		elems := make([]int, 0, l.Len())
		for i := l.Front(); i != nil; i = i.Next {
			elems = append(elems, i.Value.(int))
		}
		require.Equal(t, []int{70, 80, 60, 40, 10, 30, 50}, elems)
	})
}

func TestList_PushFront(t *testing.T) {
	l := NewList()

	t.Run("empty", func(t *testing.T) {
		i := l.PushFront(10) // [10]
		require.Equal(t, 1, l.Len())
		require.Equal(t, i, l.Front())
		require.Equal(t, i, l.Back())
	})

	t.Run("non empty", func(t *testing.T) {
		i := l.PushFront(20) // [20, 10]
		require.Equal(t, 2, l.Len())
		require.Equal(t, i, l.Front())
		require.Equal(t, i.Next, l.Back())

		i = l.PushFront(30) // [30, 20, 10]
		require.Equal(t, 3, l.Len())
		require.Equal(t, i, l.Front())
		require.Equal(t, i.Next.Next, l.Back())
	})
}

func TestList_PushBack(t *testing.T) {
	l := NewList()

	t.Run("empty", func(t *testing.T) {
		i := l.PushBack(10) // [10]
		require.Equal(t, 1, l.Len())
		require.Equal(t, i, l.Front())
		require.Equal(t, i, l.Back())
	})

	t.Run("non empty", func(t *testing.T) {
		i := l.PushBack(20) // [10, 20]
		require.Equal(t, 2, l.Len())
		require.Equal(t, i, l.Back())
		require.Equal(t, i.Prev, l.Front())

		i = l.PushBack(30) // [10, 20, 30]
		require.Equal(t, 3, l.Len())
		require.Equal(t, i, l.Back())
		require.Equal(t, i.Prev.Prev, l.Front())
	})
}

func TestList_Remove(t *testing.T) {
	t.Run("one item list", func(t *testing.T) {
		l := NewList()
		i := l.PushBack(10) // [10]

		l.Remove(i) // []
		require.Equal(t, 0, l.Len())
		require.Nil(t, l.Front())
		require.Nil(t, l.Back())
	})

	t.Run("two items list", func(t *testing.T) {
		t.Run("from front", func(t *testing.T) {
			l := NewList()
			l.PushBack(10) // [10]
			l.PushBack(20) // [10, 20]

			l.Remove(l.Front()) // [20]
			require.Equal(t, 1, l.Len())
			require.Equal(t, l.Front(), l.Back())
			require.Equal(t, 20, l.Front().Value)
		})

		t.Run("from back", func(t *testing.T) {
			l := NewList()
			l.PushBack(10) // [10]
			l.PushBack(20) // [10, 20]

			l.Remove(l.Back()) // [10]
			require.Equal(t, 1, l.Len())
			require.Equal(t, l.Front(), l.Back())
			require.Equal(t, 10, l.Front().Value)
		})
	})

	t.Run("three and more items list", func(t *testing.T) {
		t.Run("from front", func(t *testing.T) {
			l := NewList()
			front := l.PushBack(10)  // [10]
			middle := l.PushBack(20) // [10, 20]
			back := l.PushBack(30)   // [10, 20, 30]

			l.Remove(front) // [20, 30]
			require.Equal(t, 2, l.Len())
			require.Equal(t, middle, l.Front())
			require.Equal(t, back, l.Back())
		})

		t.Run("from middle", func(t *testing.T) {
			l := NewList()
			front := l.PushBack(10)  // [10]
			middle := l.PushBack(20) // [10, 20]
			back := l.PushBack(30)   // [10, 20, 30]

			l.Remove(middle) // [10, 30]
			require.Equal(t, 2, l.Len())
			require.Equal(t, front, l.Front())
			require.Equal(t, back, l.Back())
		})

		t.Run("from back", func(t *testing.T) {
			l := NewList()
			front := l.PushBack(10)  // [10]
			middle := l.PushBack(20) // [10, 20]
			back := l.PushBack(30)   // [10, 20, 30]

			l.Remove(back) // [10, 20]
			require.Equal(t, 2, l.Len())
			require.Equal(t, front, l.Front())
			require.Equal(t, middle, l.Back())
		})
	})
}

func TestList_MoveToFront(t *testing.T) {
	t.Run("one item list", func(t *testing.T) {
		l := NewList()
		i := l.PushBack(10) // [10]

		l.MoveToFront(i) // [10]
		require.Equal(t, 1, l.Len())
		require.Equal(t, l.Front(), l.Back())
	})

	t.Run("two and more items list", func(t *testing.T) {
		t.Run("from front", func(t *testing.T) {
			l := NewList()
			front := l.PushBack(10) // [10]
			back := l.PushBack(20)  // [10, 20]

			l.MoveToFront(l.Front()) // [10, 20]
			require.Equal(t, 2, l.Len())
			require.Equal(t, front, l.Front())
			require.Equal(t, back, l.Back())
		})

		t.Run("from middle", func(t *testing.T) {
			l := NewList()
			front := l.PushBack(10)  // [10]
			middle := l.PushBack(20) // [10, 20]
			back := l.PushBack(30)   // [10, 20, 30]

			l.MoveToFront(middle) // [20, 10, 30]
			require.Equal(t, 3, l.Len())
			require.Equal(t, middle, l.Front())
			require.Equal(t, front, l.Front().Next)
			require.Equal(t, back, l.Back())
		})

		t.Run("from back", func(t *testing.T) {
			l := NewList()
			front := l.PushBack(10)  // [10]
			middle := l.PushBack(20) // [10, 20]
			back := l.PushBack(30)   // [10, 20, 30]

			l.MoveToFront(l.Back()) // [30, 10, 20]
			require.Equal(t, 3, l.Len())
			require.Equal(t, back, l.Front())
			require.Equal(t, front, l.Front().Next)
			require.Equal(t, middle, l.Back())
		})
	})
}
