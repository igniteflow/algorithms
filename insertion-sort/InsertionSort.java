import java.util.Arrays;

public class InsertionSort {

	public int a[] = new int[] { 16, 7, 4, 10, 18, 15, 6, 12, 13, 5, 11, 14,
			17, 8, 2, 9, 19, 3, 1 };
	public int i, j, key;

	public void insertionSort() {
		System.out.println(Arrays.toString(a));
		for (j = 0; j < 19; j++) {
			key = a[j];
			i = j - 1;
			while (i >= 0 && a[i] > key) {
				a[i + 1] = a[i];
				i = i - 1;
			}
			a[i + 1] = key;
		}
		System.out.println(Arrays.toString(a));
	}

	public static void main(String[] args) {
		InsertionSort insertion = new InsertionSort();
		insertion.insertionSort();
	}
}
