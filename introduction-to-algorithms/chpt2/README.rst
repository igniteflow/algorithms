Performance results for insertion sort:

C
--

time ./insertion-sort::

  Unsorted: [ 16  7  4  10  18  15  6  12  13  5  11  14  17  8  2  9  20  19  3  1 ]
  Sorted: [ 1  2  3  4  5  6  7  8  9  10  11  12  13  14  15  16  17  18  19  20 ]
  
  real  0m0.003s
  user	0m0.000s
  sys	0m0.000s


Java
----

time java InsertionSort:: 
 
    [16, 7, 4, 10, 18, 15, 6, 12, 13, 5, 11, 14, 17, 8, 2, 9, 19, 3, 1]
    [1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19]
    real  0m0.104s
    user  0m0.064s
    sys	  0m0.028s                                                                       


Python
------

time ./insertion-sort.py::

  Unsorted: [16, 7, 4, 10, 18, 15, 6, 12, 13, 5, 11, 14, 17, 8, 2, 9, 19, 3, 1]
  Sorted: [1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19]
  real	0m0.034s
  user	0m0.028s
  sys	0m0.008s
