package listplus

import (
	"math/rand/v2"
	"testing"
)

func checkListIsSorted(list *List[int]) bool {
	i := 0
	for front := list.Front(); front != nil; front = front.Next() {
		if front.Next() != nil && front.Value < front.Next().Value {
			return false
		}
		i++
	}

	if i != list.Len() {
		return false
	}

	for back := list.Back(); back != nil; back = back.Prev() {
		i--
	}
	if i != 0 {
		return false
	}

	return true
}

func TestList_Sort(t *testing.T) {
	t.Run("empty list", func(t *testing.T) {
		list := New[int]()
		list.Sort(func(front, back int) bool {
			return front > back
		})
		if !checkListIsSorted(list) {
			t.Fatal("list is not sorted")
		}
	})

	t.Run("no duplicate data", func(t *testing.T) {
		list := New[int]()
		for _, v := range rand.Perm(100) {
			list.PushBack(v)
		}
		list.Sort(func(front, back int) bool {
			return front > back
		})
		if !checkListIsSorted(list) {
			t.Fatal("list is not sorted")
		}
	})

	t.Run("duplicate data", func(t *testing.T) {
		list := New[int]()
		for _, v := range rand.Perm(301) {
			list.PushBack(v)
		}
		for i := 0; i < 100; i++ {
			list.PushBack(i)
		}
		for i := 200; i < 300; i++ {
			list.PushBack(i)
		}
		list.Sort(func(front, back int) bool {
			return front > back
		})
		if !checkListIsSorted(list) {
			t.Fatal("list is not sorted")
		}
	})

	t.Run("sorted list", func(t *testing.T) {
		list := New[int]()
		for i := 0; i < 100; i++ {
			list.PushBack(i)
		}
		list.Sort(func(front, back int) bool {
			return front > back
		})
		if !checkListIsSorted(list) {
			t.Fatal("list is not sorted")
		}
	})

	t.Run("sorted list with duplicate data", func(t *testing.T) {
		list := New[int]()
		for i := 0; i < 100; i++ {
			list.PushBack(i)
		}
		for i := 0; i < 100; i++ {
			list.PushBack(i)
		}
		list.Sort(func(front, back int) bool {
			return front > back
		})
		if !checkListIsSorted(list) {
			t.Fatal("list is not sorted")
		}
	})

	t.Run("reverse data", func(t *testing.T) {
		list := New[int]()
		for i := 100; i > 0; i-- {
			list.PushBack(i)
		}
		list.Sort(func(front, back int) bool {
			return front > back
		})
		if !checkListIsSorted(list) {
			t.Fatal("list is not sorted")
		}
	})
}

func TestRandLenthList_Sort(t *testing.T) {
	for i := 0; i < 100; i++ {
		list := New[int]()
		for _, v := range rand.Perm(rand.IntN(500)) {
			list.PushBack(v)
		}
		list.Sort(func(front, back int) bool {
			return front > back
		})
		if !checkListIsSorted(list) {
			t.Fatal("list is not sorted")
		}
	}
}

func BenchmarkList_Sort(b *testing.B) {
	list := New[int]()
	for i := 0; i < b.N; i++ {
		list.Init()
		for _, v := range rand.Perm(1000) {
			list.PushBack(v)
		}

		list.Sort(func(front, back int) bool {
			return front > back
		})
	}
}

func BenchmarkList_SortSored(b *testing.B) {
	list := New[int]()
	for _, v := range rand.Perm(1000) {
		list.PushBack(v)
	}
	list.Sort(func(front, back int) bool {
		return front > back
	})
	for i := 0; i < b.N; i++ {
		list.Sort(func(front, back int) bool {
			return front > back
		})
	}
}
