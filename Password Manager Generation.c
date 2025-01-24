#include <stdio.h>
#include <stdlib.h>
#include <time.h>

#define PASSWORD_LENGTH 12  // デフォルトのパスワード長

// パスワードに使用する文字セット
const char charset[] = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
                       "abcdefghijklmnopqrstuvwxyz"
                       "0123456789"
                       "!@#$%^&*()";

void generate_password(char *password, int length) {
    int charset_size = sizeof(charset) - 1;  // 文字セットのサイズを計算
    for (int i = 0; i < length; i++) {
        password[i] = charset[rand() % charset_size];  // ランダムな文字を選択
    }
    password[length] = '\0';  // 文字列の終端を設定
}

int main() {
    srand((unsigned int)time(NULL));  // ランダムシードを設定

    int length = PASSWORD_LENGTH;  // パスワード長を設定
    char password[PASSWORD_LENGTH + 1];  // パスワード格納用バッファ

    // ユーザーがパスワード長を指定するオプション
    printf("パスワードの長さを入力してください（デフォルト: %d）: ", PASSWORD_LENGTH);
    if (scanf("%d", &length) != 1 || length <= 0) {
        printf("無効な入力です。デフォルトの長さ %d で生成します。\n", PASSWORD_LENGTH);
        length = PASSWORD_LENGTH;
    }

    generate_password(password, length);

    printf("生成されたパスワード: %s\n", password);

    return 0;
}
