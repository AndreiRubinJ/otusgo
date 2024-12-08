package statistic

import (
	"testing"
	"time"

	"github.com/AndreiRubinJ/otusgo/hw12_log_util/internal/util/log/entry"
	"github.com/stretchr/testify/assert"
)

func TestParseLogEntry(t *testing.T) {
	t.Run("ValidLogEntries", testValidLogEntries)
	t.Run("InvalidLogEntries", testInvalidLogEntries)
}

func testValidLogEntries(t *testing.T) {
	tests := []struct {
		name          string
		line          string
		filterLogType string
		expected      *entry.BasicLogEntry
		shouldMatch   bool
		expectError   bool
	}{
		{
			name:          "ValidLogEntry",
			line:          "INFO 2023-10-10T10:00:00Z This is a valid log entry",
			filterLogType: "INFO",
			expected: &entry.BasicLogEntry{
				LogType: "INFO",
				Date:    time.Date(2023, 10, 10, 10, 0, 0, 0, time.UTC),
				Message: "This is a valid log entry",
			},
			shouldMatch: true,
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			entry, matched, err := ParseLogEntry(tt.line, tt.filterLogType)
			if tt.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			if tt.shouldMatch {
				assert.Equal(t, tt.expected, entry)
			} else {
				assert.Nil(t, entry)
			}
			assert.Equal(t, tt.shouldMatch, matched)
		})
	}
}

func testInvalidLogEntries(t *testing.T) {
	tests := []struct {
		name          string
		line          string
		filterLogType string
		expected      *entry.BasicLogEntry
		shouldMatch   bool
		expectError   bool
	}{
		{
			name:          "InvalidLogFormat",
			line:          "ERROR 2023-10-10T10:00:00Z",
			filterLogType: "ERROR",
			expected:      nil,
			shouldMatch:   false,
			expectError:   true,
		},
		{
			name:          "InvalidDateFormat",
			line:          "INFO invalid-date This is a valid log entry",
			filterLogType: "INFO",
			expected:      nil,
			shouldMatch:   false,
			expectError:   true,
		},
		{
			name:          "MismatchedLogType",
			line:          "WARNING 2023-10-10T10:00:00Z This is a warning",
			filterLogType: "INFO",
			expected:      nil,
			shouldMatch:   false,
			expectError:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			entry, matched, err := ParseLogEntry(tt.line, tt.filterLogType)
			if tt.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			if tt.shouldMatch {
				assert.Equal(t, tt.expected, entry)
			} else {
				assert.Nil(t, entry)
			}
			assert.Equal(t, tt.shouldMatch, matched)
		})
	}
}
