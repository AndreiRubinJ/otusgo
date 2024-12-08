package entry

import (
	"time"
)

type LogEntry interface {
	GetType() string
	GetDate() time.Time
	GetMessage() string
}

type BasicLogEntry struct {
	LogType string
	Date    time.Time
	Message string
}

func (entry *BasicLogEntry) GetType() string {
	return entry.LogType
}

func (entry *BasicLogEntry) GetDate() time.Time {
	return entry.Date
}

func (entry *BasicLogEntry) GetMessage() string {
	return entry.Message
}
