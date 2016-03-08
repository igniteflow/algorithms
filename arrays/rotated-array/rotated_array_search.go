package rotate_array_search

import "fmt"

func binarySearch(array []int, start, end, key int) int {
  if start > end {
    return -1
  }

  // prevent overflow
  mid := start + (end - start) / 2

  if array[mid] == key {
    return mid
  }

  if key > array[mid] {
    return binarySearch(array, (mid + 1), end, key)
  }
  return binarySearch(array, start, (mid - 1), key)
}

func binarySearchRotated(array []int, key int) int {
  return binarySearch(array, 0, len(array) - 1, key)
}

func main() {
  array := []int{4,5,6,7,0,1,2}
  key := 6
  res := binarySearchRotated(array, key)
  fmt.Println(res)
}
