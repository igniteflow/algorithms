#!/usr/bin/python
"""
Merge sort
- divide-and-conquer
- recursive
- good choice for sorting a linked list
"""
m = [16, 7, 4, 10, 18, 15, 6, 12, 13, 5, 11, 14, 17, 8, 2, 9, 19, 3, 1]


def merge_sort(m):
    if len(m) <= 1:
        return m
    middle = len(m) / 2
    left, right = m[:middle], m[middle:]
    left = merge_sort(left)
    right = merge_sort(right)
    return merge(left, right)

def merge(left, right):
    result = []
    while len(left) > 0 or len(right) > 0:
        if len(left) > 0 and len(right) > 0:
            if left[0] <= right[0]:
                result.append(left.pop(0))
            else:
                result.append(right.pop(0))
        elif len(left) > 0:
            result.append(left.pop(0))
        elif len(right) > 0:
            result.append(right.pop(0))
    return result

print merge_sort(m)
