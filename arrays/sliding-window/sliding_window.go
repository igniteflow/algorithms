package main

import "fmt"
import "container/list"


// uses a linked list to find the max value in a list
// window shifting not yet implemented
func find_max_sliding_window(haystack []int, window_size int) {
  var window list.List

  for j := 0; j < len(haystack); j++ {
    tail := window.Back()

    if tail == nil {
      window.PushBack(haystack[j])
    } else {
      if haystack[j] >= tail.Value.(int) {
        window.Remove(tail)
        window.PushBack(haystack[j])
      }
    }
  }

  for e := window.Front(); e != nil; e=e.Next() {
    fmt.Println(e.Value.(int))
  }

}


func main() {
  data := []int{-3, 5, -2, 7, -1, 0, 4}
  window_size := 3
  find_max_sliding_window(data, window_size)
}
