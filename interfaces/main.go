package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
}

type Circle struct {
	x, y, radius float64
}

type Rectangle struct {
	width, height float64
}

func (c Circle) area() float64 {
	return math.Pi *c.radius *c.radius
}

func (r Rectangle) area() float64 {
	return r.width * r.height
}

func getArea(s Shape) float64 {
	return s.area()
}

func main() {

	circle1 := Circle{0, 0, 1}
	rectangle1 := Rectangle{2, 4}

	fmt.Printf("Circle Area %f\n", getArea(circle1))
	fmt.Printf("Reactancle Area %f\n", getArea(rectangle1))
	
}
