package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// GitHubのユーザー情報をマッピングする構造体
type GitHubUser struct {
	Login       string  `json:"login"`
	Name        string  `json:"name"`
	ID          int     `json:"id"`
	UserType    string  `json:"type"`
	UserView    string  `json:"user_view_type"`
	AvatarURL   string  `json:"avatar_url"`
	ProfileURL  string  `json:"html_url"`
	PublicRepos *int    `json:"public_repos"`
	Followers   *int    `json:"followers"`    
	Following   *int    `json:"following"`
	Blog        *string  `json:"blog"`
	Company     *string `json:"company"`
	Email       *string `json:"email"`
	Bio         *string `json:"bio"`
	Created     string  `json:"created_at"`
	Update      string  `json:"updated_at"`
}

func main() {
	fmt.Println("GitHubのユーザー名を入力してください。")
	var username string
	fmt.Scan(&username)

	url := "https://api.github.com/users/" + username
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	// JSONを構造体に変換
	var user GitHubUser
	err = json.Unmarshal(body, &user)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	// フォーマットして表示
	fmt.Println("====================================================")
	fmt.Printf("GitHubユーザー: %s\n", user.Login)
	fmt.Printf("GitHubニックネーム %s\n", user.Name)
	fmt.Printf("ID: %d\n", user.ID)
	fmt.Printf("ユーザータイプ: %s\n",user.UserType)
	fmt.Printf("ユーザー公開情報: %s\n" ,user.UserView)
	fmt.Printf("アバター画像: %s\n", user.AvatarURL)
	fmt.Printf("プロフィール: %s\n", user.ProfileURL)
	fmt.Printf("公開リポジトリ数: %s\n", formatInt(user.PublicRepos))
	fmt.Printf("フォロワー数: %s\n", formatInt(user.Followers))
	fmt.Printf("フォロー中: %s\n", formatInt(user.Following))
	fmt.Printf("blog: %s\n", formatString(user.Blog))
	fmt.Printf("会社名: %s\n" , formatString(user.Company))
	fmt.Printf("E-mail: %s\n" , formatString(user.Email))
	fmt.Printf("Bio: %s\n" , formatString(user.Bio))
	fmt.Printf("アカウント作成日: %s\n",user.Created)
	fmt.Printf("更新日: %s\n",user.Update)
	fmt.Println("====================================================")

	end()
}

// *int を安全に表示する関数
func formatInt(value *int) string {
	if value != nil {
		return fmt.Sprintf("%d", *value) // 値がある場合はそのまま
	}
	return "データなし" // `nil` の場合
}

// *string を安全に表示する関数
func formatString(value *string) string {
	if value != nil && *value != "" {
		return *value // 値がある場合はそのまま返す
	}
	return "null" // `nil` または空文字の場合
}

func end() {
	var next string
			fmt.Println("終了してもよろしいですか？(y)\n")
			fmt.Scan(&next)
			if next == "Yes" || next == "YES" || next == "yes" || next == "y" || next == "Y" {
				return
			}else {
				fmt.Println("無効な数値です。もう一度お願いします。")
				return
			}
}
