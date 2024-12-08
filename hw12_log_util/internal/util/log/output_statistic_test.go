package log

import (
	"bytes"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOutputStatisticToFile(t *testing.T) {
	statistics := map[string]map[string]int{
		"INFO": {
			"2023-10-10": 2,
		},
		"ERROR": {
			"2023-10-11": 1,
		},
	}

	fileName := "test_log_output.csv"

	defer func() {
		if err := os.Remove(fileName); err != nil {
			t.Fatalf("Failed to remove test file: %v", err)
		}
	}()

	OutputStatisticToFile(statistics, fileName)

	content, err := os.ReadFile(fileName)
	assert.NoError(t, err)

	expectedContent := "Log Type,Date,Count\nINFO,2023-10-10,2\nERROR,2023-10-11,1\n"
	assert.Equal(t, expectedContent, string(content))
}

func TestOutputToConsole(t *testing.T) {
	statistics := map[string]map[string]int{
		"INFO": {
			"2023-10-10": 2,
		},
		"ERROR": {
			"2023-10-11": 1,
		},
	}

	r, w, err := os.Pipe()
	if err != nil {
		t.Fatalf("failed to create pipe: %v", err)
	}

	originalStdout := os.Stdout
	defer func() { os.Stdout = originalStdout }()

	os.Stdout = w

	outC := make(chan string)
	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, r)
		outC <- buf.String()
	}()
	OutputToConsole(statistics)
	w.Close()
	output := <-outC
	expectedOutput := "Log Type: INFO\n  Date: 2023-10-10, Count: 2\nLog Type: ERROR\n  Date: 2023-10-11, Count: 1\n"
	assert.Equal(t, expectedOutput, output)
}
