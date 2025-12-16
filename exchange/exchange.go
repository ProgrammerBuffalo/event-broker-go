package exchange

import (
	"errors"

	"github.com/ProgrammerBuffalo/event-broker/event"
	"github.com/ProgrammerBuffalo/event-broker/queue"
)

type Exchange struct {
	name   string
	queues map[string]*queue.Queue
}

func NewExchange(name string) *Exchange {
	return &Exchange{
		name:   name,
		queues: make(map[string]*queue.Queue),
	}
}

func (exc *Exchange) AssignQueue(rKey string, queue *queue.Queue) {
	exc.queues[rKey] = queue
}

func (exc *Exchange) PublishEventToQueue(rKey string, event event.Event) error {
	q, exists := exc.queues[rKey]

	if !exists {
		return errors.New("routing key not found")
	}

	q.PublishEvent(event)
	return nil
}
