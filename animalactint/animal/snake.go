package animal

import "fmt"

type Snake struct {
	name       string
	food       string
	locomotion string
	noise      string
}

func InitSnake(name string) Snake {
	return Snake{name, "mice", "slither", "hsss"}
}

func (snake Snake) Eat() {
	fmt.Println(snake.food)
}

func (snake Snake) Move() {
	fmt.Println(snake.locomotion)
}

func (snake Snake) Speak() {
	fmt.Println(snake.noise)
}
