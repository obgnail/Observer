package store

import (
	"github.com/obgnail/Observer/event"
	"log"
)

const simpleChanSize = 1024

var _ EventPool = (*SimpleEventPool)(nil)

type SimpleEventPool struct {
	ch       chan *event.Event
	workFunc func(event *event.Event)
}

type workFunc func(event *event.Event)

func NewSimpleEventPool(work workFunc) *SimpleEventPool {
	ss := &SimpleEventPool{
		ch:       make(chan *event.Event, simpleChanSize),
		workFunc: work,
	}
	go ss.run()
	return ss
}

func (ss *SimpleEventPool) PutEvent(events ...*event.Event) error {
	for _, e := range events {
		ss.ch <- e
	}
	return nil
}

func (ss *SimpleEventPool) PullEvent() (*event.Event, error) {
	return <-ss.ch, nil
}

var observerDone = make(chan struct{})

func (ss *SimpleEventPool) Release() {
	close(observerDone)
}

func (ss *SimpleEventPool) run() {
	for {
		select {
		case <-observerDone:
			log.Println("event store PullEvent() closed.")
			return
		case ev := <-ss.ch:
			if ev == nil {
				continue
			}
			if ss.workFunc == nil {
				panic("store push func must init.")
			}
			ss.workFunc(ev)
		}
	}
}
