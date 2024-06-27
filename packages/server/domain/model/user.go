package model

import (
	"errors"
	"net/url"
	"regexp"
	"time"
)

type User struct {
	Id          int
	Name        string
	Email       string
	Password    string
	DisplayName *string
	AvatarUrl   *string
	Status      int
	Role        int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (u *User) Validate() error {
	if u.Name == "" {
		return errors.New("名前は必須です")
	}
	if len(u.Name) > 20 {
		return errors.New("名前は20文字以内で入力してください")
	}
	if u.Email == "" {
		return errors.New("メールアドレスは必須です")
	}
	if !isValidEmail(u.Email) {
		return errors.New("有効なメールアドレスを入力してください")
	}
	if u.Password == "" {
		return errors.New("パスワードは必須です")
	}
	if len(u.Password) < 6 {
		return errors.New("パスワードは6文字以上で入力してください")
	}
	if u.DisplayName != nil && len(*u.DisplayName) > 20 {
		return errors.New("表示名は20文字以内で入力してください")
	}
	if u.AvatarUrl != nil {
		if _, err := url.ParseRequestURI(*u.AvatarUrl); err != nil {
			return errors.New("アバターURLは有効なURLである必要があります")
		}
	}
	if u.Status < 0 || u.Status > 2 {
		return errors.New("無効なステータスです")
	}
	if u.Role < 0 || u.Role > 1 {
		return errors.New("無効なロールです")
	}
	return nil
}

func isValidEmail(email string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(email)
}
