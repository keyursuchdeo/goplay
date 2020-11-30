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
	ch := make(chan []int)
	go mergeSort(input, ch)
	output := <-ch
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

func mergeSort(input []int, channel chan []int) {
	length := len(input)
	if length <= 1 {
		channel <- input
	} else {
		chLeft := make(chan []int)
		chRight := make(chan []int)
		go mergeSort(input[:length/2], chLeft)
		go mergeSort(input[length/2:], chRight)
		left := <-chLeft
		right := <-chRight
		mergedOutput := merge(left, right)
		// fmt.Println(mergedOutput)
		channel <- mergedOutput
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
