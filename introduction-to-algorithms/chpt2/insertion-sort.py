#!/usr/bin/python

import random

# generate a shuffled list of numbers 0-19
a = range(1, 21)
random.shuffle(a)
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
