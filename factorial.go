package main

import (
	"fmt"
	"log"
)

func FirstFactorial(num int) int {

	if num == 1 || num == 0 {
		return num
	}
	return num * FirstFactorial(num-1)

}

func main() {
	var input int
	fmt.Println("First Factorial:")
	_, err := fmt.Scanln(&input)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(FirstFactorial(input))
}
