package main

import (
	"fmt"
	"sync"
)

func main() {
	var numOfPhilos int
	var numOfTimesToEat int
	var numOfPhilosEatingConc int
	fmt.Print("Please enter number of philosophers: ")
	_, err := fmt.Scan(&numOfPhilos)
	if err == nil {
		fmt.Print("Please enter number of times philosophers can eat: ")
		_, err := fmt.Scan(&numOfTimesToEat)
		if err == nil {
			fmt.Print("Please enter number of philosophers eating concurrently: ")
			_, err := fmt.Scan(&numOfPhilosEatingConc)
			if err == nil {
				philos := initializePhilosophers(numOfPhilos)
				wg := sync.WaitGroup{}
				wg.Add(numOfPhilos)
				go orchestrateEating(philos, numOfPhilosEatingConc, numOfTimesToEat, &wg)
				wg.Wait()
			} else {
				fmt.Printf("Error %s while reading number of philosophers eating concurrently\n", err)
			}
		} else {
			fmt.Printf("Error %s while reading number of philosophers eating concurrently\n", err)
		}
	} else {
		fmt.Printf("Error %s while reading number of philosophers\n", err)
	}
}

func orchestrateEating(philos []*Philosopher, concPhilos int, numOfTimesToEat int, wg *sync.WaitGroup) {
	eatingChan := make(chan struct{}, concPhilos)
	for _, philo := range philos {
		eatingChan <- struct{}{}
		go philo.Eat(numOfTimesToEat, eatingChan, wg)
	}
}

func initializeChopsticks(philoCount int) []*Chopstick {
	chopsticks := make([]*Chopstick, philoCount)
	for i := 0; i < philoCount; i++ {
		chopsticks[i] = new(Chopstick)
	}
	return chopsticks
}

func initializePhilosophers(philoCount int) []*Philosopher {
	chopsticks := initializeChopsticks(philoCount)
	philos := make([]*Philosopher, philoCount)
	for i := 0; i < philoCount; i++ {
		philos[i] = &Philosopher{i, chopsticks[i], chopsticks[(i+1)%philoCount]}
	}
	return philos
}
