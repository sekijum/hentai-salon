package resource

import (
	"server/domain/model"
	"server/infrastructure/ent"
	"time"
)

type ThreadCommentResource struct {
	ID              int                                 `json:"id"`
	GuestName       *string                             `json:"guestName,omitempty"`
	Content         string                              `json:"content"`
	CreatedAt       string                              `json:"createdAt"`
	ParentCommentID *int                                `json:"parentCommentId,omitempty"`
	ReplyCount      *int                                `json:"replyCount"`
	User            *UserResource                       `json:"user,omitempty"`
	Thread          *ThreadResource                     `json:"thread,omitempty"`
	ParentComment   *ThreadCommentResource              `json:"parentComment,omitempty"`
	Attachments     []*ThreadCommentAttachmentResource  `json:"attachments,omitempty"`
	Replies         *Collection[*ThreadCommentResource] `json:"replies,omitempty"`
	IsLiked         *bool                               `json:"isLiked,omitempty"`
}

type NewThreadCommentResourceParams struct {
	ThreadComment                                                                        *model.ThreadComment
	Limit, Offset                                                                        int
	UserID, ReplyCount                                                                   *int
	IncludeUser, IncludeThread, IncludeParentComment, IncludeAttachments, IncludeReplies bool
}

func NewThreadCommentResource(params NewThreadCommentResourceParams) *ThreadCommentResource {

	var userResource *UserResource
	if params.IncludeUser {
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
	}

	var parentComment *ThreadCommentResource
	if params.IncludeParentComment {
		if params.ThreadComment.EntThreadComment.Edges.ParentComment != nil {
			parentComment = NewThreadCommentResource(NewThreadCommentResourceParams{
				ThreadComment: model.NewThreadComment(model.NewThreadCommentParams{
					EntThreadComment: params.ThreadComment.EntThreadComment.Edges.ParentComment,
				}),
			})
		}
	}

	var attachmentResourceList []*ThreadCommentAttachmentResource
	if params.IncludeAttachments {
		for _, attachment_i := range params.ThreadComment.EntThreadComment.Edges.Attachments {
			attachmentResourceList = append(attachmentResourceList, NewThreadCommentAttachmentResource(NewThreadCommentAttachmentResourceParams{
				ThreadCommentAttachment: model.NewThreadCommentAttachment(model.NewThreadCommentAttachmentParams{
					EntAttachment: attachment_i,
				}),
			}))
		}
	}

	var replyResourceList []*ThreadCommentResource
	if params.IncludeReplies {
		for _, reply_i := range params.ThreadComment.EntThreadComment.Edges.Replies {
			replyCount := len(reply_i.Edges.Replies)
			replyResourceList = append(replyResourceList, NewThreadCommentResource(NewThreadCommentResourceParams{
				ThreadComment: model.NewThreadComment(model.NewThreadCommentParams{EntThreadComment: reply_i}),
				Offset:        params.Offset,
				ReplyCount:    &replyCount,
				UserID:        params.UserID,
			}))
		}
	}

	var replyCollection *Collection[*ThreadCommentResource]
	if params.ReplyCount != nil {
		replyCollection = NewCollection(NewCollectionParams[*ThreadCommentResource]{
			Data:       replyResourceList,
			TotalCount: *params.ReplyCount,
			Limit:      params.Limit,
			Offset:     params.Offset,
		})
	}

	var threadResource *ThreadResource
	if params.IncludeThread {
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
	}

	var isLiked *bool
	if params.UserID != nil {
		liked := false
		for _, likedUser_i := range params.ThreadComment.EntThreadComment.Edges.LikedUsers {
			if likedUser_i.ID == *params.UserID {
				liked = true
				break
			}
		}
		isLiked = &liked
	}

	return &ThreadCommentResource{
		ID:              params.ThreadComment.EntThreadComment.ID,
		User:            userResource,
		GuestName:       params.ThreadComment.EntThreadComment.GuestName,
		Content:         params.ThreadComment.EntThreadComment.Content,
		CreatedAt:       params.ThreadComment.EntThreadComment.CreatedAt.Format(time.RFC3339),
		ParentCommentID: params.ThreadComment.EntThreadComment.ParentCommentID,
		ParentComment:   parentComment,
		Thread:          threadResource,
		Attachments:     attachmentResourceList,
		ReplyCount:      params.ReplyCount,
		Replies:         replyCollection,
		IsLiked:         isLiked,
	}
}
