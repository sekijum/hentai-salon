package response

import (
	"server/domain/model"
	"server/infrastructure/ent"
	"strconv"
	"time"
)

type ThreadCommentResponse struct {
	ID              string                              `json:"id"`
	GuestName       *string                             `json:"guestName,omitempty"`
	Content         string                              `json:"content"`
	CreatedAt       string                              `json:"createdAt"`
	ParentCommentID *string                             `json:"parentCommentId,omitempty"`
	ReplyCount      *int                                `json:"replyCount"`
	User            *UserResponse                       `json:"user,omitempty"`
	Thread          *ThreadResponse                     `json:"thread,omitempty"`
	ParentComment   *ThreadCommentResponse              `json:"parentComment,omitempty"`
	Attachments     []*ThreadCommentAttachmentResponse  `json:"attachments,omitempty"`
	Replies         *Collection[*ThreadCommentResponse] `json:"replies,omitempty"`
	IsLiked         *bool                               `json:"isLiked,omitempty"`
}

type NewThreadCommentResponseParams struct {
	ThreadComment                                                                        *model.ThreadComment
	Limit, Offset                                                                        int
	UserID, ReplyCount                                                                   *int
	IncludeUser, IncludeThread, IncludeParentComment, IncludeAttachments, IncludeReplies bool
}

func NewThreadCommentResponse(params NewThreadCommentResponseParams) *ThreadCommentResponse {

	var userResponse *UserResponse
	if params.IncludeUser {
		if params.ThreadComment.EntThreadComment.Edges.Author != nil {
			userResponse = NewUserResponse(NewUserResponseParams{
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

	var parentComment *ThreadCommentResponse
	if params.IncludeParentComment {
		if params.ThreadComment.EntThreadComment.Edges.ParentComment != nil {
			parentComment = NewThreadCommentResponse(NewThreadCommentResponseParams{
				ThreadComment: model.NewThreadComment(model.NewThreadCommentParams{
					EntThreadComment: params.ThreadComment.EntThreadComment.Edges.ParentComment,
				}),
			})
		}
	}

	var attachmentResponseList []*ThreadCommentAttachmentResponse
	if params.IncludeAttachments {
		for _, attachment_i := range params.ThreadComment.EntThreadComment.Edges.Attachments {
			var commentAuthorName *string
			if params.ThreadComment.EntThreadComment.Edges.Author != nil {
				commentAuthorName = &params.ThreadComment.EntThreadComment.Edges.Author.Name
			} else if params.ThreadComment.EntThreadComment.GuestName != nil {
				commentAuthorName = params.ThreadComment.EntThreadComment.GuestName
			}

			createdAt := params.ThreadComment.EntThreadComment.CreatedAt.Format(time.RFC3339)

			attachmentResponseList = append(attachmentResponseList, NewThreadCommentAttachmentResponse(NewThreadCommentAttachmentResponseParams{
				ThreadCommentAttachment: model.NewThreadCommentAttachment(model.NewThreadCommentAttachmentParams{
					EntAttachment: attachment_i,
				}),
				CommentAuthorName: commentAuthorName,
				CommentContent:    &params.ThreadComment.EntThreadComment.Content,
				CreatedAt:         &createdAt,
			}))
		}
	}

	var replyResponseList []*ThreadCommentResponse
	if params.IncludeReplies {
		for _, reply_i := range params.ThreadComment.EntThreadComment.Edges.Replies {
			replyCount := len(reply_i.Edges.Replies)
			replyResponseList = append(replyResponseList, NewThreadCommentResponse(NewThreadCommentResponseParams{
				ThreadComment:      model.NewThreadComment(model.NewThreadCommentParams{EntThreadComment: reply_i}),
				Offset:             params.Offset,
				ReplyCount:         &replyCount,
				UserID:             params.UserID,
				IncludeAttachments: true,
			}))
		}
	}

	var replyCollection *Collection[*ThreadCommentResponse]
	if params.ReplyCount != nil {
		replyCollection = NewCollection(NewCollectionParams[*ThreadCommentResponse]{
			Data:       replyResponseList,
			TotalCount: *params.ReplyCount,
			Limit:      params.Limit,
			Offset:     params.Offset,
		})
	}

	var threadResponse *ThreadResponse
	if params.IncludeThread {
		if params.ThreadComment.EntThreadComment.Edges.Thread != nil {
			var description *string
			if params.ThreadComment.EntThreadComment.Edges.Thread.Description != nil {
				description = params.ThreadComment.EntThreadComment.Edges.Thread.Description
			}
			threadResponse = &ThreadResponse{
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

	var parentCommentID *string
	if params.ThreadComment.EntThreadComment.ParentCommentID != nil {
		idStr := strconv.FormatUint(*params.ThreadComment.EntThreadComment.ParentCommentID, 10)
		parentCommentID = &idStr
	}

	return &ThreadCommentResponse{
		ID:              strconv.FormatUint(params.ThreadComment.EntThreadComment.ID, 10),
		User:            userResponse,
		GuestName:       params.ThreadComment.EntThreadComment.GuestName,
		Content:         params.ThreadComment.EntThreadComment.Content,
		CreatedAt:       params.ThreadComment.EntThreadComment.CreatedAt.Format(time.RFC3339),
		ParentCommentID: parentCommentID,
		ParentComment:   parentComment,
		Thread:          threadResponse,
		Attachments:     attachmentResponseList,
		ReplyCount:      params.ReplyCount,
		Replies:         replyCollection,
		IsLiked:         isLiked,
	}
}
