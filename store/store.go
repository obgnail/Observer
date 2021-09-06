package store

import "github.com/obgnail/Observer/event"

type EventPool interface {
	PutEvent(events ...*event.Event) error
	PullEvent() (*event.Event, error)
}
