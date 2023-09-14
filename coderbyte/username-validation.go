package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func CodelandUsernameValidation(str string) string {

	// Rule 1: Check if the length is between 4 and 25 characters.
	if len(str) < 4 || len(str) > 25 {
		return "false"
	}

	// Rule 2: Check if it starts with a letter.
	if match, _ := regexp.MatchString("^[a-zA-Z]", str); !match {
		return "false"
	}

	// Rule 3: Check if it contains only letters, numbers, and underscores.
	if match, _ := regexp.MatchString("^[a-zA-Z0-9_]+$", str); !match {
		return "false"
	}

	// Rule 4: Check if it doesn't end with an underscore.
	if str[len(str)-1] == '_' {
		return "false"
	}

	// If all rules pass, return true.
	return "true"

}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter string:")
	scanner.Scan()
	input := scanner.Text()
	fmt.Println(CodelandUsernameValidation(input))

}
