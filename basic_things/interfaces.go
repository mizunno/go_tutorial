package main

import (
	"fmt"
	"math"
)

// A type implements an interface by implementing its methods.
// There is no explicit declaration of intent, no "implements"
// keyword. Implicit interfaces decouple the definition of an
// interface from its implementation, which could then appear
// in any package without prearrangement.

type Shape interface {
    Area()			float64
	IncreaseSize(delta float64)
	fmt.Stringer // Interface for printables
}

type Circle struct {
	radius	float64
}

func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (c Circle) String() string {
	return fmt.Sprintf("Circle {Radius: %.2f}", c.radius)
}

func (c *Circle) IncreaseSize(delta float64) {
	c.radius += delta
}

type Rectangle struct {
	width 	float64
	height	float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func (r Rectangle) String() string {
	return fmt.Sprintf("Rectangle {Width: %.2f} {Height: %.2f}", r.width, r.height)
}

func (r *Rectangle) IncreaseSize(delta float64) {
	r.height += delta
	r.width += delta
}


func main() {
	c := &Circle{
		radius: 5,
	}

	fmt.Printf("%T\n", c)

	r := &Rectangle{
		width: 3,
		height: 4,
	}

	fmt.Printf("%T\n", r)

	PrintShape(c)
	PrintShape(r)

	var circle Shape = &Circle{radius: 3}
	PrintShape(circle)

	fmt.Printf("%T\n", circle)

	// Type Assertions
	// How can we access to a property of Rectangle
	// from a type Shape? with type assertions
	fmt.Println(circle.(*Circle).radius)


	circle.IncreaseSize(3)

	fmt.Println(circle.(*Circle).radius)

}

func PrintShape (s Shape) {
	fmt.Printf("Area of %s is %.2f\n", s.String(), s.Area())
}
