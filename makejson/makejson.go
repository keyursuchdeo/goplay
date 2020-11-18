package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var name string
	var address string
	var person map[string]string
	fmt.Print("Please enter name: ")
	_, err := fmt.Scan(&name)
	if err == nil {
		fmt.Print("Please enter address: ")
		_, err := fmt.Scan(&address)
		if err == nil {
			person = make(map[string]string)
			person["name"] = name
			person["address"] = address
			jsonPerson, err := json.Marshal(person)
			if err == nil {
				fmt.Println(string(jsonPerson))
			} else {
				fmt.Printf("Error occurred while marshaling person %s", person)
			}
		} else {
			fmt.Printf("Error %s while reading address", err)
		}
	} else {
		fmt.Printf("Error %s while reading name", err)
	}
}
