package l121

import "fmt"

type Creature interface {
	Say()
}

type Human struct{}

func (h *Human) Say() {
	fmt.Println("Says something...")
}

type Dog struct{}

func (d *Dog) Bark() {
	fmt.Println("Woof-woof!")
}

type DogAdapter struct {
	DogExample *Dog
}

func (da *DogAdapter) Say() {
	fmt.Println("Adapter transform dog's barking to human speech")
	da.DogExample.Bark()
}
