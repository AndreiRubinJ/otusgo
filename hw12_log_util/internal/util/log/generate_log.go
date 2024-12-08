package log

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"os"
	"sync"
	"time"

	"github.com/AndreiRubinJ/otusgo/hw12_log_util/internal/util/log/entry"
)

const fileName = "information/log/app/generated_logs.log"
const numThreads = 5
const entriesPerThread = 100

var logTypes = []string{"INFO", "DEBUG", "WARNING", "ERROR"}
var logMessages = []string{
	"User profile updated",
	"Start sequence completed",
	"Non-responsive service detected",
	"Failed to connect to database",
}

func generateRandomDate() (time.Time, error) {
	startUnix := time.Now().AddDate(-1, 0, 0).Unix()
	endUnix := time.Now().Unix()
	delta := endUnix - startUnix

	nBig, err := rand.Int(rand.Reader, big.NewInt(delta))
	if err != nil {
		return time.Time{}, err
	}
	sec := startUnix + nBig.Int64()
	return time.Unix(sec, 0), nil
}

func GenerateRandomLogEntry() (entry.LogEntry, error) {
	logTypeIndex, err := getSecureRandomIndex(int64(len(logTypes)))
	if err != nil {
		return nil, err
	}

	logMessageIndex, err := getSecureRandomIndex(int64(len(logMessages)))
	if err != nil {
		return nil, err
	}

	date, err := generateRandomDate()
	if err != nil {
		return nil, err
	}

	return &entry.BasicLogEntry{
		LogType: logTypes[logTypeIndex],
		Date:    date,
		Message: logMessages[logMessageIndex],
	}, nil
}

func getSecureRandomIndex(max int64) (int, error) {
	n, err := rand.Int(rand.Reader, big.NewInt(max))
	if err != nil {
		return 0, err
	}
	return int(n.Int64()), nil
}

func EnsureDir() error {
	dir := "information/log/app"
	return os.MkdirAll(dir, os.ModePerm)
}

func WriteLogEntries(wg *sync.WaitGroup) {
	defer wg.Done()

	if err := EnsureDir(); err != nil {
		fmt.Printf("Error creating directories: %v\n", err)
		return
	}

	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	for i := 0; i < entriesPerThread; i++ {
		logEntry, err := GenerateRandomLogEntry()
		if err != nil {
			fmt.Printf("Error generating log entry: %v\n", err)
			return
		}
		entryStr := fmt.Sprintf("%s %s %s\n",
			logEntry.GetType(),
			logEntry.GetDate().Format(time.RFC3339), logEntry.GetMessage())
		if _, err := file.WriteString(entryStr); err != nil {
			fmt.Printf("Error writing to file: %v\n", err)
			return
		}
	}
}

func GenerateLogs() {
	var wg sync.WaitGroup

	if _, err := os.Stat(fileName); err == nil {
		if err := os.Remove(fileName); err != nil && !os.IsNotExist(err) {
			fmt.Printf("Error removing old log file: %v\n", err)
			return
		}
	}

	for i := 0; i < numThreads; i++ {
		wg.Add(1)
		go WriteLogEntries(&wg)
	}

	wg.Wait()
	fmt.Println("Log generation complete.")
}

func GetFileNameByDefoult() string {
	return fileName
}
