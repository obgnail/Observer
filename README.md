## Use example

```go
// Custom observer0
type OperationObserver struct{}
// implement Emit() function
func (ob *OperationObserver) Emit(e *event.Event) (err error) {
	switch e.Action {
	case event.ActionAdd:
		fmt.Println("action observer receive add action", e)
	case event.ActionUpdate:
		fmt.Println("action observer receive update action", e)
	case event.ActionDelete:
		fmt.Println("action observer receive delete action", e)
	}
	return nil
}

// Custom observer1
type TypeObserver struct{}
func (ob *TypeObserver) Emit(e *event.Event) (err error) {
	fmt.Println("type observer receive event:", e)
	return nil
}

func main() {
	eng := engine.Default()  // no delay
	// eng := engine.NewTickerEngine(5 * time.Second) // events in five seconds will be delayed

	eng.AddObserver(&TypeObserver{})
	eng.AddObserver(&OperationObserver{})

	ev0 := event.Default()
	ev1 := event.NewDeleteEvent("myType1", nil)  // Custom event
	ev2 := event.NewEvent("myType3", "myAction", nil)
	eng.PutEvent(ev0, ev1)              // no delay
	eng.PutDelayEvent(time.Second, ev2) // ev3 will push queue after one second delay

	// dummy sleep
	time.Sleep(10 * time.Second)
}
```

