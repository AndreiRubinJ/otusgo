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

func CountWords(text string) map[string]int {
	wordCount := make(map[string]int)
	words := strings.FieldsFunc(text, checkIsNotTheWorld)

	for _, word := range words {
		word = strings.ToLower(word)
		if utf8.RuneCountInString(word) > 1 {
			wordCount[word]++
		}
	}
	return wordCount
}
func CountWordsPlural(text string) map[string]int {
	wordCount := make(map[string]int)
	words := strings.Split(text, " ")
	for _, word := range words {
		isNotTheWorld := false
		for idx, char := range word {
			if !unicode.IsLetter(char) || unicode.IsNumber(char) {
				if idx == 0 || idx == len(word)-1 {
					word = word[:idx] + word[idx+1:]
					continue
				}
				isNotTheWorld = true
				break
			}
		}
		if !isNotTheWorld {
			word = strings.ToLower(word)
			if utf8.RuneCountInString(word) > 1 {
				wordCount[word]++
			}
		}
	}
	return wordCount
}

func CountWordsStatistic(text string, funcCount func(string) map[string]int) {
	wordCount := funcCount(text)
	for word, count := range wordCount {
		fmt.Printf("World: '%s', Count: %d\n", word, count)
	}
}
