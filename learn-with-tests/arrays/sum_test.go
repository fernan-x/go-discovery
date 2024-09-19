package main

import (
	"reflect"
	"slices"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	result := SumAll([]int{1, 2}, []int{0, 9})
	expected := []int{3, 9}

	if !slices.Equal(result, expected) {
		t.Errorf("got %v want %v", result, expected)
	}
}

func TestSumAllTails(t *testing.T) {
	t.Run("slices with 2 elements", func(t *testing.T) {
		result := SumAllTails([]int{1, 2}, []int{0, 9})
		expected := []int{2, 9}

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("got %v want %v", result, expected)
		}
	})

	t.Run("slices with 3 elements", func(t *testing.T) {
		result := SumAllTails([]int{1, 2, 3}, []int{0, 9, 8})
		expected := []int{5, 17}

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("got %v want %v", result, expected)
		}
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}
