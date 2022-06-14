package model

// import "errors"

type User struct {
	ID     int
	UserID int
	Name   string
}

type Link struct {
	ID      int
	User1ID int
	User2ID int
}

// NewUser userのコンストラクタ
// user_idは重複しないようにランダムな数値にする
// nameもユニークにする
// func NewUser(name string) error {
// 	if name == "" {
// 		return errors.New("nameを入力してください")
// 	}

// 	// nameがユニークか確認する処理

// 	// ユニークなuser_idを作成する処理
// 	userID := 0

// 	user := &User{
// 		UserID: userID,
// 		Name: name,
// 	}
// 	println(user)
// 	// dbにuserを登録する処理

// 	return nil
// }


// セッターとかいろいろ追加する
