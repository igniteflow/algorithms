#!/usr/bin/python

a = [16, 7, 4, 10, 18, 15, 6, 12, 13, 5, 11, 14, 17, 8, 2, 9, 19, 3, 1]
print 'Unsorted: %s' % a

def insertion_sort(a):
    for j in range(1, len(a)):
        key = a[j]
        i = j - 1
        while i >= 0 and a[i] > key:
	    a[i+1] = a[i]
	    i = i - 1
 	a[i+1] = key
    return a

# execute the sort
print 'Sorted: %s' % insertion_sort(a)
