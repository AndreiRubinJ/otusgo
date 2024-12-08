package log

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func OutputStatisticToFile(statistics map[string]map[string]int, fileName string) {
	if _, err := os.Stat(fileName); err == nil {
		err := os.Remove(fileName)
		if err != nil {
			fmt.Printf("Failed to remove existing file: %v", err)
			return
		}
	}

	file, err := os.Create(fileName)
	if err != nil {
		fmt.Printf("Failed to open file: %v", err)
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()

	header := []string{"Log Type", "Date", "Count"}
	if err := writer.Write(header); err != nil {
		writer.Flush() // Ensure flushing happens
		fmt.Printf("Failed to write header: %v", err)
		return
	}

	for logType, dateMap := range statistics {
		for date, count := range dateMap {
			row := []string{logType, date, fmt.Sprintf("%d", count)}
			if err := writer.Write(row); err != nil {
				log.Printf("Failed to write row: %v", err)
				return
			}
		}
	}
}
func OutputToConsole(statistics map[string]map[string]int) {
	for logType, dateMap := range statistics {
		fmt.Printf("Log Type: %s\n", logType)
		for date, count := range dateMap {
			fmt.Printf("  Date: %s, Count: %d\n", date, count)
		}
	}
}
