package resource

import (
	"server/domain/model"
	"time"
)

type ThreadResource struct {
	ID           int                                 `json:"id"`
	Board        *BoardResource                      `json:"board"`
	Title        string                              `json:"title"`
	Description  *string                             `json:"description"`
	ThumbnailURL *string                             `json:"thumbnailUrl"`
	TagNameList  []string                            `json:"tagNameList"`
	CreatedAt    string                              `json:"createdAt"`
	CommentCount int                                 `json:"commentCount"`
	Comments     *Collection[*ThreadCommentResource] `json:"comments"`
	Attachments  []*ThreadCommentAttachmentResource  `json:"attachments"`
	IsLiked      bool                                `json:"isLiked"`
}

type NewThreadResourceParams struct {
	Thread                      *model.Thread
	UserID                      *int
	CommentCount, Limit, Offset int
}

func NewThreadResource(params NewThreadResourceParams) *ThreadResource {
	var tagNameList []string
	for _, tag_i := range params.Thread.EntThread.Edges.Tags {
		tagNameList = append(tagNameList, tag_i.Name)
	}

	var boardResource *BoardResource
	if params.Thread.EntThread.Edges.Board != nil {
		boardResource = &BoardResource{
			Id:    params.Thread.EntThread.Edges.Board.ID,
			Title: params.Thread.EntThread.Edges.Board.Title,
		}
	}

	var commentResourceList []*ThreadCommentResource
	var attachments []*ThreadCommentAttachmentResource
	for _, comment_i := range params.Thread.EntThread.Edges.Comments {
		commentResourceList = append(commentResourceList, NewThreadCommentResource(NewThreadCommentResourceParams{
			ThreadComment: model.NewThreadComment(model.NewThreadCommentParams{
				EntThreadComment: comment_i,
			}),
			ReplyCount: len(comment_i.Edges.Replies),
			UserID:     params.UserID,
		}))

		for _, attachment_i := range comment_i.Edges.Attachments {
			attachments = append(attachments, NewThreadCommentAttachmentResource(NewThreadCommentAttachmentResourceParams{
				ThreadCommentAttachment: model.NewThreadCommentAttachment(model.NewThreadCommentAttachmentParams{
					EntAttachment: attachment_i,
				}),
			}))
		}
	}

	commentCollection := NewCollection(NewCollectionParams[*ThreadCommentResource]{
		Data:       commentResourceList,
		TotalCount: params.CommentCount,
		Limit:      params.Limit,
		Offset:     params.Offset,
	})

	isLiked := false
	if params.UserID != nil {
		for _, likedUser_i := range params.Thread.EntThread.Edges.LikedUsers {
			if likedUser_i.ID == *params.UserID {
				isLiked = true
				break
			}
		}
	}

	return &ThreadResource{
		ID:           params.Thread.EntThread.ID,
		Board:        boardResource,
		Title:        params.Thread.EntThread.Title,
		Description:  params.Thread.EntThread.Description,
		ThumbnailURL: params.Thread.EntThread.ThumbnailURL,
		TagNameList:  tagNameList,
		CreatedAt:    params.Thread.EntThread.CreatedAt.Format(time.RFC3339),
		CommentCount: params.CommentCount,
		Comments:     commentCollection,
		Attachments:  attachments,
		IsLiked:      isLiked,
	}
}
