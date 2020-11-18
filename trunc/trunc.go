package main

import (
	"fmt"
	"strconv"
)

func main() {
	var input float64
	fmt.Printf("Enter a floating point number: ")
	_, err := fmt.Scan(&input)
	if err == nil {
		fmt.Println(strconv.Itoa(int(input)))
	} else {
		fmt.Println("Please enter a valid floating point number")
	}
}
