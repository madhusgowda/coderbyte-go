package main

import (
	"fmt"
	"log"
)

func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int)

	for i, num := range nums {
		complement := target - num
		if index, ok := numMap[complement]; ok {
			return []int{index, i}
		}
		numMap[num] = i
	}
	return nil
}

func main() {
	var target, arraySize int
	var err error

	fmt.Println("Enter size of array:")
	_, err = fmt.Scan(&arraySize)
	if err != nil {
		log.Fatal(err.Error())
	}

	nums := make([]int, arraySize)

	fmt.Println("Enter the value of an array:")
	for i := 0; i < arraySize; i++ {
		_, err = fmt.Scan(&nums[i])
		if err != nil {
			log.Fatal(err.Error())
		}
	}
	fmt.Println("Enter the target value:")
	_, err = fmt.Scan(&target)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(twoSum(nums, target))
}
