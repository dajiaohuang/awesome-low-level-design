package pubsubsystem

import (
	"sync"
)

type Topic struct {
	Name        string
	Subscribers map[Subscriber]struct{}
	mu          sync.RWMutex
}

func NewTopic(name string) *Topic {
	return &Topic{
		Name:        name,
		Subscribers: make(map[Subscriber]struct{}),
	}
}

func (t *Topic) AddSubscriber(subscriber Subscriber) {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.Subscribers[subscriber] = struct{}{}
}

func (t *Topic) RemoveSubscriber(subscriber Subscriber) {
	t.mu.Lock()
	defer t.mu.Unlock()
	delete(t.Subscribers, subscriber)
}

func (t *Topic) Publish(message *Message) {
	t.mu.RLock()
	subscribers := make([]Subscriber, 0, len(t.Subscribers))
	for subscriber := range t.Subscribers {
		subscribers = append(subscribers, subscriber)
	}
	t.mu.RUnlock()
	for _, subscriber := range subscribers {
		subscriber.OnMessage(message)
	}
}
