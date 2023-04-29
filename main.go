package main

import (
	"context"
	"fmt"
	"log"
	"test/db"
	"test/ent/user"

	_ "github.com/lib/pq"
)

func crud() {
	client := db.NewDBClient()
	ctx := context.Background()

	// 1件追加
	usr, err := client.Debug().User.
		Create().
		SetName("user1").
		SetAge(33).
		Save(ctx)
	if err != nil {
		fmt.Printf("failed creating user: %v", err)
		return
	}

	// 1件更新
	updatedUser, err := client.Debug().User.Update().Where(user.ID(usr.ID)).SetAge(29).Save(ctx)
	if err != nil {
		fmt.Printf("failed updating user: %v", err)
		return
	}
	log.Printf("user: %+v", updatedUser)

	// 名前がuser1のユーザーを取得
	users, err := client.Debug().User.Query().Where(user.Name("user1")).All(ctx)
	if err != nil {
		fmt.Printf("failed getting users: %v", err)
		return
	}

	for _, usr := range users {
		fmt.Printf("user: %+v", usr)
	}

	// 1件削除
	_, err = client.Debug().User.Delete().Where(user.Name("user1")).Exec(ctx)
	if err != nil {
		fmt.Printf("failed deleting user: %v", err)
		return
	}

	// DB接続を閉じる
	db.CloseDB(client)
}

func main() {
	crud()
}
