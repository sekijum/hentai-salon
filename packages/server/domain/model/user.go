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
	Status      UserStatus
	Role        UserRole
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

type UserStatus int

const (
	UserRoleActive UserStatus = iota
	UserRoleWithdrawn
	UserRoleSuspended
	UserRoleInactive
)

func (s UserStatus) String() string {
	switch s {
	case UserRoleActive:
		return "Active"
	case UserRoleWithdrawn:
		return "Withdrawn"
	case UserRoleSuspended:
		return "Suspended"
	case UserRoleInactive:
		return "Inactive"
	default:
		return "Unknown"
	}
}

func (s UserStatus) Validate() error {
	switch s {
	case UserRoleActive, UserRoleWithdrawn, UserRoleSuspended, UserRoleInactive:
		return nil
	default:
		return errors.New("無効なユーザー権限です")
	}
}

func (u UserStatus) ToInt() int {
	UserStatusToIntToInt := map[UserStatus]int{
		UserRoleActive:    0,
		UserRoleWithdrawn: 1,
		UserRoleSuspended: 2,
		UserRoleInactive:  3,
	}
	return UserStatusToIntToInt[u]
}

func (s UserStatus) Label() string {
	switch s {
	case UserRoleActive:
		return "有効"
	case UserRoleWithdrawn:
		return "退会済"
	case UserRoleSuspended:
		return "凍結"
	case UserRoleInactive:
		return "無効"
	default:
		return "不明なステータス"
	}
}

type UserRole int

const (
	UserRoleMember UserRole = iota
	UserRoleAdmin
)

func (s UserRole) String() string {
	switch s {
	case UserRoleMember:
		return "Member"
	case UserRoleAdmin:
		return "Admin"
	default:
		return "Unknown"
	}
}

func (s UserRole) Validate() error {
	switch s {
	case UserRoleMember, UserRoleAdmin:
		return nil
	default:
		return errors.New("無効なユーザー権限です")
	}
}

func (u UserRole) ToInt() int {
	boardStatusToIntToInt := map[UserRole]int{
		UserRoleMember: 0,
		UserRoleAdmin:  1,
	}
	return boardStatusToIntToInt[u]
}

func (s UserRole) Label() string {
	switch s {
	case UserRoleMember:
		return "会員"
	case UserRoleAdmin:
		return "管理者"
	default:
		return "不明なステータス"
	}
}

func (u *User) IsAdmin() bool {
	return u.Role == UserRoleAdmin
}
