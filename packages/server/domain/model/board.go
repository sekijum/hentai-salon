package model

import (
	"time"
)

type Board struct {
	Id           int
	UserId       int
	Title        string
	Description  *string
	ThumbnailUrl *string
	Status       BoardStatus
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Threads      []*Thread
}

func NewBoard(UserId int,
	Title string,
	Description *string,
	ThumbnailUrl *string,
	Status BoardStatus,
	CreatedAt time.Time,
	UpdatedAt time.Time,
	Threads []*Thread) *Board {
	board := &Board{}
	return board
}

type BoardStatus int

const (
	BoardStatusPublic BoardStatus = iota
	BoardStatusPrivate
	BoardStatusArchived
	BoardStatusDeleted
)

func (s BoardStatus) String() string {
	switch s {
	case BoardStatusPublic:
		return "Public"
	case BoardStatusPrivate:
		return "Private"
	case BoardStatusArchived:
		return "Archived"
	case BoardStatusDeleted:
		return "Deleted"
	default:
		return "Unknown"
	}
}
