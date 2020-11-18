package main

import "fmt"

func main() {
	acc, initV, initD, time, err := readInput()
	if err == nil {
		displacementFn := GenDisplaceFn(acc, initV, initD)
		fmt.Println(displacementFn(time))
	} else {
		fmt.Printf("Error %s while reading input\n", err)
	}
}

func GenDisplaceFn(acc float64, initV float64, initD float64) func(float64) float64 {
	return func(time float64) float64 {
		return (0.5 * acc * time * time) + (initV * time) + initD
	}
}

func readInput() (float64, float64, float64, float64, error) {
	var acceleration float64
	var initVelocity float64
	var initDisplacement float64
	var time float64

	fmt.Print("Please enter acceleration: ")
	_, err := fmt.Scan(&acceleration)
	if err == nil {
		fmt.Print("Please enter initial value of velocity: ")
		_, err := fmt.Scan(&initVelocity)
		if err == nil {
			fmt.Print("Please enter initial value of displacement: ")
			_, err := fmt.Scan(&initDisplacement)
			if err == nil {
				fmt.Print("Please enter time: ")
				_, err := fmt.Scan(&time)
				if err == nil {
					return acceleration, initVelocity, initDisplacement, time, nil
				} else {
					return 0, 0, 0, 0, err
				}
			} else {
				return 0, 0, 0, 0, err
			}
		} else {
			return 0, 0, 0, 0, err
		}
	} else {
		return 0, 0, 0, 0, err
	}
}
