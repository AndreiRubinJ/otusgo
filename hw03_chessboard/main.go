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
	msgToPrompt   = "Please enter board '%s' (from 2 to 25) \n"
	maxNumber     = 25
	minNumber     = 2
	maxRetryCount = 3
)

func printBoard(height int, length int) {
	printHeader(length)
	for i := 1; i <= height; i++ {
		printRowBoundary(length)
		printRow(i, length)
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
	fmt.Println(rowNum)
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

func readPromptValue(paramType string) string {
	fmt.Printf(msgToPrompt, paramType)
	reader := bufio.NewReader(os.Stdin)
	value, _ := reader.ReadString('\n')
	return strings.TrimSpace(value)
}

func getBoardParam(param *int, paramType string) {
	var err error
	var retryCount = 1
	for {
		value := readPromptValue(paramType)
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
	getBoardParam(&height, "height")
	getBoardParam(&length, "length")
	printBoard(height, length)
}

func main() {
	start()
}
