package main

import (
	"context"
	"fmt"
	"time"
)

var events = []*Event{
	{ID: "1", Type: "click", Value: 9},
	{ID: "2", Type: "view", Value: 16},
	{ID: "3", Type: "click", Value: 1},
	{ID: "4", Type: "view", Value: 4},
}

type Event struct {
	ID        string    `json:"id"`
	Type      string    `json:"type"`  // Тип события
	Value     int       `json:"value"` // Значение, которое мы будем агрегировать
	Timestamp time.Time `json:"timestamp"`
}

func ReciverEvent(ctx context.Context, eventCh <-chan *Event, store map[string]int) {
	for {
		select {
		case <-ctx.Done():
			return

		case event := <-eventCh:

			if event != nil {
				store[event.Type] += event.Value
			} else {
				return
			}

		default:
			continue
		}
	}
}

func ProduceEvent(eventCh chan<- *Event, events []*Event) {
	for _, event := range events {
		eventCh <- event
	}

	close(eventCh)
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// var wg sync.WaitGroup

	eventCh := make(chan *Event)
	store := make(map[string]int, 10)

	go func() {
		ReciverEvent(ctx, eventCh, store)
	}()

	ProduceEvent(eventCh, events)

	// Print
	fmt.Println(store["click"], store["view"])

}
