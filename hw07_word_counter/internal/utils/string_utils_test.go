package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// checkIsNotTheWorld проверяет, является ли символ разделителем (не буква и не цифра).

// countWords принимает строку текста и возвращает мапу, содержащую количество каждого слова в тексте.

func TestCountWords(t *testing.T) {
	tests := []struct {
		input    string
		expected map[string]int
	}{
		{"hello world.", map[string]int{"hello": 1, "world": 1}},
		{"hello world world!", map[string]int{"hello": 1, "world": 2}},
		{"hello 123", map[string]int{"hello": 1}},
		{"1hello 123", map[string]int{"hello": 1}},
		{"1hel11lo 123", map[string]int{"hel": 1, "lo": 1}},
		{"hello HELLO", map[string]int{"hello": 2}},
		{"a a b b", map[string]int{}},
		{"1231", map[string]int{}},
		{"привет, мир!", map[string]int{"привет": 1, "мир": 1}},
		{"привет, мир! привет! мир! ", map[string]int{"привет": 2, "мир": 2}},
		{"   привет, мир! привет! мир! ", map[string]int{"привет": 2, "мир": 2}},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, CountWords(test.input))
	}
}

func TestCheckIsNotTheWorld(t *testing.T) {
	tests := []struct {
		input    rune
		expected bool
	}{
		{'a', false},
		{'1', true},
		{' ', true},
		{'$', true},
		{'@', true},
		{'.', true},
		{',', true},
		{'!', true},
		{'?', true},
		{':', true},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, checkIsNotTheWorld(test.input))
	}
}
func TestCountWordsPlural(t *testing.T) {
	tests := []struct {
		input    string
		expected map[string]int
	}{
		{"hello world", map[string]int{"hello": 1, "world": 1}},
		{"hello 123", map[string]int{"hello": 1}},
		{"hello HELLO", map[string]int{"hello": 2}},
		{"a a b b", map[string]int{}},
		{"", map[string]int{}},
		{"1323", map[string]int{}},
		{"привет, мир!", map[string]int{"привет": 1, "мир": 1}},
		{"привет, мир! привет! мир! ", map[string]int{"привет": 2, "мир": 2}},
		{"   !привет, мир! привет! мир! ", map[string]int{"привет": 2, "мир": 2}},
		{"  @@@@ !привет, мир! привет! мир! ", map[string]int{"привет": 2, "мир": 2}},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, CountWordsPlural(test.input))
	}
}
