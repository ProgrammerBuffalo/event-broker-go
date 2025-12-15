package queue

import "github.com/ProgrammerBuffalo/event-driven/event"

type Queue struct {
	name      string
	messages  chan event.Event
	consumers chan chan event.Event
}

func NewQueue(name string, qSize int) *Queue {
	q := &Queue{
		name:      name,
		messages:  make(chan event.Event, qSize),
		consumers: make(chan chan event.Event),
	}

	go q.dispatch()
	return q
}

// Push to available consumer message from queue
func (q *Queue) dispatch() {
	for message := range q.messages {
		consumer := <-q.consumers
		consumer <- message
	}
}
