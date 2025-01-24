#include <stdio.h>
#include <stdlib.h>
#include <string.h>

void show_menu() {
    printf("==============================\n");
    printf("   ターミナル操作プログラム    \n");
    printf("==============================\n");
    printf("1.メッセージを表示\n");
    printf("2.簡単な計算(足し算)\n");
    printf("3.終了\n");
    printf("選択肢を選択してください: ");
}

void display_message() {
    printf("こんにちは!これはCで作られたプログラムです。 \n");
}

void perform_calculation() {
    double num1,num2;
    printf("最初の数値を入力してください: ");
    scanf("%lf",&num1);
    printf("次の数値を入力してください: ");
    scanf("%lf",&num2);
    printf("結果: %.2lf + %.2lf = %.2lf\n",num1,num2,num1 + num2);
}

int main() {
    int choice;

    while (1) {
        show_menu();
        scanf("%d", &choice);

        switch (choice) {
            case 1:
                display_message();
                break;
            case 2:
                perform_calculation();
                break;
            case 3:
                printf("プログラムを終了します。 \n");
                exit(0);
            default:
                printf("無効な選択です。もう一度お試しください。 \n");
        }

        printf("\n続けるにはEnterキーを押してください...\n");
        getchar();
        getchar();
    }

    return 0;
}