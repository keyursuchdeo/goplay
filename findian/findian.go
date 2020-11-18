package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var input string

	fmt.Printf("Please enter a string:")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		input = scanner.Text()
	}
	lowerInput := strings.ToLower(strings.Trim(input, " "))
	if strings.HasPrefix(lowerInput, "i") && strings.ContainsRune(lowerInput, 'a') && strings.HasSuffix(lowerInput, "n") {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}
