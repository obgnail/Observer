package observer

import (
	"fmt"
	"github.com/obgnail/Observer/event"
)

type OperationObserver struct{}

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
