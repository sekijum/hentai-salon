package resource

import (
	"fmt"
	"server/domain/model"
	"time"
)

type ThreadBoardResource struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
}

type ThreadResource struct {
	Id           int                  `json:"id"`
	Board        *ThreadBoardResource `json:"board"`
	Title        string               `json:"title"`
	Description  string               `json:"description"`
	ThumbnailUrl string               `json:"thumbnailUrl"`
	Tags         []string             `json:"tags"`
	CreatedAt    string               `json:"createdAt"`
	CommentCount int                  `json:"commentCount"`
	Popularity   string               `json:"popularity"`
}

func NewThreadBoardResource(b *model.Board) *ThreadBoardResource {
	return &ThreadBoardResource{
		Id:    b.EntBoard.ID,
		Title: b.EntBoard.Title,
	}
}

func NewThreadResource(t *model.Thread) *ThreadResource {
	description := ""
	if t.EntThread.Description != "" {
		description = t.EntThread.Description
	}

	thumbnailUrl := t.EntThread.ThumbnailURL

	var tagNames []string
	for _, tag := range t.EntThread.Edges.Tags {
		tagNames = append(tagNames, tag.Name)
	}

	commentCount := len(t.EntThread.Edges.Comments)

	popularity := fmt.Sprintf("%d%%", t.Popularity)

	var boardResource *ThreadBoardResource
	if t.EntThread.Edges.Board != nil {
		boardResource = NewThreadBoardResource(&model.Board{EntBoard: t.EntThread.Edges.Board})
	}

	return &ThreadResource{
		Id:           t.EntThread.ID,
		Board:        boardResource,
		Title:        t.EntThread.Title,
		Description:  description,
		ThumbnailUrl: thumbnailUrl,
		Tags:         tagNames,
		CreatedAt:    t.EntThread.CreatedAt.Format(time.RFC3339),
		CommentCount: commentCount,
		Popularity:   popularity,
	}
}
