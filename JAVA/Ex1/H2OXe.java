import java.util.Scanner;

public class H2OXe {
    public static void main(String args[]) {
        int N;
        Scanner sc;
        sc = new Scanner(System.in);
        N = sc.nextInt();
        for(int i=100;i<=N;++i) {
            int tot = 0;
            int op = i;
            while(op != 0) {
                int t = op % 10;
                op /= 10;
                tot += t * t * t;
            }
            if (tot == i) {
                System.out.println(i);
            }
        }
    }
}