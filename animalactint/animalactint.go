package main

import (
	"animalactint/animal"
	"fmt"
)

func main() {
	var command string
	var animalName string
	var animalType string
	var info string
	var animals = make(map[string]animal.Animal)

	for {
		fmt.Print("Please enter your command: ")
		_, err := fmt.Scan(&command)
		if err == nil {
			if command == "newanimal" {
				fmt.Print("Please enter animal name and animal type: ")
				_, err := fmt.Scan(&animalName, &animalType)
				if err == nil {
					if animals[animalName] == nil {
						if animalType == "cow" {
							animals[animalName] = animal.InitCow(animalName)
							fmt.Println("Created it")
						} else if animalType == "bird" {
							animals[animalName] = animal.InitBird(animalName)
							fmt.Println("Created it")
						} else if animalType == "snake" {
							animals[animalName] = animal.InitSnake(animalName)
							fmt.Println("Created it")
						} else {
							fmt.Printf("Invalid animal type %s\n", animalType)
						}
					} else {
						fmt.Printf("Animal by name %s is already present %s\n", animalName, animals[animalName])
					}
				} else {
					fmt.Printf("Error while reading animal name and type %s", err)
				}
			} else if command == "query" {
				fmt.Print("Please enter animal name and information required: ")
				_, err := fmt.Scan(&animalName, &info)
				if err == nil {
					if animals[animalName] == nil {
						fmt.Printf("Animal by name %s doesn't exist\n", animalName)
					} else {
						animal := animals[animalName]
						if info == "eat" {
							animal.Eat()
						} else if info == "move" {
							animal.Move()
						} else if info == "speak" {
							animal.Speak()
						} else {
							fmt.Printf("Invalid info requested %s\n", info)
						}
					}
				} else {
					fmt.Printf("Error while reading animal name and information %s", err)
				}
			} else {
				fmt.Printf("Invalid command %s\n", command)
			}
		} else {
			fmt.Printf("Invalid input %s %s %s\n", command, animalName, animalType)
		}
	}
}
