package model

type Event struct {
	ID      string `json:"id"`
	Content string `json:"content"`
}

func NewEvent(id string, content string) *Event {
	return &Event{
		ID:      id,
		Content: content,
	}
}

func (e *Event) GetId() string {
	return e.ID
}

func (e *Event) GetContent() string {
	return e.Content
}
