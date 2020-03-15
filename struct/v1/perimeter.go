package main

import "math"

// Shape ...
type Shape interface {
	Area() float64
}

// Rectangle ...
type Rectangle struct {
	Width  float64
	Height float64
}

// Area ...
func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

// Circle ...
type Circle struct {
	Radius float64
}

// Area ...
func (c Circle) Area() float64 {
	return c.Radius * c.Radius * math.Pi
}

// Perimeter ...
func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

// Area ...
func Area(rectangle Rectangle) float64 {
	return rectangle.Width * rectangle.Height
}
