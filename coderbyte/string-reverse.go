package main

import (
	"bufio"
	"fmt"
	"os"
)

func FirstReverse(str string) string {
	var result string
	for _, s := range str {
		result = string(s) + result
	}
	return result

}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter string:")
	scanner.Scan()
	input := scanner.Text()
	fmt.Println(FirstReverse(input))
}
