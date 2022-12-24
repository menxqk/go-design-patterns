package main

import "fmt"

// The Observer design pattern (also called publisher/subscriber or publisher/listener) is
// used to uncouple an event from its possible handlers.
type Observer interface {
	Notify(string)
}

type Consumer struct {
	Name string
}

func (c *Consumer) Notify(s string) {
	fmt.Printf("Consumer %s received: '%s'\n", c.Name, s)
}

type Observable interface {
	Subscribe(Observer)
	Unsubscribe(Observer)
	Emit(string)
}

type Publisher struct {
	ObserversList []Observer
}

func (p *Publisher) Subscribe(o Observer) {
	for _, x := range p.ObserversList {
		if x == o {
			return
		}
	}
	p.ObserversList = append(p.ObserversList, o)
}

func (p *Publisher) Unsubscribe(o Observer) {
	for i, x := range p.ObserversList {
		if x == o {
			newSlice := p.ObserversList[:i]
			if i+1 < len(p.ObserversList) {
				newSlice = append(newSlice, p.ObserversList[i+1:]...)
			}
			p.ObserversList = newSlice
			break
		}
	}
}

func (p *Publisher) Emit(s string) {
	for _, o := range p.ObserversList {
		o.Notify(s)
	}
}

func main() {
	consumerA := Consumer{Name: "A"}
	consumerB := Consumer{Name: "B"}
	consumerC := Consumer{Name: "C"}

	publisher := Publisher{}
	publisher.Subscribe(&consumerA)
	publisher.Subscribe(&consumerB)
	publisher.Subscribe(&consumerC)

	publisher.Emit("First Message")
	fmt.Println()

	publisher.Unsubscribe(&consumerB)

	publisher.Emit("Second Message")
}
