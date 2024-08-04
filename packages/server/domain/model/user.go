package model

import (
	"server/infrastructure/ent"
)

type User struct {
	EntUser *ent.User
}

type NewUserParams struct {
	EntUser    *ent.User
	OptionList []func(*User)
}

func NewUser(params NewUserParams) *User {
	user := &User{EntUser: params.EntUser}

	for _, option_i := range params.OptionList {
		option_i(user)
	}

	return user
}

type UserStatus int

const (
	UserStatusActive UserStatus = iota
	UserStatusWithdrawn
	UserStatusSuspended
	UserStatusInactive
)

type UserRole int

const (
	UserRoleMember UserRole = iota
	UserRoleAdmin
)

func WithUserStatus(status UserStatus) func(*User) {
	return func(u *User) {
		u.EntUser.Status = int(status)
	}
}

func WithUserRole(role UserRole) func(*User) {
	return func(u *User) {
		u.EntUser.Role = int(role)
	}
}

func (m *User) StatusToString() string {
	switch UserStatus(m.EntUser.Status) {
	case UserStatusActive:
		return "Active"
	case UserStatusWithdrawn:
		return "Withdrawn"
	case UserStatusSuspended:
		return "Suspended"
	case UserStatusInactive:
		return "Inactive"
	default:
		return "Unknown"
	}
}

func (m *User) StatusToLabel() string {
	switch UserStatus(m.EntUser.Status) {
	case UserStatusActive:
		return "有効"
	case UserStatusWithdrawn:
		return "退会済"
	case UserStatusSuspended:
		return "凍結"
	case UserStatusInactive:
		return "無効"
	default:
		return "不明なステータス"
	}
}

func (m *User) RoleToString() string {
	switch UserRole(m.EntUser.Role) {
	case UserRoleMember:
		return "Member"
	case UserRoleAdmin:
		return "Admin"
	default:
		return "Unknown"
	}
}

func (m *User) RoleToLabel() string {
	switch UserRole(m.EntUser.Role) {
	case UserRoleMember:
		return "会員"
	case UserRoleAdmin:
		return "管理者"
	default:
		return "不明なステータス"
	}
}
