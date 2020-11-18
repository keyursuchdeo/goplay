package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := takeInput()
	bubbleSort(input)
	for _, num := range input {
		fmt.Printf("%v ", num)
	}
}

func swap(from int, to int, nums []int) {
	temp := nums[from]
	nums[from] = nums[to]
	nums[to] = temp
}

func bubbleSort(nums []int) {
	length := len(nums)

	for i := 0; i < length; i++ {
		for index, num := range nums {
			if index == length-i-1 {
				break
			} else {
				if num > nums[index+1] {
					swap(index, index+1, nums)
				}
			}
		}
	}
}

func takeInput() []int {
	var nums []int
	var input string
	fmt.Print("Please enter up to 10 integers:")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		input = scanner.Text()
	}
	numstrings := strings.Fields(input)
	if len(numstrings) > 10 {
		fmt.Printf("More than 10 nums entered %s\n", input)
	} else {
		for _, strnum := range numstrings {
			num, err := strconv.Atoi(strnum)
			if err == nil {
				nums = append(nums, num)
			} else {
				fmt.Printf("Invalid num %s\n", strnum)
			}
		}
	}
	return nums
}
