package rotate_array_search

import "testing"


func TestBinarySearchRotated(t *testing.T) {
  array := []int{4,5,6,7,0,1,2}
  key := 6
  actual := binarySearchRotated(array, key)
  expected := 2

  if actual != expected {
    t.Errorf("Test failed, expected: '%d', got:  '%d'", expected, actual)
  }
}
