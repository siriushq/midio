package pubsub

import (
	"fmt"
	"testing"
	"time"
)

func TestSubscribe(t *testing.T) {
	ps := New()
	ch1 := make(chan interface{}, 1)
	ch2 := make(chan interface{}, 1)
	doneCh := make(chan struct{})
	defer close(doneCh)
	ps.Subscribe(ch1, doneCh, nil)
	ps.Subscribe(ch2, doneCh, nil)
	ps.Lock()
	defer ps.Unlock()
	if len(ps.subs) != 2 {
		t.Errorf("expected 2 subscribers")
	}
}

func TestUnsubscribe(t *testing.T) {
	ps := New()
	ch1 := make(chan interface{}, 1)
	ch2 := make(chan interface{}, 1)
	doneCh1 := make(chan struct{})
	doneCh2 := make(chan struct{})
	ps.Subscribe(ch1, doneCh1, nil)
	ps.Subscribe(ch2, doneCh2, nil)

	close(doneCh1)
	// Allow for the above statement to take effect.
	time.Sleep(100 * time.Millisecond)
	ps.Lock()
	if len(ps.subs) != 1 {
		t.Errorf("expected 1 subscriber")
	}
	ps.Unlock()
	close(doneCh2)
}

func TestPubSub(t *testing.T) {
	ps := New()
	ch1 := make(chan interface{}, 1)
	doneCh1 := make(chan struct{})
	defer close(doneCh1)
	ps.Subscribe(ch1, doneCh1, func(entry interface{}) bool { return true })
	val := "hello"
	ps.Publish(val)
	msg := <-ch1
	if msg != "hello" {
		t.Errorf(fmt.Sprintf("expected %s , found %s", val, msg))
	}
}

func TestMultiPubSub(t *testing.T) {
	ps := New()
	ch1 := make(chan interface{}, 1)
	ch2 := make(chan interface{}, 1)
	doneCh := make(chan struct{})
	defer close(doneCh)
	ps.Subscribe(ch1, doneCh, func(entry interface{}) bool { return true })
	ps.Subscribe(ch2, doneCh, func(entry interface{}) bool { return true })
	val := "hello"
	ps.Publish(val)

	msg1 := <-ch1
	msg2 := <-ch2
	if msg1 != "hello" && msg2 != "hello" {
		t.Errorf(fmt.Sprintf("expected both subscribers to have%s , found %s and  %s", val, msg1, msg2))
	}
}
