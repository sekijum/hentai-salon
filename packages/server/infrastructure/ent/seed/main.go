package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"

	"server/infrastructure/ent"

	"github.com/google/uuid"
)

func main() {
	client, cleanup, err := ent.ProvideClient()
	if err != nil {
		log.Fatalf("failed to initialize database client: %v", err)
	}
	defer cleanup()

	ctx := context.Background()
	if err := seed(ctx, client); err != nil {
		log.Fatalf("failed seeding data: %v", err)
	}
}

func randomImageURL(index int) string {
	width := rand.Intn(500) + 200  // 200から700のランダムな幅
	height := rand.Intn(500) + 200 // 200から700のランダムな高さ
	url := fmt.Sprintf("https://picsum.photos/seed/%d/%d/%d.webp", index, width, height)
	return url
}

func videoURL() string {
	return "https://www.w3schools.com/html/mov_bbb.mp4"
}

func seed(ctx context.Context, client *ent.Client) error {
	rand.Seed(time.Now().UnixNano())

	tx, err := client.Tx(ctx)
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			if rerr := tx.Rollback(); rerr != nil {
				log.Fatalf("tx rollback failed: %v", rerr)
			}
			return
		}
		if err := tx.Commit(); err != nil {
			log.Fatalf("tx commit failed: %v", err)
		}
	}()

	// Create Users
	users := make([]*ent.User, 100)
	for i := 0; i < 90; i++ {
		name := "ユーザー" + uuid.New().String()[:8]
		avatarURL := randomImageURL(i)
		createUser := tx.User.Create().
			SetName(name).
			SetEmail("user" + uuid.New().String()[:8] + "@example.com").
			SetPassword("pass1234").
			SetCreatedAt(time.Now()).
			SetUpdatedAt(time.Now())
		if rand.Intn(1000) >= 100 { // 1000件ほどnullにする
			createUser.SetAvatarURL(avatarURL)
		}
		users[i] = createUser.SaveX(ctx)
	}

	// Create Admin Users
	for i := 90; i < 100; i++ {
		name := "管理者" + uuid.New().String()[:8]
		avatarURL := randomImageURL(i)
		createAdmin := tx.User.Create().
			SetName(name).
			SetEmail("admin" + uuid.New().String()[:8] + "@example.com").
			SetPassword("pass1234").
			SetRole(1). // 管理者の役割を設定
			SetCreatedAt(time.Now()).
			SetUpdatedAt(time.Now())
		if rand.Intn(1000) >= 100 { // 1000件ほどnullにする
			createAdmin.SetAvatarURL(avatarURL)
		}
		users[i] = createAdmin.SaveX(ctx)
	}

	// Create Boards
	boards := make([]*ent.Board, 100)
	for i := 0; i < 100; i++ {
		thumbnailURL := randomImageURL(i)
		createBoard := tx.Board.Create().
			SetTitle("掲示板" + uuid.New().String()[:8]).
			SetDescription("掲示板の説明" + uuid.New().String()[:8]).
			SetUserID(users[rand.Intn(100)].ID).
			SetCreatedAt(time.Now()).
			SetUpdatedAt(time.Now())
		if rand.Intn(1000) >= 100 { // 1000件ほどnullにする
			createBoard.SetThumbnailURL(thumbnailURL)
		}
		boards[i] = createBoard.SaveX(ctx)
	}

	// Create Threads
	threads := make([]*ent.Thread, 100)
	for i := 0; i < 100; i++ {
		thumbnailURL := randomImageURL(i)
		createThread := tx.Thread.Create().
			SetTitle("スレッド" + uuid.New().String()[:8]).
			SetDescription("スレッドの説明" + uuid.New().String()[:8]).
			SetIPAddress("127.0.0.1").
			SetBoardID(boards[rand.Intn(100)].ID).
			SetUserID(users[rand.Intn(100)].ID).
			SetCreatedAt(time.Now()).
			SetUpdatedAt(time.Now())
		if rand.Intn(1000) >= 100 { // 1000件ほどnullにする
			createThread.SetThumbnailURL(thumbnailURL)
		}
		threads[i] = createThread.SaveX(ctx)
	}

	// Create Thread Comments with nested parent_comment_id
	comments := make([]*ent.ThreadComment, 100000)
	for i := 0; i < 100000; i++ {
		var parentID *int
		if i >= 1000 {
			pID := comments[rand.Intn(1000)].ID
			parentID = &pID
		}
		createComment := tx.ThreadComment.Create().
			SetContent("コメント内容" + uuid.New().String()[:8]).
			SetIPAddress("127.0.0.1").
			SetThreadID(threads[rand.Intn(100)].ID).
			SetUserID(users[rand.Intn(100)].ID).
			SetCreatedAt(time.Now()).
			SetUpdatedAt(time.Now())
		if parentID != nil {
			createComment.SetParentCommentID(*parentID)
		}
		comments[i] = createComment.SaveX(ctx)
	}

	// Create Thread Comment Attachments
	for i := 0; i < 300000; i++ {
		var attachmentURL string
		if i%2 == 0 {
			attachmentURL = randomImageURL(i)
		} else {
			attachmentURL = videoURL()
		}
		tx.ThreadCommentAttachment.Create().
			SetURL(attachmentURL).
			SetDisplayOrder(int(rand.Int63n(100))).
			SetType(i % 2). // typeを0か1に設定
			SetCreatedAt(time.Now()).
			SetCommentID(comments[rand.Intn(100000)].ID).
			SaveX(ctx)
	}

	// Create Tags
	tags := make([]*ent.Tag, 100)
	for i := 0; i < 100; i++ {
		tags[i] = tx.Tag.Create().
			SetName("タグ" + uuid.New().String()[:8]).
			SetCreatedAt(time.Now()).
			SaveX(ctx)
	}

	// Create Thread Tags
	threadTags := make(map[string]struct{})
	for i := 0; i < 100; i++ {
		var threadID, tagID int
		var key string
		for {
			threadID = threads[rand.Intn(100)].ID
			tagID = tags[rand.Intn(100)].ID
			key = fmt.Sprintf("%d-%d", threadID, tagID)
			if _, exists := threadTags[key]; !exists {
				break
			}
		}
		threadTags[key] = struct{}{}
		tx.ThreadTag.Create().
			SetThreadID(threadID).
			SetTagID(tagID).
			SaveX(ctx)
	}

	// Create User Comment Likes
	userCommentLikes := make(map[string]struct{})
	for i := 0; i < 100; i++ {
		var userID, commentID int
		var key string
		for {
			userID = users[rand.Intn(100)].ID
			commentID = comments[rand.Intn(100000)].ID
			key = fmt.Sprintf("%d-%d", userID, commentID)
			if _, exists := userCommentLikes[key]; !exists {
				break
			}
		}
		userCommentLikes[key] = struct{}{}
		tx.UserCommentLike.Create().
			SetLikedAt(time.Now()).
			SetUserID(userID).
			SetCommentID(commentID).
			SaveX(ctx)
	}

	// Create User Comment Subscriptions
	userCommentSubscriptions := make(map[string]struct{})
	for i := 0; i < 100; i++ {
		var userID, commentID int
		var key string
		for {
			userID = users[rand.Intn(100)].ID
			commentID = comments[rand.Intn(100000)].ID
			key = fmt.Sprintf("%d-%d", userID, commentID)
			if _, exists := userCommentSubscriptions[key]; !exists {
				break
			}
		}
		userCommentSubscriptions[key] = struct{}{}
		tx.UserCommentSubscription.Create().
			SetIsNotified(true).
			SetIsChecked(false).
			SetSubscribedAt(time.Now()).
			SetUserID(userID).
			SetCommentID(commentID).
			SaveX(ctx)
	}

	// Create User Thread Likes
	userThreadLikes := make(map[string]struct{})
	for i := 0; i < 100; i++ {
		var userID, threadID int
		var key string
		for {
			userID = users[rand.Intn(100)].ID
			threadID = threads[rand.Intn(100)].ID
			key = fmt.Sprintf("%d-%d", userID, threadID)
			if _, exists := userThreadLikes[key]; !exists {
				break
			}
		}
		userThreadLikes[key] = struct{}{}
		tx.UserThreadLike.Create().
			SetLikedAt(time.Now()).
			SetUserID(userID).
			SetThreadID(threadID).
			SaveX(ctx)
	}

	// Create User Thread Subscriptions
	userThreadSubscriptions := make(map[string]struct{})
	for i := 0; i < 100; i++ {
		var userID, threadID int
		var key string
		for {
			userID = users[rand.Intn(100)].ID
			threadID = threads[rand.Intn(100)].ID
			key = fmt.Sprintf("%d-%d", userID, threadID)
			if _, exists := userThreadSubscriptions[key]; !exists {
				break
			}
		}
		userThreadSubscriptions[key] = struct{}{}
		tx.UserThreadSubscription.Create().
			SetIsNotified(true).
			SetIsChecked(false).
			SetSubscribedAt(time.Now()).
			SetUserID(userID).
			SetThreadID(threadID).
			SaveX(ctx)
	}

	return nil
}
