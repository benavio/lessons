package main

import "fmt"

type Share interface {
	Area() float64
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return c.Radius * c.Radius * 3.14
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width * 0.5
}

func main() {
	fmt.Println(Share.Area(Circle{3.12}))
	fmt.Println(Share.Area(Rectangle{3.12, 1.12}))
}
