package animal

import "fmt"

type Cow struct {
	name       string
	food       string
	locomotion string
	noise      string
}

func InitCow(name string) Cow {
	return Cow{name, "grass", "walk", "moo"}
}

func (cow Cow) Eat() {
	fmt.Println(cow.food)
}

func (cow Cow) Move() {
	fmt.Println(cow.locomotion)
}

func (cow Cow) Speak() {
	fmt.Println(cow.noise)
}
