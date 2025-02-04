import java.util.Random;
import java.util.Scanner;

public class random {
    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);
        Random random = new Random();
        
        // ユーザーに2～5の範囲で選択させる
        int maxRange;
        do {
            System.out.print("2~5の範囲で最大値を選択してください(2~5): ");
            while (!scanner.hasNextInt()) {
                System.out.println("数値を入力してください。");
                scanner.next(); // 無効な入力を読み飛ばす
            }
            maxRange = scanner.nextInt();
        } while (maxRange < 2 || maxRange > 5);
        
        // 1～maxRangeの範囲でランダムな数値を選択
        int selectedNumber = random.nextInt(maxRange) + 1;
        
        // 結果を出力
        System.out.println("選ばれた数値: " + selectedNumber);
        
        scanner.close();
    }
}
