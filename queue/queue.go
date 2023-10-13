package queue

type Item struct {
	Key      string
	Data     []byte
	SendTime uint64
}

type Queue interface {
	Pop() (item Item, err error)
	Insert(item Item) error
	Del(key string) error
}
