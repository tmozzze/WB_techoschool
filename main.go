package main

import (
	l121 "github.com/tmozzze/WB_techoschool/L1/L1_21"
)

func main() {
	var c1, c2 l121.Creature

	c1 = &l121.Human{}
	c1.Say()

	dog := &l121.Dog{}
	dog.Bark()

	c2 = &l121.DogAdapter{DogExample: dog}
	c2.Say()

}
