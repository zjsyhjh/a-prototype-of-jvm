public class InterpreterTest2 {

    public int instanceVar;
    public static int staticVar;
    public static void main(String[] args) {
        
        int x = 100000;
        InterpreterTest2 obj = new InterpreterTest2();
        obj.staticVar = x;
        x = obj.staticVar;
        obj.instanceVar = x;
        x = obj.instanceVar;
        Object other = obj;
        if (other instanceof InterpreterTest2) {
            obj = (InterpreterTest2) other;
            System.out.println(obj.instanceVar);
        }
    }
}
        
        
