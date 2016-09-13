import java.util.ArrayList;


public class BoxTest {

	public static void main(String[] args) {
		ArrayList<Integer> l = new ArrayList<>();
		l.add(new Integer(1));
		l.add(2);

		l.add(3);

		for (int num : l) {
			System.out.println(num);
		}
	}
}
