package main

import (
	"fmt"

	"github.com/AndreiRubinJ/otusgo/hw08_binary_search/internal/utils"
)

func main() {
	arr := []int{0, 2, 3, 4, 10, 40, 60, 100, 250, 1000}
	value := 55
	result := utils.SearchValueFromSlice(arr, value)
	if result != -1 {
		fmt.Printf("Element found at index %d\n", result)
	} else {
		fmt.Println("Element not found in array")
	}
}
