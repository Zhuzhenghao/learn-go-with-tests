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

// Triangle ...
type Triangle struct {
	Height float64
	Base   float64
}

// Area ..
func (t Triangle) Area() float64 {
	return t.Height * t.Base / 2
}
