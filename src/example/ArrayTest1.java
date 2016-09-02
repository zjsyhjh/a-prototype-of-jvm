public class ArrayTest1 {

	public static void main(String[] args) {

		int intArray[] = {10, 1, 2, 6, 3, 9, 7, 8, 0, 5, 4};
		
		bubbleSort(intArray);

		for (int num : intArray) {
			System.out.println(num);
		}
	}

	private static void bubbleSort(int[] intArray) {

		int i = 1;
		boolean flag = true;
		while (flag) {
			flag = false;
			for (int j = 0;  j < intArray.length - i; j++) {
				if (intArray[j] > intArray[j + 1]) {
					int tmp = intArray[j + 1];
					intArray[j + 1] = intArray[j];
					intArray[j] = tmp;
					flag = true;
				}
			}
			i++;
		}
	}
}
