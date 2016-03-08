package sliding_window

import "testing"


func TestSlidingWindow(t *testing.T) {
  data := []int{-3, 5, -2, 7, -1, 0, 4}
  window_size := 3
  actual := find_max_sliding_window(data, window_size)
  expected := 4

  if actual != expected {
    t.Errorf("Test failed, expected: '%d', got:  '%d'", expected, actual)
  }
}
