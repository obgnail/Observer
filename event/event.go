package event

const (
	ActionAdd    = "add"
	ActionUpdate = "update"
	ActionDelete = "delete"
)

type Event struct {
	Type   string                 `json:"type"`   // event的类型,可自定义
	Action string                 `json:"action"` // 默认提供Add Update Delete,可自定义
	Ctx    map[string]interface{} `json:"ctx"`    // 上下文,可存放需要传递的内容,可自定义
}

func Default() *Event {
	return NewEvent("", "", nil)
}

func NewEvent(Type string, action string, ctx map[string]interface{}) *Event {
	return &Event{Type: Type, Action: action, Ctx: ctx}
}

func (e *Event) SetAction(action string) *Event {
	e.Action = action
	return e
}

func (e *Event) SetType(typ string) *Event {
	e.Type = typ
	return e
}

func (e *Event) SetCtx(key string, value interface{}) *Event {
	e.Ctx[key] = value
	return e
}

func NewAddEvent(Type string, ctx map[string]interface{}) *Event {
	return NewEvent(Type, ActionAdd, ctx)
}

func NewUpdateEvent(Type string, ctx map[string]interface{}) *Event {
	return NewEvent(Type, ActionUpdate, ctx)
}

func NewDeleteEvent(Type string, ctx map[string]interface{}) *Event {
	return NewEvent(Type, ActionDelete, ctx)
}
