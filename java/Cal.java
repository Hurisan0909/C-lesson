import java.util.Scanner;

public class Cal {
    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);

        System.out.print("1つ目の数字を入力してください: ");
        double num1 = scanner.nextDouble();


        System.out.print("2つ目の数字を入力してください:");
        double num2 = scanner.nextDouble();

        System.out.println("足し算: " + (num1 + num2));
        System.out.println("引き算: " + (num1 - num2));
        System.out.println("掛け算: " + (num1 * num2));

        if (num2 != 0) {
            System.out.println("割り算: " + (num1 / num2));
        } else{
            System.out.println("割り算: 0では割れません");
        }

        scanner.close();
    }
}