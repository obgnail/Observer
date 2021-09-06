package observer

import (
	"fmt"
	"github.com/obgnail/Observer/event"
)

type TypeObserver struct{}

func (ob *TypeObserver) Emit(e *event.Event) (err error) {
	fmt.Println("type observer receive event:", e)
	return nil
}
