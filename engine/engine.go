package engine

import (
	"log"
	"time"

	"github.com/ivpusic/grpool"
	"github.com/obgnail/Observer/event"
	"github.com/obgnail/Observer/observer"
	"github.com/obgnail/Observer/store"
)

const (
	WorkerMax   = 4
	JobQueueMax = 64
)

type EventEngine struct {
	// event的存储中心
	eventStore store.EventPool
	// event的打工人
	eventWorker *grpool.Pool
	// observers
	observers []observer.Observer
}

func Default() *EventEngine {
	engine := &EventEngine{}
	engine.eventWorker = grpool.NewPool(WorkerMax, JobQueueMax)
	engine.eventStore = store.NewSimpleEventPool(engine.processWork)
	return engine
}

func NewTickerEngine(tickerTime time.Duration) *EventEngine {
	engine := &EventEngine{}
	engine.eventWorker = grpool.NewPool(WorkerMax, JobQueueMax)
	engine.eventStore = store.NewTickerEventPool(tickerTime, engine.processWork)
	return engine
}

func (eng *EventEngine) AddObserver(ob observer.Observer) {
	eng.observers = append(eng.observers, ob)
}

func (eng *EventEngine) processWork(event *event.Event) {
	eng.eventWorker.JobQueue <- func() {
		eng.processEvent(event)
	}
}

func (eng *EventEngine) processEvent(e *event.Event) {
	for _, ob := range eng.observers {
		err := ob.Emit(e)
		if err != nil {
			log.Println("[Error] Observer update err:", err)
		}
	}
}

func (eng *EventEngine) Release() {
	if eng.eventWorker != nil {
		eng.eventWorker.Release()
	}
}

func (eng *EventEngine) PutEvent(e ...*event.Event) {
	err := eng.eventStore.PutEvent(e...)
	if err != nil {
		log.Println("put to store was err:", err)
	}
}

func (eng *EventEngine) PutDelayEvent(delay time.Duration, e ...*event.Event) {
	time.Sleep(delay)
	eng.PutEvent(e...)
}
