package main

import "testing"

func TestSum(t *testing.T) {
	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		sum := Sum(numbers)
		want := 6

		if sum != want {
			t.Errorf("got '%d' want '%d' given, '%v'", sum, want, numbers)
		}
	})

}
