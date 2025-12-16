package main

import (
	"fmt"
	"time"

	"github.com/ProgrammerBuffalo/event-broker/event"
	"github.com/ProgrammerBuffalo/event-broker/eventbus"
)

func main() {
	// Initialize event bus
	eventBus := eventbus.NewEventBus()

	// Create queues in event bus
	eventBus.AddNewQueue("order-created-queue")
	eventBus.AddNewQueue("order-rejected-queue")

	// Create new exchange
	eventBus.AddNewExchange("order-exchange")

	// Assign queues to exchange
	eventBus.AssignQueueToExchange("order-exchange", "order-created-routing-key", "order-created-queue")
	eventBus.AssignQueueToExchange("order-exchange", "order-rejected-routing-key", "order-rejected-queue")

	// Initialize consumer #1 channel
	ch1 := make(chan event.Event)

	// Initialize consumer #2 creation channel
	ch2 := make(chan event.Event)

	// Initialize consumer #2 rejection channel
	ch3 := make(chan event.Event)

	ackCh1, err1 := eventBus.AssignConsumerToQueue("order-created-queue", ch1)

	if err1 != nil {
		fmt.Println(err1)
	}

	ackCh2, err2 := eventBus.AssignConsumerToQueue("order-created-queue", ch2)

	if err2 != nil {
		fmt.Println(err2)
	}

	ackCh3, err3 := eventBus.AssignConsumerToQueue("order-rejected-queue", ch3)

	if err3 != nil {
		fmt.Println(err3)
	}

	// Start 1 consumer
	go func() {
		for {
			fmt.Println("Consumer #1 listen event bus...")
			event := <-ch1
			fmt.Println("Consumer #1 start to consume", event.Payload)
			time.Sleep(1 * time.Second)
			fmt.Println("Consumer #1 end to consume", event.Payload)
			ackCh1 <- ch1
		}
	}()

	// Start 2 consumer
	go func() {
		for {
			fmt.Println("Consumer #2 listen event bus...")
			select {
			case createdEvent := <-ch2:
				fmt.Println("Consumer #2 start to consume", createdEvent.Payload)
				time.Sleep(500 * time.Millisecond)
				fmt.Println("Consumer #2 end to consume", createdEvent.Payload)
				ackCh2 <- ch2
			case rejectedEvent := <-ch3:
				fmt.Println("Consumer #2 start to consume", rejectedEvent.Payload)
				time.Sleep(500 * time.Millisecond)
				fmt.Println("Consumer #2 end to consume", rejectedEvent.Payload)
				ackCh3 <- ch3
			}
		}
	}()

	// Publish events to queues
	eventBus.PublishEvent("order-exchange", "order-created-routing-key", *event.NewEvent("Order #1 Created", time.Now()))
	eventBus.PublishEvent("order-exchange", "order-created-routing-key", *event.NewEvent("Order #2 Created", time.Now()))
	eventBus.PublishEvent("order-exchange", "order-rejected-routing-key", *event.NewEvent("Order #2 Rejected", time.Now()))
	eventBus.PublishEvent("order-exchange", "order-created-routing-key", *event.NewEvent("Order #3 Created", time.Now()))
	eventBus.PublishEvent("order-exchange", "order-created-routing-key", *event.NewEvent("Order #4 Created", time.Now()))

	time.Sleep(5 * time.Second)
}
