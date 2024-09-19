package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/AndreiRubinJ/otusgo/hw07_word_counter/internal/utils"
)

func main() {
	fmt.Println("please type any string:")
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	fmt.Println("===========================Count Words============================")
	utils.CountWordsStatistic(text, utils.CountWords)
	fmt.Println("===========================Count Words Plural============================")
	utils.CountWordsStatistic(text, utils.CountWordsPlural)
}
