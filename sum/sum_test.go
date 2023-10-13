package main

import "testing"

func TestSum(t *testing.T) {
	t.Run("Should sum two numbers", func(t *testing.T) {
		numbers := []int{1, 1, 1, 1, 2}
		got := Add(numbers)
		expected := 6

		if got != expected {
			t.Errorf("got %q but expected %q", got, expected)
		}

	})
}
