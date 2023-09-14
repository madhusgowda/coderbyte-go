package main

import (
	"bufio"
	"fmt"
	"os"
)

import "strconv"
import "strings"

func QuestionsMarks(str string) string {

	chars := strings.Split(str, "")

	qCount := 0
	lastNum := -1
	pairs := 0
	valid := 0

	// 	// fmt.Println(QuestionsMarks("acc?7??sss?3rr1??????5"))
	for _, c := range chars {

		num, err := strconv.Atoi(c)

		// we have a letter
		if err != nil {
			if c != "?" {
				continue
			}

			if lastNum != -1 {
				qCount++
			}

			continue
		}

		if lastNum+num == 10 {
			pairs++

			if qCount == 3 {
				valid++
			}
		}
		// reset
		qCount = 0
		lastNum = num
	}

	if pairs > 0 && pairs == valid {
		return "true"
	}
	return "false"

}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter string:")
	scanner.Scan()
	input := scanner.Text()
	fmt.Println(QuestionsMarks(input))

}
