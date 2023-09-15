package main

import (
	"fmt"
	"log"
)

func addDigits(num int) int {
	var result int
	for {
		for num > 0 {
			result += num % 10
			num /= 10
		}
		num, result = result, 0
		if num < 10 {
			break
		}
	}
	return num
}

func main() {
	var num int
	fmt.Println("Enter a value:")
	_, err := fmt.Scan(&num)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(addDigits(num))
}
