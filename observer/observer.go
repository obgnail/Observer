package observer

import "github.com/obgnail/Observer/event"

type Observer interface {
	Emit(e *event.Event) error
}
