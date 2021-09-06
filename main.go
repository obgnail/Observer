package main

import (
	"time"

	"github.com/obgnail/Observer/engine"
	"github.com/obgnail/Observer/event"
	"github.com/obgnail/Observer/observer"
)

func main() {
	//eng := engine.Default()  // no delay
	eng := engine.NewTickerEngine(5 * time.Second) // events in five seconds will be delayed

	eng.AddObserver(&observer.TypeObserver{}) // Custom observer
	eng.AddObserver(&observer.OperationObserver{})

	ev0 := event.Default()
	ev1 := event.NewDeleteEvent("myType1", nil) // Custom event
	ev2 := event.NewUpdateEvent("myType2", nil)
	ev3 := event.NewEvent("myType3", "add", nil)
	eng.PutEvent(ev0, ev1, ev2)         // no delay
	eng.PutDelayEvent(time.Second, ev3) // ev3 will push queue after one second delay

	// dummy sleep
	time.Sleep(10 * time.Second)
}
