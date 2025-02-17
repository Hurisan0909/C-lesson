package main

import (
    "fmt"
    "math/rand"
    "strings"
    "time"
)

func main() {
    // 乱数のシードを設定
    rand.Seed(time.Now().UnixNano())
    
    for {
        playGame()
        
        // もう一度プレイするか確認
        var answer string
        fmt.Print("\nもう一度プレイしますか？ (はい/いいえ): ")
        fmt.Scan(&answer)
        
        // 小文字に変換して比較
        answer = strings.ToLower(answer)
        if answer == "はい" || answer == "y" || answer == "yes" {
            fmt.Println("\n新しいゲームを開始します！")
        } else if answer == "いいえ" || answer == "n" || answer == "no" {
            survey()
            fmt.Println("ゲームを終了します。遊んでくれてありがとう！")
            break
        }else {
            fmt.Println("無効な入力です。ゲームを終了します。")
            break
        }
    }
}

func playGame() {
    // 1から100までのランダムな数を生成
    target := rand.Intn(100) + 1
    
    fmt.Println("1から100までの数を当ててください！")
    //fmt.Println("Debug: ", target)
    
    // ゲームのメインループ
    for attempts := 0; ; attempts++ {
        var guess int
        fmt.Print("あなたの予想は？: ")
        fmt.Scan(&guess)
        
        // 入力された数字を判定
        if guess < target {
            fmt.Println("もっと大きい数字です！")
        } else if guess > target {
            fmt.Println("もっと小さい数字です！")
        } else {
            fmt.Printf("正解です！%d回で当てることができました！\n", attempts+1)
            break
        }
    }
}

func survey() {
    fmt.Println("アンケートにご協力ください！")
    fmt.Println("1. とても良い")
    fmt.Println("2. 良い")
    fmt.Println("3. 普通")
    fmt.Println("4. 悪い")
    fmt.Println("5. とても悪い")

    var choice int
    fmt.Print("あなたの評価は？: ")
    fmt.Scan(&choice)

    switch choice {
    case 1:
        fmt.Println("とても良いと評価していただき、ありがとうございます！")
        break
    case 2:
        fmt.Println("良いと評価していただき、ありがとうございます！")
        break
    case 3:
        fmt.Println("普通ですか...もっと良くするために頑張ります！")
        break
    case 4:
        fmt.Println("悪いですか...申し訳ございません。")
        break
    case 5:
        fmt.Println("ええぇ...どこが悪いんですか...?勘弁してください :(")
        break
    default:
        fmt.Println("無効な入力です。")
    }
}