#include <iostream>
#include <cstdlib>  // rand(), srand()
#include <ctime>    // time()

int main() {
    std::srand(std::time(0)); // 現在の時間をシードとして乱数を初期化
    int target = std::rand() % 100 + 1; // 1から100の範囲の乱数を生成

    std::cout << "数当てゲームを開始します！" << std::endl;
    std::cout << "1から100の数を当ててください。" << std::endl;

    std::cout << "（デバッグ用：正解は " << target << "）" << std::endl; // デバッグ用

    int guess;
    int attempts = 0;

    while (true) {
        std::cout << "数を入力してください: ";
        std::cin >> guess;

        // 入力エラーをチェック
        if (!std::cin) {
            std::cout << "無効な入力です。整数を入力してください。" << std::endl;
            std::cin.clear();
            std::cin.ignore(10000, '\n');
            continue;
        }

        attempts++;

        if (guess < target) {
            std::cout << "もっと大きい数です！ " << attempts << "回目" << std::endl;
        } else if (guess > target) {
            std::cout << "もっと小さい数です！ " << attempts << "回目" << std::endl;
        } else {
            std::cout << "正解です！ " << attempts << "回で当たりました！" << std::endl;
            break;
        }
    }

    return 0;
}

