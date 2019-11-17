import java.util.Scanner;

public class ReverseString {
    public static void main(String[] args) {
        String s;
        Scanner sc;
        sc = new Scanner(System.in);
        s = sc.nextLine();
        int l = s.length();
        for(int i=l - 1;i>=0;--i) {
            System.out.printf("%c", s.charAt(i));
        }
        System.out.println();
    }
}