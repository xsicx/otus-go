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

	t.Run("correct work with many types", func(t *testing.T) {
		l := NewList()

		l.PushFront([2]int{99, 1}) // [slice]
		l.PushFront("string")      // ["string", slice]
		l.PushFront(10)            // [10, "string", slice]
		l.PushBack(true)           // [10, "string", slice, true]

		require.Equal(t, 4, l.Len())
		require.Equal(t, 10, l.Front().Value)
		require.Equal(t, true, l.Back().Value)

		victim := l.Back().Prev // [2]int{99, 1}
		l.Remove(victim)        // [10, "string", true]
		require.Equal(t, 3, l.Len())
		l.MoveToFront(l.Back()) // [true, 10, "string"]

		require.Equal(t, true, l.Front().Value)
		require.Equal(t, "string", l.Back().Value)
	})

	t.Run("first item move to front", func(t *testing.T) {
		l := NewList()
		l.PushFront(10) // [10]
		l.PushBack(20)  // [10, 20]
		l.PushBack(30)  // [10, 20, 30]
		require.Equal(t, 10, l.Front().Value)
		l.MoveToFront(l.Front()) // [10, 20, 30]
		require.Equal(t, 10, l.Front().Value)
	})
}
