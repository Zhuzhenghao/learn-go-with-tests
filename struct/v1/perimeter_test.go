package main

import "testing"

func TestPerimeter(t *testing.T) {
	t.Run("perimeter of rectangle", func(t *testing.T) {
		got := Perimeter(Rectangle{10.0, 10.0})
		want := 40.0

		if got != want {
			t.Errorf("got %2.f want %2.f", got, want)
		}
	})

	checkArea := func(t *testing.T, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()

		if got != want {
			t.Errorf("got %2.f want %2.f", got, want)
		}
	}

	t.Run("area of rectangle", func(t *testing.T) {
		rectangle := Rectangle{12.0, 6.0}
		checkArea(t, rectangle, 72.0)
	})

	t.Run("area of circle", func(t *testing.T) {
		circle := Circle{10.0}
		checkArea(t, circle, 314.1592653589793)
	})
}
