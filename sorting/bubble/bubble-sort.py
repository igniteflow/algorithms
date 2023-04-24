from typing import List


def bubble(x: List[int]) -> List[int]:
    """
    Pairs of adjacent elements are swapped if they are not in order
    Quadratic time 0(n^2)
    """
    reordered = False
    for i in range(len(x)):
        if i == len(x) - 1:
            break
        first = x[i]
        second = x[i + 1]
        if first > second:
            second_val = x.pop(i + 1)
            x.insert(i, second_val)
            reordered = True
    if reordered:
        # recurse
        bubble(x)
    return x


if __name__ == "__main__":
    x = [2, 4, 7, 6, 5, 9, 1, 8, 3]
    print(bubble(x))
