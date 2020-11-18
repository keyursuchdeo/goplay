package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	var originalLen = 3
	const MaxValue = 2147483647
	var slice = []int{MaxValue, MaxValue, MaxValue}
	var index = 0
	for {
		var inputStr string
		fmt.Println("Please enter an integer. Enter X to stop ")
		_, err := fmt.Scan(&inputStr)
		if err == nil {
			if inputStr == "X" {
				break
			} else {
				input, e := strconv.Atoi(inputStr)
				if e == nil {
					if index < originalLen {
						slice[index] = input
						index = index + 1
						sort.Ints(slice)
						fmt.Printf("%v", slice[0:index])
					} else {
						slice = append(slice, input)
						sort.Ints(slice)
						fmt.Printf("%v", slice)
					}
				} else {
					fmt.Printf("Invalid input %s", inputStr)
					continue
				}
			}
		}

	}
}
