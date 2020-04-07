package arrays

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	assertSum := func(t *testing.T, got, want int, numbers []int) {
		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	}

	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		got := Sum(numbers)
		want := 15
		assertSum(t, got, want, numbers)
	})

	t.Run("collection of 3 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		got := Sum(numbers)
		want := 6
		assertSum(t, got, want, numbers)
	})
}

func TestSumAll(t *testing.T) {
	assertSum := func(t *testing.T, got, want []int) {
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %d want %d given", got, want)
		}
	}

	t.Run("sum all integer arrays", func(t *testing.T) {
		got := SumAll([]int{1, 2}, []int{0, 9})
		want := []int{3, 9}
		assertSum(t, got, want)
	})
}

func TestSumAllTails(t *testing.T) {
	assertSum := func(t *testing.T, got, want []int) {
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %d want %d given", got, want)
		}
	}

	t.Run("sum all tails integer arrays", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{3, 4, 5})
		want := []int{2, 9}
		assertSum(t, got, want)
	})

	t.Run("sum all tails w/ empty array", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}
		assertSum(t, got, want)
	})

	t.Run("sum all tails w/ lots of array", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5, 6, 7, 8}, []int{10})
		want := []int{0, 30, 0}
		assertSum(t, got, want)
	})
}
