package binary_search

import "testing"


func TestBinarySearchRecursive(t *testing.T) {
  haystack := [8]int{1, 10, 20, 47, 59, 63, 75, 88}
  key := 47
  low := 0
  high := len(haystack) - 1
  actual := binary_search_recursive(haystack, key, low, high)
  expected := 3

  if actual != expected {
    t.Errorf("Test failed, expected: '%d', got:  '%d'", expected, actual)
  }
}

func TestBinarySearchIterative(t *testing.T) {
  haystack := [8]int{1, 10, 20, 47, 59, 63, 75, 88}
  key := 47
  actual := binary_search_iterative(haystack, key)
  expected := 3

  if actual != expected {
    t.Errorf("Test failed, expected: '%d', got:  '%d'", expected, actual)
  }
}
