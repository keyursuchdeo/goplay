package animal

import "fmt"

type Bird struct {
	name       string
	food       string
	locomotion string
	noise      string
}

func InitBird(name string) Bird {
	return Bird{name, "worms", "fly", "peep"}
}

func (bird Bird) Eat() {
	fmt.Println(bird.food)
}

func (bird Bird) Move() {
	fmt.Println(bird.locomotion)
}

func (bird Bird) Speak() {
	fmt.Println(bird.noise)
}
