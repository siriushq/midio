package pubsub

import (
	"sync"
	"sync/atomic"
)

// Sub - subscriber entity.
type Sub struct {
	ch     chan interface{}
	filter func(entry interface{}) bool
}

// PubSub holds publishers and subscribers
type PubSub struct {
	subs           []*Sub
	numSubscribers int32
	sync.RWMutex
}

// Publish message to the subscribers.
// Note that publish is always nob-blocking send so that we don't block on slow receivers.
// Hence receivers should use buffered channel so as not to miss the published events.
func (ps *PubSub) Publish(item interface{}) {
	ps.RLock()
	defer ps.RUnlock()

	for _, sub := range ps.subs {
		if sub.filter == nil || sub.filter(item) {
			select {
			case sub.ch <- item:
			default:
			}
		}
	}
}

// Subscribe - Adds a subscriber to pubsub system
func (ps *PubSub) Subscribe(subCh chan interface{}, doneCh <-chan struct{}, filter func(entry interface{}) bool) {
	ps.Lock()
	defer ps.Unlock()

	sub := &Sub{subCh, filter}
	ps.subs = append(ps.subs, sub)
	atomic.AddInt32(&ps.numSubscribers, 1)

	go func() {
		<-doneCh

		ps.Lock()
		defer ps.Unlock()

		for i, s := range ps.subs {
			if s == sub {
				ps.subs = append(ps.subs[:i], ps.subs[i+1:]...)
			}
		}
		atomic.AddInt32(&ps.numSubscribers, -1)
	}()
}

// NumSubscribers returns the number of current subscribers
func (ps *PubSub) NumSubscribers() int32 {
	return atomic.LoadInt32(&ps.numSubscribers)
}

// New inits a PubSub system
func New() *PubSub {
	return &PubSub{}
}
