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
	firstItem *ListItem
	lastItem  *ListItem
	len       int
}

func (l list) Len() int {
	return l.len
}

func (l list) Front() *ListItem {
	return l.firstItem
}

func (l list) Back() *ListItem {
	return l.lastItem
}

func (l *list) PushFront(v interface{}) *ListItem {
	newListItem := new(ListItem)
	newListItem.Value = v

	if l.firstItem == nil {
		l.lastItem = newListItem
	} else {
		newListItem.Next = l.firstItem
		l.firstItem.Prev = newListItem
	}

	l.firstItem = newListItem
	l.len++

	return newListItem
}

func (l *list) PushBack(v interface{}) *ListItem {
	newListItem := new(ListItem)
	newListItem.Value = v

	if l.lastItem == nil {
		l.firstItem = newListItem
	} else {
		newListItem.Prev = l.lastItem
		l.lastItem.Next = newListItem
	}

	l.lastItem = newListItem
	l.len++

	return newListItem
}

func (l *list) Remove(i *ListItem) {
	if i.Next == nil {
		l.lastItem = i.Prev
	} else {
		i.Next.Prev = i.Prev
	}

	if i.Prev == nil {
		l.firstItem = i.Next
	} else {
		i.Prev.Next = i.Next
	}

	l.len--
}

func (l *list) MoveToFront(i *ListItem) {
	l.Remove(i)
	l.PushFront(i.Value)
}

func NewList() List {
	return new(list)
}
