package hw04lrucache

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

type list struct {
	size  int
	front *ListItem
	back  *ListItem
}

func (l *list) Len() int {
	return l.size
}

func (l *list) Front() *ListItem {
	return l.front
}

func (l *list) Back() *ListItem {
	return l.back
}

func (l *list) PushFront(v interface{}) *ListItem {
	i := &ListItem{
		Value: v,
		Next:  l.front,
		Prev:  nil,
	}

	if l.front != nil {
		l.front.Prev = i
	}

	l.size++
	l.front = i
	if l.back == nil {
		l.back = i
	}

	return i
}

func (l *list) PushBack(v interface{}) *ListItem {
	i := &ListItem{
		Value: v,
		Next:  nil,
		Prev:  l.back,
	}

	if l.back != nil {
		l.back.Next = i
	}

	l.size++
	l.back = i
	if l.front == nil {
		l.front = i
	}

	return i
}

func (l *list) Remove(i *ListItem) {
	if i.Prev != nil {
		i.Prev.Next = i.Next
	} else {
		l.front.Prev = nil
		l.front = l.front.Next
	}

	if i.Next != nil {
		i.Next.Prev = i.Prev
	} else {
		l.back.Next = nil
		l.back = l.back.Prev
	}

	l.size--
}

func (l *list) MoveToFront(i *ListItem) {
	if i != l.front {
		i.Prev.Next = i.Next

		if i.Next != nil {
			i.Next.Prev = i.Prev
		} else {
			l.back = l.back.Prev
			l.back.Next = nil
		}

		i.Prev = nil
		i.Next = l.front
		i.Next.Prev = i
		l.front = i
	}
}

func NewList() List {
	return &list{
		size:  0,
		front: nil,
		back:  nil,
	}
}
