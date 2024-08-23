package main

import (
	"fmt"
	"strings"
)

const (
	spaceCount = 1
	spaceType  = "\t"
)

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

func main() {
	printBoard(8, 8)
}
