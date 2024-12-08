package log

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/AndreiRubinJ/otusgo/hw12_log_util/internal/util/log/entry"
	"github.com/AndreiRubinJ/otusgo/hw12_log_util/internal/util/log/statistic"
)

func GetStatistic(fileName string, filterLogType string) map[string]map[string]int {
	statistics := make(map[string]map[string]int)
	lines := readFile(fileName)
	if len(lines) > 0 {
		var logs []entry.LogEntry
		for _, line := range lines {
			if line == "" {
				continue
			}
			logEntry, matches, err := statistic.ParseLogEntry(line, filterLogType)
			if err != nil {
				fmt.Printf("Error parsing log entry: %v\n", err)
				continue
			}

			if matches {
				logs = append(logs, logEntry)
			}
		}
		statistics = statistic.AddToStatistics(logs)
	}

	return statistics
}

func readFile(fileName string) []string {
	var logs []string
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return logs
	}
	defer func() {
		if cerr := file.Close(); cerr != nil {
			fmt.Printf("Error closing file: %v\n", cerr)
		}
	}()
	content, err := io.ReadAll(file)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return logs
	}
	fmt.Printf("Read %d bytes\n", len(content))
	logs = strings.Split(string(content), "\n")
	return logs
}
