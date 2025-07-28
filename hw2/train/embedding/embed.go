package main

// НАСЛЕДОВАНИЕ

import "fmt"

type Vechicle struct{}

type Car struct {
	Vechicle
}

func (v *Vechicle) Start() {
	fmt.Println("we are started...")
}

func (c *Car) Drive() {
	fmt.Println("we are Driving...")
}

func main() {
	car := Car{}

	car.Start()
	car.Drive()

}
