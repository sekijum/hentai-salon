package model

import "time"

type Board struct {
	Id           int
	UserId       int
	Title        string
	Description  string
	ThumbnailUrl string
	Order        int
	Status       string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Threads      []*Thread
}
