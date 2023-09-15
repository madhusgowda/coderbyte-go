package main

import (
	"fmt"
	"log"
)

func main() {
	var input string
	fmt.Println("Enter a value:")
	_, err := fmt.Scan(&input)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(lengthOfLongestSubstring(input))
}

func lengthOfLongestSubstring(s string) int {
	location := [256]int{}
	for i := range location {
		location[i] = -1
	}
	maxLen, left := 0, 0
	for i := 0; i < len(s); i++ {
		if location[s[i]] >= left {
			left = location[s[i]] + 1
		} else if i+1-left > maxLen {
			maxLen = i + 1 - left
		}
		location[s[i]] = i
	}

	return maxLen
}

//func lengthOfLongestSubstring(s string) int {
//	if len(s) == 1 {
//		return 1
//	}
//	startMax, endMax, tmpStart, tmpIndex := 0, 0, 0, 0
//	for i := 1; i < len(s); i++ {
//		tmpIndex = -1
//		for j := tmpStart; j < i; j++ {
//			if s[j] == s[i] {
//				tmpIndex = j
//				break
//			}
//		}
//		if tmpIndex >= 0 {
//			if (i - tmpStart - tmpIndex) > (endMax - startMax) {
//				startMax = tmpIndex
//				endMax = i
//			}
//			tmpStart = tmpIndex + 1
//		} else {
//			if (i - tmpStart + 1) > (endMax - startMax) {
//				startMax = tmpStart
//				endMax = i + 1
//			}
//		}
//	}
//	return endMax - startMax
//}
