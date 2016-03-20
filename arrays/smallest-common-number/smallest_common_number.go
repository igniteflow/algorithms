package smallest_common_number

// Given three integer arrays sorted in ascending order, find the smallest
// number that is common in all three arrays

func findSmallestCommonNumber(arr1, arr2, arr3 []int) int {
    // initiate the array index counters to zero
    i, j, k := 0, 0, 0

    for i < len(arr1) && j < len(arr2) && k < len(arr3) {
        if arr1[i] == arr2[j] && arr2[j] == arr3[k] {
            // all values are the same, so return
            return arr1[i]
        }

        // as the arrays are sorted, just increment the smallest value's
        // index for the next comparison check
        if arr1[i] <= arr2[j] && arr1[i] <= arr3[k] {
            i += 1
        } else if arr2[j] <= arr1[i] && arr2[j] <= arr3[k] {
            j += 1
        } else if arr3[k] <= arr1[i] && arr3[k] <= arr2[j] {
            k += 1
        }
    }

    // number common to all arrays not found
    return -1
}
