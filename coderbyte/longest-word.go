package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func LongestWord(sen string) string {
	words := strings.Fields(sen)
	result, length := "", 0
	for _, word := range words {
		cleanedWord := strings.TrimFunc(word, func(r rune) bool {
			return !((r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9'))
		})

		wordLength := len(cleanedWord)
		if wordLength > length {
			result = cleanedWord
			length = wordLength
		}
	}
	return result

}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter string:")
	scanner.Scan()
	input := scanner.Text()
	fmt.Println(LongestWord(input))

}
