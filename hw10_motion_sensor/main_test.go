package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestProcessSensorData(t *testing.T) {
	sensorDataChan := make(chan float64)
	processedDataChan := make(chan float64)

	go processSensorData(sensorDataChan, processedDataChan)

	go func() {
		for i := 1; i <= 20; i++ {
			sensorDataChan <- float64(i)
		}
		close(sensorDataChan)
	}()

	expectedAverages := []float64{5.5, 15.5}

	var receivedAverages []float64
	for {
		select {
		case avg, ok := <-processedDataChan:
			if !ok {
				assert.Equal(t, expectedAverages, receivedAverages, "Expected and received averages do not match")
				return
			}
			receivedAverages = append(receivedAverages, avg)
		case <-time.After(time.Second):
			t.Fatal("Timeout waiting for processed data")
		}
	}
}
