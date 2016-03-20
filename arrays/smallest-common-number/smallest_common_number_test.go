package smallest_common_number

import "testing"


func TestSmallestCommonNumber(t *testing.T) {
  arr1 := []int{6,7,8,10,25,30,63,64}
  arr2 := []int{-1,4,5,6,8,50}
  arr3 := []int{8,10,14}

  actual := findSmallestCommonNumber(arr1, arr2, arr3)
  expected := 8

  if actual != expected {
    t.Errorf("Test failed, expected: '%d', got:  '%d'", expected, actual)
  }
}
