package queue

import "github.com/liyue201/gostl/ds/skiplist"

type slkey struct {
	key   string
	score uint64
}

func (s *slkey) Compare(cmp slkey) int {
	if s.score < cmp.score {
		return -1
	} else if s.score > cmp.score {
		return 1
	}
	return s.CompareKey(cmp)
}

func (s *slkey) CompareKey(cmp slkey) int {
	if s.key < cmp.key {
		return -1
	} else if s.key > cmp.key {
		return 1
	}
	return 0
}

func slKeyComparator(a, b slkey) int {
	return a.Compare(b)
}

type Comparator[T any] func(a, b T) int

type SkipList struct {
	list *skiplist.Skiplist[slkey, *Item]
}

func New() *SkipList {
	list := skiplist.New[slkey, *Item](slKeyComparator, skiplist.WithMaxLevel(15))
	return &SkipList{
		list,
	}
}

func (s *SkipList) Insert(key string, socre uint64, value *Item) {
	s.list.Insert(slkey{key, socre}, value)
}

func (s *SkipList) Delete(key slkey) bool {
	return s.list.Remove(key)
}

func (s *SkipList) Get(key slkey) *Item {
	v, e := s.list.Get(key)
	if e != nil {
		return nil
	}
	return v
}

func (s *SkipList) First() *Item {
	var value *Item
	s.list.Traversal(func(k slkey, v *Item) bool {
		value = v
		return false
	})
	return value
}
