package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"time"
)

func main() {
	sensorDataChan := make(chan float64, 10)
	processedDataChan := make(chan float64, 10)

	startTime := time.Now()
	fmt.Printf("Start time: %s\n", startTime.Format(time.RFC3339))

	go simulateSensorRead(sensorDataChan, time.Minute)
	go processSensorData(sensorDataChan, processedDataChan)

	for average := range processedDataChan {
		fmt.Printf("Received average: %f\n", average)
	}

	endTime := time.Now()
	fmt.Printf("End time: %s\n", endTime.Format(time.RFC3339))
}

func simulateSensorRead(sensorDataChan chan float64, duration time.Duration) {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	timeout := time.NewTimer(duration)
	defer timeout.Stop()

	for {
		select {
		case <-timeout.C:
			close(sensorDataChan)
			return
		case <-ticker.C:
			value, err := randFloat64()
			if err != nil {
				fmt.Println("Error generating random number:", err)
				close(sensorDataChan)
				return
			}
			select {
			case sensorDataChan <- value * 100:
			default:
				fmt.Println("sensorDataChan is full, skipping this value.")
			}
		}
	}
}

func processSensorData(sensorDataChan, processedDataChan chan float64) {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	var dataBatch []float64

	for {
		select {
		case data, ok := <-sensorDataChan:
			if !ok {
				select {
				case processedDataChan <- calculateAverage(dataBatch):
				default:
					fmt.Println("processedDataChan is full, skipping this batch.")
				}
				close(processedDataChan)
				return
			}
			dataBatch = append(dataBatch, data)
		case <-ticker.C:
			if len(dataBatch) > 0 {
				select {
				case processedDataChan <- calculateAverage(dataBatch):
				default:
					fmt.Println("processedDataChan is full, skipping this batch.")
				}
				dataBatch = []float64{}
			}
		}
	}
}

func calculateAverage(dataBatch []float64) float64 {
	var sum float64
	for _, value := range dataBatch {
		sum += value
	}
	return sum / float64(len(dataBatch))
}

func randFloat64() (float64, error) {
	n, err := rand.Int(rand.Reader, big.NewInt(100))
	if err != nil {
		return 0, err
	}
	return float64(n.Int64()), nil
}
