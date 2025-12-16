package queue

import (
	"github.com/ProgrammerBuffalo/event-driven/event"
)

type Queue struct {
	name      string
	events    chan event.Event
	consumers chan chan event.Event
}

func NewQueue(name string, qSize int) *Queue {
	q := &Queue{
		name:      name,
		events:    make(chan event.Event, qSize),
		consumers: make(chan chan event.Event, 5),
	}

	go q.dispatch()
	return q
}

func (q *Queue) PublishEvent(event event.Event) {
	q.events <- event
}

func (q *Queue) AssignConsumer(consumer chan event.Event) chan chan event.Event {
	q.consumers <- consumer
	return q.consumers
}

func (q *Queue) dispatch() {
	for event := range q.events {
		consumer := <-q.consumers
		consumer <- event
	}
}
