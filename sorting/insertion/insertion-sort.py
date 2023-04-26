from typing import List

"""
O(n²) in the average and worst case, and O(n) in the best case.
"""


def insertion_sort(x: List[int]):
    """
    For each element in a list, walk backwards until a smaller value is found,
    then insert the element after that value
    """
    # start at the second element as position 0 is already considered sorted
    for i_idx in range(1, len(x)):
        # position behind key
        j_idx = i_idx - 1

        # i starts on second list element
        key = x[i_idx]

        while j_idx >= 0 and x[j_idx] > key:
            # move j to the right as it is greater than i
            x[j_idx + 1] = x[j_idx]

            # walk backwards until we find an element
            # smaller than key
            j_idx -= 1

        # j is smaller than key so insert key after j
        x[j_idx + 1] = key
    return x


if __name__ == "__main__":
    a = [2, 4, 7, 6, 5, 9, 1, 8, 3]
    print("Before:", a)
    print("After:", insertion_sort(a))
