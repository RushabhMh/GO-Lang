package main

import (
	"fmt"
	"sync"
)

// VERY VERY IMP

type PubSub struct {
	mu      sync.RWMutex
	subbers map[string][]chan interface{}
}

func NewPubSub() *PubSub {
	return &PubSub{
		subbers: make(map[string][]chan interface{}),
	}
}

func (p *PubSub) Subscribe(topic string) <-chan interface{} {
	ch := make(chan interface{})
	p.mu.Lock()
	p.subbers[topic] = append(p.subbers[topic], ch)
	p.mu.Unlock()
	return ch
}

func (p *PubSub) Unsubscribe(topic string, ch <-chan interface{}) {
	p.mu.Lock()
	defer p.mu.Unlock()
	subs := p.subbers[topic]
	for i, sub := range subs {
		if sub == ch {
			p.subbers[topic] = append(subs[:i], subs[i+1:]...)
			close(sub)
			break
		}
	}
}

func (p *PubSub) Publish(topic string, msg interface{}) {

	p.mu.RLock()
	defer p.mu.RUnlock()
	for _, ch := range p.subbers[topic] {
		ch <- msg
	}

}

func main() {
	ps := NewPubSub()

	// Subscribe to the "example" topic
	sub := ps.Subscribe("example")

	// Start a goroutine to listen for messages
	go func() {
		for msg := range sub {
			fmt.Println("Received:", msg)
		}
	}()

	// Publish messages to the "example" topic
	ps.Publish("example", 12133121)
	ps.Publish("example", "Another message")

	// Unsubscribe from the "example" topic
	ps.Unsubscribe("example", sub)
}
