package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"time"
)

func main() {
	sensorDataChan := make(chan float64)
	processedDataChan := make(chan float64)

	go simulateSensorRead(sensorDataChan)
	go processSensorData(sensorDataChan, processedDataChan)

	for average := range processedDataChan {
		fmt.Printf("Received average: %f\n", average)
	}
}

func simulateSensorRead(sensorDataChan chan float64) {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	timeout := time.After(time.Minute)
	for {
		select {
		case <-timeout:
			close(sensorDataChan)
			return
		case <-ticker.C:
			value, err := randFloat64()
			if err != nil {
				fmt.Println("Error generating random number:", err)
				close(sensorDataChan)
				return
			}
			sensorDataChan <- value * 100
		}
	}
}

func processSensorData(sensorDataChan, processedDataChan chan float64) {
	dataBatch := make([]float64, 0, 10)
	for data := range sensorDataChan {
		dataBatch = append(dataBatch, data)
		if len(dataBatch) == 10 {
			var sum float64
			for _, value := range dataBatch {
				sum += value
			}
			average := sum / float64(len(dataBatch))
			processedDataChan <- average
			dataBatch = []float64{}
		}
	}
	if len(dataBatch) > 0 {
		var sum float64
		for _, value := range dataBatch {
			sum += value
		}
		average := sum / float64(len(dataBatch))
		processedDataChan <- average
	}
	close(processedDataChan)
}
func randFloat64() (float64, error) {
	n, err := rand.Int(rand.Reader, big.NewInt(100))
	if err != nil {
		return 0, err
	}
	return float64(n.Int64()), nil
}
