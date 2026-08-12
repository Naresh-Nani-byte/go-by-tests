package main

import (
	"slices"
	"testing"
)

func TestArrayAndSlices(t *testing.T){
	t.Run("let's do a sum from array", func(t *testing.T) {
		numbers := [5]int{2,3,4,5,6}
		got := ArrayAndSlices(numbers)
		want := 20
		assertMessage(t, got, want)
	})

	t.Run("let's have advanced method of range", func(t *testing.T) {
		numbers := [5]int{2,3,4,5,6}
		got := ArrayAndSlicesWithAdv(numbers)
		want := 20
		assertMessage(t, got, want)
	})
	t.Run("let's find out the average of array items", func(t *testing.T) {
		numbers := []int{2,3,4,5,6}
		got := CaulculateAverageOfArrayItems(numbers[1:4])
		want := 4
		assertMessage(t, got, want)
		// if slices.Equal(got, want){
		// 	t.Errorf("got %d but want %d")
		// }
	})

	t.Run("Sum of numbers in a slice", func(t *testing.T) {
		numbers := []int{2,3,4,5,6}
		got := CaulculateAverageOfArrayItems(numbers[1:4])
		want := 4
		assertMessage(t, got, want)
	})
}

func TestSumAll(t *testing.T){
	got := SumAll([]int{2, 3, 4}, []int{5, 6, 7}, []int{10, 9, 8})
	want := []int{9, 18, 27}
	if !slices.Equal(got,want){
		t.Errorf("got %v but want %v ", got, want)
	}
}

func TestSumAllTails(t *testing.T){
	
	checkSum := func (t testing.TB, got, want []int)  {
		if !slices.Equal(got, want){
			t.Errorf("got %v but want %v", got, want)
		}	
	}
	t.Run("make sum of tails of ", func(t *testing.T) {
		got := SumAllTails([]int{2, 3}, []int{5, 6}, []int{10, 9})
		want := []int{3,6,9}
		checkSum(t, got, want)
	})

	t.Run("include the zero in sum of tails of", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{1,3}, []int{0, 9})
		want := []int{0,3,9}
		checkSum(t, got, want)
	})
}

func assertMessage(t testing.TB, got, want int){
	t.Helper()
	if got != want {
		t.Errorf("got %d but want %d", got, want)
	}
}