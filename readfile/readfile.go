package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	type Person struct {
		fName string
		lName string
	}
	var persons []Person
	var fileName string

	fmt.Print("Please enter file name containing names: ")
	_, err := fmt.Scan(&fileName)
	if err == nil {
		file, err := os.Open(fileName)
		if err == nil {
			scanner := bufio.NewScanner(file)
			for scanner.Scan() {
				line := scanner.Text()
				rawName := strings.Split(line, " ")
				persons = append(persons, Person{fName: rawName[0], lName: rawName[1]})
			}
			for _, person := range persons {
				fmt.Printf("%s %s\n", person.fName, person.lName)
			}
		} else {
			fmt.Printf("Error %s while opening file %s\n", err, fileName)
		}
	} else {
		fmt.Printf("Error %s while reading file name\n", err)
	}

}
