package entry

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestBasicLogEntry_GetType(t *testing.T) {
	expectedType := "INFO"
	entry := &BasicLogEntry{LogType: expectedType}

	assert.Equal(t, expectedType, entry.GetType(), "LogEntry type should match")
}

func TestBasicLogEntry_GetDate(t *testing.T) {
	expectedDate := time.Now()
	entry := &BasicLogEntry{Date: expectedDate}

	assert.True(t, entry.GetDate().Equal(expectedDate), "LogEntry dates should be equal")
}

func TestBasicLogEntry_GetMessage(t *testing.T) {
	expectedMessage := "This is a log message."
	entry := &BasicLogEntry{Message: expectedMessage}

	assert.Equal(t, expectedMessage, entry.GetMessage(), "LogEntry message should match")
}
