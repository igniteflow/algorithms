#include <stdio.h>
#include <stdlib.h>

/*
    Compile with: 

    cc insertion-sort.c -o insertion-sort
*/
int main(int argc, char **argv) 
{
   unsigned int a[20] = {16, 7, 4, 10, 18, 15, 6, 12, 13, 5, 11, 14, 17, 8, 2, 9, 19, 3, 1};
   unsigned int i, j, key;

   printf("Unsorted: [");
   for ( i = 0; i < 20; i++ ) {
       printf(" %d ", a[i]);
   }
   printf("]\n");

   for ( j = 1 ; j < 20 ; j++ ) 
   {
       key = a[j];
       i = j -1;
       while ( i > 0 && a[i] > key ) {
           a[i + 1] = a[i];
           i = i - 1;
       }    
       a[i + 1] = key;
   }

   printf("Sorted: [");
   for ( i = 0; i < 20; i++ ) {
       printf(" %d ", a[i]);
   }
   printf("]\n");

   print_to_terminal("hello", a);
}

