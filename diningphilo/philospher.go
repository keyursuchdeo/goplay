package main

import (
	"fmt"
	"sync"
)

type Chopstick struct {
	sync.Mutex
}

type Philosopher struct {
	id    int
	left  *Chopstick
	right *Chopstick
}

func (p *Philosopher) pick() {
	p.left.Lock()
	p.right.Lock()
}

func (p *Philosopher) keep() {
	p.left.Unlock()
	p.right.Unlock()
}

func (p *Philosopher) Eat(numOfTimesToEat int, ch chan struct{}, wg *sync.WaitGroup) {
	for i := 0; i < numOfTimesToEat; i++ {
		p.pick()
		fmt.Printf("Starting to eat %v\n", p.id)
		fmt.Printf("Finishing eating %v\n", p.id)
		p.keep()
	}
	<-ch
	wg.Done()
}
