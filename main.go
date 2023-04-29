package main

import (
	"context"
	"fmt"
	"log"
	"test/db"
	"test/ent/comment"
	"test/ent/user"

	"entgo.io/ent/dialect/sql"
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

func cleanUp() {
	client := db.NewDBClient()
	ctx := context.Background()

	_, err := client.Debug().User.Delete().Exec(ctx)
	if err != nil {
		fmt.Printf("failed deleting user: %v", err)
		return
	}

	db.CloseDB(client)
}

func addUserAndComment() {
	client := db.NewDBClient()
	ctx := context.Background()

	// 1件追加
	usr, err := client.Debug().User.
		Create().
		SetName("user2").
		SetAge(30).
		Save(ctx)
	if err != nil {
		fmt.Printf("failed creating user: %v", err)
		return
	}

	// コメント1件追加
	_, err = client.Debug().Comment.
		Create().
		SetUserID(usr.ID).
		SetComment("comment1").
		Save(ctx)
	if err != nil {
		fmt.Printf("failed creating comment: %v", err)
		return
	}

	// コメント1件追加
	_, err = client.Debug().Comment.
		Create().
		SetUserID(usr.ID).
		SetComment("comment2").
		Save(ctx)
	if err != nil {
		fmt.Printf("failed creating comment: %v", err)
		return
	}

	// user2のコメントを全件取得
	comments, err := usr.QueryComments().All(ctx)
	if err != nil {
		fmt.Printf("failed getting comments: %v", err)
		return
	}

	for _, comment := range comments {
		fmt.Println(comment.Comment)
	}

	// 'comment2'を持つユーザー一覧を取得
	usrs, err := client.Debug().User.Query().Where(func(s *sql.Selector) {
		t := sql.Table(comment.Table)
		s.Join(t).On(s.C(user.FieldID), t.C(comment.FieldUserID))
		s.Where(sql.EQ(t.C(comment.FieldComment), "comment2"))
	}).All(ctx)
	if err != nil {
		fmt.Printf("failed getting users: %v", err)
		return
	}

	for _, usr := range usrs {
		fmt.Println(usr.Name) // user2
	}

	db.CloseDB(client)
}

func main() {
	cleanUp()

	// crud()

	addUserAndComment()
}
