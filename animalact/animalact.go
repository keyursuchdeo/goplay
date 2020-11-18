package main

import (
	"errors"
	"fmt"
)

type Animal struct {
	food     string
	location string
	noise    string
}

func initAnimal(name string) (Animal, error) {
	if name == "cow" {
		return initCow(), nil
	} else if name == "bird" {
		return initBird(), nil
	} else if name == "snake" {
		return initSnake(), nil
	} else {
		return Animal{}, errors.New("Invalid animal")
	}
}

func initCow() Animal {
	return Animal{"grass", "walk", "moo"}
}

func initBird() Animal {
	return Animal{"worms", "fly", "peep"}
}

func initSnake() Animal {
	return Animal{"mice", "slither", "hsss"}
}

func (animal Animal) Act(action string) error {
	if action == "eat" {
		animal.Eat()
		return nil
	} else if action == "move" {
		animal.Move()
		return nil
	} else if action == "speak" {
		animal.Speak()
		return nil
	} else {
		return errors.New("Invalid action")
	}
}

func (animal Animal) Eat() {
	fmt.Println(animal.food)
}

func (animal Animal) Move() {
	fmt.Println(animal.location)
}

func (animal Animal) Speak() {
	fmt.Println(animal.noise)
}

func main() {
	var animalName string
	var action string
	for {
		fmt.Print("Please enter your request in 'animal action' format: ")
		_, err := fmt.Scan(&animalName, &action)
		if err == nil {
			animal, err := initAnimal(animalName)
			if err == nil {
				err := animal.Act(action)
				if err == nil {
					continue
				} else {
					fmt.Printf("Error occurred while acting: %s\n", err)
				}
			} else {
				fmt.Printf("Error occurred while initializing animal: %s\n", err)
			}
		} else {
			fmt.Printf("Error occurred while scanning input: %s\n", err)
		}
	}
}
