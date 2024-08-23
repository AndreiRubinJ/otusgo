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
	spaceCount = 1
	spaceType  = "\t"
	msgHeight  = "Please enter board height (from 2 to 10)"
	msgLength  = "Please enter board length (from 2 to 10)"
)

func init() {
	start()
}

func printBoard(height int, length int) {
	for i := 0; i <= height+1; i++ {
		switch {
		case i == 0:
			printHeader(length)
		case i <= height:
			printRowBoundary(length)
			printRow(i, length)
		default:
			printFooter(height, length)
		}
		fmt.Println("")
	}
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
	count := 0
	for {
		value := readPromptValue(msg)
		*param, err = convertParam(value)
		if count > 2 {
			fmt.Println("You put more then 3 time wrong value \n program exit")
			os.Exit(0)
		} else if err != nil || *param < 2 || *param > 10 {
			fmt.Println("You putted wrong value")
			count++
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
