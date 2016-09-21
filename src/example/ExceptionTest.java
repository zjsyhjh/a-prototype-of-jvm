public class ExceptionTest {

	public static void main(String[] args) {

		method1(args);
	}

	private static void method1(String[] args) {
		try {
			method2(args);
		} catch(java.lang.NumberFormatException e) {
			System.out.println(e.getMessage());
		}
	}

	private static void method2(String[] args) {
		if (args.length == 0) {
			throw new IndexOutOfBoundsException("No args!");
		}
		System.out.println(Integer.parseInt(args[0]));
	}
}