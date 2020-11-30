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
	output := mergeSort(input)
	fmt.Println(output)
}

func merge(input1 []int, input2 []int) []int {
	output := make([]int, 0, len(input1)+len(input2))
	for len(input1) > 0 || len(input2) > 0 {
		if len(input1) == 0 {
			return append(output, input2...)
		} else if len(input2) == 0 {
			return append(output, input1...)
		} else if input1[0] <= input2[0] {
			output = append(output, input1[0])
			input1 = input1[1:]
		} else {
			output = append(output, input2[0])
			input2 = input2[1:]
		}
	}
	return output
}

func mergeSort(input []int) []int {
	length := len(input)
	if length <= 1 {
		return input
	} else {
		left := mergeSort(input[:length/2])
		right := mergeSort(input[length/2:])
		return merge(left, right)
	}
}

func takeInput() []int {
	var nums []int
	var input string
	fmt.Print("Please enter integers separated by space:")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		input = scanner.Text()
	}
	numstrings := strings.Fields(input)
	for _, strnum := range numstrings {
		num, err := strconv.Atoi(strnum)
		if err == nil {
			nums = append(nums, num)
		} else {
			fmt.Printf("Invalid num %s\n", strnum)
		}
	}
	return nums
}
