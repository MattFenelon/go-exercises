package binarysearch

import (
	"testing"
)

func Test_recursive_when_element_is_not_in_slice(t *testing.T) {
	a := []int{0, 1, 2, 3}

	actualIndex := RecursiveBinarySearch(a, 4)

	t.Log("It should return -1")
	if actualIndex != -1 {
		t.Errorf("Expected -1 but was %v", actualIndex)
	}
}

func Test_recursive_when_slice_is_empty(t *testing.T) {
	a := []int{}

	actualIndex := RecursiveBinarySearch(a, 1)

	t.Log("It should return -1")
	if actualIndex != -1 {
		t.Errorf("Expected -1 but was %v", actualIndex)
	}
}

func Test_recursive_when_slice_length_is_even_and_n_is_in_the_middle(t *testing.T) {
	a := []int{0, 1, 2, 3}

	actualIndex := RecursiveBinarySearch(a, 2)

	t.Log("It should return the index of n")
	if actualIndex != 2 {
		t.Errorf("Expected 2 but was %v", actualIndex)
	}
}

func Test_recursive_when_slice_length_is_even_and_n_is_at_the_beginning(t *testing.T) {
	a := []int{0, 1, 2, 3}

	actualIndex := RecursiveBinarySearch(a, 0)

	t.Log("It should return the index of n")
	if actualIndex != 0 {
		t.Errorf("Expected 0 but was %v", actualIndex)
	}
}

func Test_recursive_when_slice_length_is_even_and_n_is_at_the_end(t *testing.T) {
	a := []int{0, 1, 2, 3}

	actualIndex := RecursiveBinarySearch(a, 3)

	t.Log("It should return the index of n")
	if actualIndex != 3 {
		t.Errorf("Expected 3 but was %v", actualIndex)
	}
}

func Test_recursive_when_slice_length_is_odd_and_n_is_right_of_the_middle(t *testing.T) {
	a := []int{0, 1, 2, 3, 4}

	actualIndex := RecursiveBinarySearch(a, 3)

	t.Log("It should return the index of n")
	if actualIndex != 3 {
		t.Errorf("Expected 3 but was %v", actualIndex)
	}
}

func Test_recursive_when_slice_length_is_odd_and_n_is_left_of_the_middle(t *testing.T) {
	a := []int{0, 1, 2, 3, 4}

	actualIndex := RecursiveBinarySearch(a, 1)

	t.Log("It should return the index of n")
	if actualIndex != 1 {
		t.Errorf("Expected 1 but was %v", actualIndex)
	}
}

func Test_recursive_when_slice_is_not_in_sequence_and_n_is_a_missing_element(t *testing.T) {
	a := []int{0, 1, 3, 4}

	actualIndex := RecursiveBinarySearch(a, 2)

	t.Log("It should return -1")
	if actualIndex != -1 {
		t.Errorf("Expected -1 but was %v", actualIndex)
	}
}
