package model

import (
	"server/infrastructure/ent"
)

type User struct {
	EntUser                    *ent.User
	ThreadCommentCount         int         // コメント数
	ThreadCount                int         // スレッド数
	ThreadCommentCountMap      map[int]int // スレッド毎のコメント数
	ThreadCommentReplyCountMap map[int]int // コメント毎のリプライ数
}

type UserStatus int

const (
	UserStatusActive UserStatus = iota
	UserStatusWithdrawn
	UserStatusSuspended
	UserStatusInactive
)

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

type UserRole int

const (
	UserRoleMember UserRole = iota
	UserRoleAdmin
)

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

func (m *User) IsAdmin() bool {
	return UserRole(m.EntUser.Role) == UserRoleAdmin
}

func (m *User) IsMember() bool {
	return UserRole(m.EntUser.Role) == UserRoleMember
}
