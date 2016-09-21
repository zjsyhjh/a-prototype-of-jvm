public class CloneTest implements Cloneable {

	public int value = 1;

	public CloneTest clone() {
		try {
			return (CloneTest)super.clone();
		} catch (java.lang.CloneNotSupportedException e) {
			System.out.println(e);
		}
		return null;
	}

	public static void main(String[] args) {

		CloneTest obj1 = new CloneTest();
		CloneTest obj2 = obj1.clone();

		System.out.println(obj1.value);

		obj1.value = 2;
		System.out.println(obj2.value);
	}
}