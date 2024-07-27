package resource

import (
	"server/domain/model"
	"server/infrastructure/ent"
	"time"
)

type ThreadCommentResource struct {
	ID              int                                 `json:"id"`
	User            *UserResource                       `json:"user,omitempty"`
	GuestName       *string                             `json:"guestName,omitempty"`
	Content         string                              `json:"content"`
	ParentCommentID *int                                `json:"parentCommentId"`
	CreatedAt       string                              `json:"createdAt"`
	UpdatedAt       string                              `json:"updatedAt"`
	Thread          *ThreadResource                     `json:"thread,omitempty"`
	ParentComment   *ThreadCommentResource              `json:"parentComment"`
	Attachments     []*ThreadCommentAttachmentResource  `json:"attachments"`
	ReplyCount      int                                 `json:"replyCount"`
	Replies         *Collection[*ThreadCommentResource] `json:"replies"`
}

type NewThreadCommentResourceParams struct {
	ThreadComment             *model.ThreadComment
	CommentIDs                []int
	ReplyCount, Limit, Offset int
}

func NewThreadCommentResource(params NewThreadCommentResourceParams) *ThreadCommentResource {

	var userResource *UserResource
	if params.ThreadComment.EntThreadComment.Edges.Author != nil {
		userResource = NewUserResource(NewUserResourceParams{
			User: model.NewUser(model.NewUserParams{
				EntUser: &ent.User{
					ID:          params.ThreadComment.EntThreadComment.Edges.Author.ID,
					Name:        params.ThreadComment.EntThreadComment.Edges.Author.Name,
					ProfileLink: params.ThreadComment.EntThreadComment.Edges.Author.ProfileLink,
				},
			}),
		})
	}

	var parentComment *ThreadCommentResource
	if params.ThreadComment.EntThreadComment.Edges.ParentComment != nil {
		parentComment = NewThreadCommentResource(NewThreadCommentResourceParams{
			ThreadComment: model.NewThreadComment(model.NewThreadCommentParams{
				EntThreadComment: params.ThreadComment.EntThreadComment.Edges.ParentComment,
			}),
		})
	}

	var attachmentResourceList []*ThreadCommentAttachmentResource
	for _, attachment_i := range params.ThreadComment.EntThreadComment.Edges.Attachments {
		attachmentResourceList = append(attachmentResourceList, NewThreadCommentAttachmentResource(NewThreadCommentAttachmentResourceParams{
			ThreadCommentAttachment: model.NewThreadCommentAttachment(model.NewThreadCommentAttachmentParams{
				EntAttachment: attachment_i,
			}),
		}))
	}

	var replyResourceList []*ThreadCommentResource
	for _, reply_i := range params.ThreadComment.EntThreadComment.Edges.Replies {
		replyResourceList = append(replyResourceList, NewThreadCommentResource(NewThreadCommentResourceParams{
			ThreadComment: model.NewThreadComment(model.NewThreadCommentParams{EntThreadComment: reply_i}),
			Offset:        params.Offset,
			ReplyCount:    len(reply_i.Edges.Replies),
		}))
	}

	replyCollection := NewCollection(NewCollectionParams[*ThreadCommentResource]{
		Data:       replyResourceList,
		TotalCount: params.ReplyCount,
		Limit:      params.Limit,
		Offset:     params.Offset,
	})

	var threadResource *ThreadResource
	if params.ThreadComment.EntThreadComment.Edges.Thread != nil {
		var description *string
		if params.ThreadComment.EntThreadComment.Edges.Thread.Description != nil {
			description = params.ThreadComment.EntThreadComment.Edges.Thread.Description
		}
		threadResource = &ThreadResource{
			ID:          params.ThreadComment.EntThreadComment.Edges.Thread.ID,
			Title:       params.ThreadComment.EntThreadComment.Edges.Thread.Title,
			Description: description,
		}
	}

	return &ThreadCommentResource{
		ID:              params.ThreadComment.EntThreadComment.ID,
		User:            userResource,
		GuestName:       params.ThreadComment.EntThreadComment.GuestName,
		Content:         params.ThreadComment.EntThreadComment.Content,
		ParentCommentID: params.ThreadComment.EntThreadComment.ParentCommentID,
		ParentComment:   parentComment,
		CreatedAt:       params.ThreadComment.EntThreadComment.CreatedAt.Format(time.RFC3339),
		UpdatedAt:       params.ThreadComment.EntThreadComment.UpdatedAt.Format(time.RFC3339),
		Thread:          threadResource,
		Attachments:     attachmentResourceList,
		ReplyCount:      params.ReplyCount,
		Replies:         replyCollection,
	}
}
