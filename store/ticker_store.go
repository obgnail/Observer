package store

import (
	"github.com/obgnail/Observer/event"
	"time"
)

const tickerChanSize = 1024

var _ EventPool = (*TickerEventPool)(nil)

// 将固定时间内的信息汇总
type TickerEventPool struct {
	size       int
	tickerTime time.Duration
	workFunc   func(event *event.Event)

	inputChan  chan *event.Event
	outputChan chan *event.Event
	tickerChan chan struct{}
}

func NewTickerEventPool(tickerTime time.Duration, work func(event *event.Event)) *TickerEventPool {
	tc := &TickerEventPool{
		inputChan:  make(chan *event.Event, tickerChanSize),
		outputChan: make(chan *event.Event, 1),
		tickerChan: make(chan struct{}, 1),
		tickerTime: tickerTime,
		workFunc:   work,
	}

	go tc.output()
	go tc.ticker()
	go tc.run()
	return tc
}

func (tc *TickerEventPool) PutEvent(events ...*event.Event) error {
	for _, e := range events {
		tc.inputChan <- e
	}
	return nil
}

func (tc *TickerEventPool) PullEvent() (*event.Event, error) {
	return <-tc.outputChan, nil
}

func (tc *TickerEventPool) ticker() {
	ticker := time.NewTicker(tc.tickerTime)
	go func() {
		for range ticker.C {
			tc.tickerChan <- struct{}{}
		}
	}()
}

func (tc *TickerEventPool) output() {
	var dataset []*event.Event
	for {
		select {
		case d := <-tc.inputChan:
			dataset = append(dataset, d)
		case <-tc.tickerChan:
			for _, d := range dataset {
				tc.outputChan <- d
			}
			dataset = []*event.Event{}
		}
	}
}

func (tc *TickerEventPool) run() {
	for {
		select {
		case ev := <-tc.outputChan:
			if ev == nil {
				continue
			}
			if tc.workFunc == nil {
				panic("store push func must init.")
			}
			tc.workFunc(ev)
		}
	}
}
