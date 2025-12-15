package exchange

import (
	"github.com/ProgrammerBuffalo/event-driven/queue"
)

type Exchange struct {
	name   string
	queues map[string]queue.Queue
}

func NewExchange(name string) *Exchange {
	return &Exchange{queues: make(map[string]queue.Queue)}
}
