// Given a sorted array of integers, return the index of the given key. Return -1 if not found.
package binary_search

func binary_search_recursive(haystack [8]int, key int, low int, high int) int {
  // Runtime complexity: Logarithmic, O(logn)
  // Memory complexity: Logarithmic, O(logn)
  if low > high {
    return -1
  }

  mid := low + ((high - low) / 2)
  if haystack[mid] == key {
    return mid
  } else if key < haystack[mid] {
    return binary_search_recursive(haystack, key, low, mid - 1)
  } else {
    return binary_search_recursive(haystack, key, mid + 1, high)
  }
}

func binary_search_iterative(haystack [8]int, key int) int {
  // Runtime complexity: Logarithmic, O(logn)
  // Memory complexity: Constant, O(1)
  low, high := 0, len(haystack) - 1

  for low <= high {
    mid := low + ((high - low) / 2)

    if haystack[mid] == key {
      return mid
    }

    if key < haystack[mid] {
      high = mid - 1
    } else {
      low = mid + 1
    }
  }

  return -1
}
