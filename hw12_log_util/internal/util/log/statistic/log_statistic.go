package statistic

import (
	"fmt"
	"strings"
	"time"

	"github.com/AndreiRubinJ/otusgo/hw12_log_util/internal/util/log/entry"
)

func ParseLogEntry(line string, filterLogType string) (*entry.BasicLogEntry, bool, error) {
	parts := strings.SplitN(line, " ", 3)
	if len(parts) < 3 {
		return nil, false, fmt.Errorf("invalid log format")
	}
	logType := parts[0]
	dateStr := parts[1]
	message := parts[2]

	if logType != filterLogType {
		return nil, false, nil
	}

	date, err := time.Parse(time.RFC3339, dateStr)
	if err != nil {
		return nil, false, fmt.Errorf("invalid date format: %w", err)
	}

	return &entry.BasicLogEntry{
		LogType: logType,
		Date:    date,
		Message: message,
	}, true, nil
}

func AddToStatistics(logs []entry.LogEntry) map[string]map[string]int {
	statistics := make(map[string]map[string]int)

	for _, log := range logs {
		logType := log.GetType()
		dateKey := log.GetDate().Format("2006-01-02")

		if _, ok := statistics[logType]; !ok {
			statistics[logType] = make(map[string]int)
		}

		statistics[logType][dateKey]++
	}

	return statistics
}
