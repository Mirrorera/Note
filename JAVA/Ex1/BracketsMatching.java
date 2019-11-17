import java.util.Scanner;
import java.io.*;

public class BracketsMatching {
    public static void main(String[] args) throws FileNotFoundException{
        Scanner sc;
        sc = new Scanner(new File("correct.in"));

        char[] stack;
        stack = new char[80 + 5];
        int top = 0;

        String s;
        s = sc.nextLine();
        int l = s.length();

        System.out.println(s);
        boolean flag = true;
        for(int i=0;i<l;++i) {
            
            switch (s.charAt(i)) {
                case '(':
                    stack[top] = s.charAt(i);
                    top++;
                    break;
                case '{':
                    stack[top] = s.charAt(i);
                    top++;
                    break;
                case '[':
                    stack[top] = s.charAt(i);
                    top++;
                    break;
                case ')':
                    if (top == 0 || stack[top - 1] != '(') {
                        flag = false;
                    }
                    top--;
                    break;
                case '}':
                    if (top == 0 || stack[top - 1] != '{') {
                        flag = false;
                    }
                    top--;
                    break;
                case ']':
                    if (top == 0 || stack[top - 1] != '[') {
                        flag = false;
                    }
                    top--;
                    break;
                default:
                    break;
            }
            if (!flag) {
                break;
            }
        }

        File ofile = new File("correct.out");
        PrintStream ofstream = new PrintStream(new FileOutputStream(ofile));
        System.setOut(ofstream);

        if(flag && top == 0) {
            System.out.println("True");
        } else {
            System.out.println("False");
        }

    }

}