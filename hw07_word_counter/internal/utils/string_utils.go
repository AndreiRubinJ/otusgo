package utils

import (
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"
)

func checkIsNotTheWorld(symbol rune) bool {
	return !unicode.IsLetter(symbol) || unicode.IsNumber(symbol)
}

func setStatisticWord(word string, wordCount *map[string]int) {
	word = strings.ToLower(word)
	if utf8.RuneCountInString(word) > 1 {
		(*wordCount)[word]++
	}
}

func CountWords(text string) map[string]int {
	wordCount := make(map[string]int)
	words := strings.FieldsFunc(text, checkIsNotTheWorld)

	for _, word := range words {
		setStatisticWord(word, &wordCount)
	}
	return wordCount
}
func CountWordsPlural(text string) map[string]int {
	wordCount := make(map[string]int)
	words := strings.Split(text, " ")
	for _, word := range words {
		for _, char := range word {
			if checkIsNotTheWorld(char) {
				word = strings.ReplaceAll(word, string(char), "")
			}
		}
		setStatisticWord(word, &wordCount)
	}
	return wordCount
}

func CountWordsStatistic(text string, funcCount func(string) map[string]int) {
	wordCount := funcCount(text)
	for word, count := range wordCount {
		fmt.Printf("World: '%s', Count: %d\n", word, count)
	}
}
