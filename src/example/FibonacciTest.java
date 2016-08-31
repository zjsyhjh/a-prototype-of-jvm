public class FibonacciTest {

	public static void main(String[] args) {
		long result = fibonacci(20);
		System.out.println(result);
	}

	private static long fibonacci(int n) {
		if (n <= 1) {
			return n;
		} else {
			return fibonacci(n -1) + fibonacci(n - 2);
		}
	}
}
