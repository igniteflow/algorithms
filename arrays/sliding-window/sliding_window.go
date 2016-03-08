package sliding_window

import "container/list"


// uses a linked list to find the max value in a list
// window shifting not yet implemented
func find_max_sliding_window(haystack []int, window_size int) int {
  if window_size > len(haystack) {
    return -1
  }

  // instantiate a linked list
  var window list.List

  // populate first window
  for i := 0; i < window_size; i++ {
    tail := window.Back()

    // remove any tail elements that are less than the next value
    if tail != nil && haystack[i] >= tail.Value.(int) {
      window.Remove(tail)
    }
    window.PushBack(haystack[i])
  }

  for i := window_size; i < len(haystack); i++ {
    // push tail element off the heap if it's less than the current haystack value
    tail := window.Back()
    if haystack[i] > tail.Value.(int) {
      window.Remove(tail)
    }

    // remove head if it is no longer in the window
    head := window.Front()

    if head != nil {
      for index, needle := range haystack {
          if needle == head.Value.(int) {
              if index < i {
                  window.Remove(head)
              }
          }
      }
    }

    // add the current element to the tail
    window.PushBack(haystack[i])
  }

  return window.Front().Value.(int)
}
