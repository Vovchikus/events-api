package model

type Event struct {
	ID        string `json:"id"`
	EventType string `json:"type"`
	Content   string `json:"content"`
	CreatedAt string `json:"createdAt"`
}

func NewEvent(id string, content string, eventType string, createdAt string) *Event {
	return &Event{
		ID:        id,
		EventType: eventType,
		Content:   content,
		CreatedAt: createdAt,
	}
}
