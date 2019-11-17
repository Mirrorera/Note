import java.util.Scanner;

public class DivideInt {
    public static void main(String[] args) {
        Scanner sc;
        sc = new Scanner(System.in);
        String s;
        s = sc.nextLine();
        int l = s.length();
        int res = l % 3;

        for(int i=0;i<l;++i) {
            System.out.printf("%c", s.charAt(i));
            if ((i + 4 - res) % 3 == 0 && i != l - 1) {
                System.out.printf(",");
            }
        }
        System.out.println();
    }
}