#!/usr/bin/python

import random

# generate a shuffled list of numbers 0-19
a = range(20)
random.shuffle(a)
print 'Unsorted: %s' % a

def insertion_sort(a):
    for i in range(1, len(a)):
        key = a[i]
        for j in range(0, i):
            if(key < a[j]):
                z = a[j]
                a[j] = key
		key = z
        a[i] = key
    return a

# execute the sort
print 'Sorted: %s' % insertion_sort(a)
