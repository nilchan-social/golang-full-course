package scanner

import "time"

type Event struct {
	Description string
	UserInput   string
	DateAt      time.Time
}

func NewEvent(description string, userInput string) Event {
	return Event{
		Description: description,
		UserInput:   userInput,
		DateAt:      time.Now(),
	}
}
