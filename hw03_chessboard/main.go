package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	height int
	length int
)

const (
	spaceCount    = 1
	spaceType     = "\t"
	msgHeight     = "Please enter board 'height' (from 2 to 25)"
	msgLength     = "Please enter board 'length' (from 2 to 25)"
	maxNumber     = 25
	minNumber     = 2
	maxRetryCount = 3
)

func init() {
	start()
}

func printBoard(height int, length int) {
	printHeader(length)
	for i := 1; i <= height; i++ {
		if i <= height {
			printRowBoundary(length)
			printRow(i, length)
		}
	}
	printFooter(height, length)
}

func printHeader(length int) {
	printLetter(length)
}

func printLetter(length int) {
	printTub()
	for j := 0; j < length; j++ {
		letter := byte('A') + byte(j)
		fmt.Printf("  %c ", letter)
	}
	fmt.Println("")
}

func printRowBoundary(length int) {
	printTub()
	fmt.Println(strings.Repeat(" ---", length))
}

func printRow(rowNum int, length int) {
	fmt.Print(rowNum)
	printTub()
	fmt.Print("|")
	for j := 0; j < length; j++ {
		if (rowNum+j)%2 == 0 {
			fmt.Print("   |")
		} else {
			fmt.Print(" X |")
		}
	}
	printTub()
	fmt.Print(rowNum)
	fmt.Println("")
}

func printFooter(_ int, length int) {
	printRowBoundary(length)
	printLetter(length)
}

func printTub() {
	fmt.Print(strings.Repeat(spaceType, spaceCount))
}

func convertParam(param string) (i int, err error) {
	intParam, err := strconv.Atoi(param)
	if err != nil {
		return 0, err
	}
	return intParam, nil
}

func readPromptValue(prompt string) string {
	fmt.Println(prompt)
	reader := bufio.NewReader(os.Stdin)
	value, _ := reader.ReadString('\n')
	return strings.TrimSpace(value)
}

func getBoardParam(param *int, msg string) {
	var err error
	retryCount := 1
	for {
		value := readPromptValue(msg)
		*param, err = convertParam(value)
		if retryCount >= maxRetryCount {
			fmt.Println("You put 3 time wrong value \n program exit")
			os.Exit(0)
		} else if err != nil || *param < minNumber || *param > maxNumber {
			fmt.Println("You putted wrong value")
			retryCount++
			continue
		}
		break
	}
}

func start() {
	getBoardParam(&height, msgHeight)
	getBoardParam(&length, msgLength)
}

func main() {
	if height > 0 && length > 0 {
		printBoard(height, length)
	}
}
