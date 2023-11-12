package queue

import (
	"sync"
)

type index struct {
	keymap map[string]uint64
	rwx    sync.RWMutex
}

var (
	defaultIndexShard          = 10
	defaultShardingSpan uint64 = 1000 * 60 * 60 * 24
)

var (
	indexes   = make([]index, defaultIndexShard)
	skipLists = make(map[uint64]string)
)

type shardingQueue struct {
}

func (s *shardingQueue) Pop() (item Item, err error) {

	return Item{}, nil
}

func (s *shardingQueue) Insert(item Item) error {
	return nil
}

func (s *shardingQueue) Del(key string) error {
	return nil
}
