package main

import (
	"fmt"
	"strconv"
)

type Animal interface {
	Eat(thing interface{})
}

type Human struct {
	weight uint32
}

type Food interface {
	GetWeight() uint32
	GetName() string
}

type Apple struct {
	weight uint32
}

func (this *Apple) GetWeight() uint32 {
	return this.weight
}

func (this *Apple) GetName() string {
	return "apple"
}

// func (this Apple) GetName() string {
// 	return "apple"
// }

func (this Human) Eat(thing Food) {
	fmt.Println("Human eat " + thing.GetName() + ", weight:" + strconv.Itoa(int(thing.GetWeight())))
}

// func (this *Human) Eat(thing Food) {
// 	fmt.Println("Human eat " + thing.GetName() + ", weight:" + strconv.Itoa(int(thing.GetWeight())))
// }

// func (this *Human) Eat(thing *Food) {
// 	fmt.Println("Human eat " + thing.GetName() + ", weight:" + strconv.Itoa(int(thing.GetWeight())))
// }

func main() {
	var human Human
	apple := &Apple{weight: 500}
	human.Eat(apple)
}
