package eventbus

import (
	"errors"

	"github.com/ProgrammerBuffalo/event-broker/event"
	"github.com/ProgrammerBuffalo/event-broker/exchange"
	"github.com/ProgrammerBuffalo/event-broker/queue"
)

type EventBus struct {
	exchanges map[string]*exchange.Exchange
	queues    map[string]*queue.Queue
}

func NewEventBus() *EventBus {
	return &EventBus{
		exchanges: make(map[string]*exchange.Exchange),
		queues:    make(map[string]*queue.Queue, 10),
	}
}

func (bus *EventBus) AddNewExchange(excName string) error {
	if _, exists := bus.exchanges[excName]; exists {
		return errors.New("exchange already exists")
	}

	bus.exchanges[excName] = exchange.NewExchange(excName)
	return nil
}

func (bus *EventBus) AddNewQueue(qName string) error {
	if _, exists := bus.queues[qName]; exists {
		return errors.New("queue already exists")
	}

	bus.queues[qName] = queue.NewQueue(qName, 10)
	return nil
}

func (bus *EventBus) AssignQueueToExchange(excName string, rKey string, qName string) error {
	exc, err := bus.findExchange(excName)

	if err != nil {
		return err
	}

	q, err := bus.findQueue(qName)

	if err != nil {
		return err
	}

	exc.AssignQueue(rKey, q)
	return nil
}

func (bus *EventBus) AssignConsumerToQueue(qName string, consumer chan event.Event) (chan chan event.Event, error) {
	q, err := bus.findQueue(qName)

	if err != nil {
		return nil, err
	}

	ch := q.AssignConsumer(consumer)
	return ch, nil
}

func (bus *EventBus) PublishEvent(excName string, rKey string, event event.Event) error {
	exc, err := bus.findExchange(excName)

	if err != nil {
		return err
	}

	err = exc.PublishEventToQueue(rKey, event)

	if err != nil {
		return err
	}

	return nil
}

func (bus *EventBus) findExchange(excName string) (*exchange.Exchange, error) {
	exc, exists := bus.exchanges[excName]

	if !exists {
		return nil, errors.New("exchange doesn't exist")
	}

	return exc, nil
}

func (bus *EventBus) findQueue(qName string) (*queue.Queue, error) {
	q, exists := bus.queues[qName]

	if !exists {
		return nil, errors.New("queue doesn't exist")
	}

	return q, nil
}
