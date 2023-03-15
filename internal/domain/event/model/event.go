package model

type Event struct {
	id      string
	content string
}

func NewEvent(id string, content string) *Event {
	return &Event{
		id:      id,
		content: content,
	}
}

func (e *Event) Id() string {
	return e.id
}

func (e *Event) Content() string {
	return e.content
}
