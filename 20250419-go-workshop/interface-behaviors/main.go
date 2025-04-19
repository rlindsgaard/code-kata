package main

import (
	"fmt"
	"math"
)

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.Radius, 2)
}

func (c Circle) String() string {
	return fmt.Sprintf("Circle {Radius: %.2f}", c.Radius)
}

type Square struct {
	Side float64
}

func (s Square) Area() float64 {
	return s.Side * s.Side
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) String() string {
	return fmt.Sprintf("Rectangle {Width: %.2f, Height: %.2f}", r.Width, r.Height)
}

type Sizer interface {
	Area() float64
}

type Shaper interface {
	Sizer
	fmt.Stringer
}

func main() {
	c := Circle{Radius: 10}
	PrintArea(c)
	// s := Square{Side: 4}
	r := Rectangle{Width: 5, Height: 10}
	PrintArea(r)

	l := Less(c, r)
	fmt.Printf("%+v is the smallest\n", l)
}

func Less(a, b Sizer) Sizer {
	if a.Area() < b.Area() {
		return a
	}
	return b
}

func PrintArea(s Shaper) {
	fmt.Printf("The area of %s is %.2f\n", s.String(), s.Area())
}
