package response

import (
	"server/domain/model"
)

type ThreadResponse struct {
	ID           int                                           `json:"id"`
	Title        string                                        `json:"title"`
	Description  *string                                       `json:"description,omitempty"`
	ThumbnailURL *string                                       `json:"thumbnailUrl,omitempty"`
	TagNameList  *[]string                                     `json:"tagNameList,omitempty"`
	CommentCount *int                                          `json:"commentCount"`
	Comments     *Collection[*ThreadCommentResponse]           `json:"comments,omitempty"`
	Board        *BoardResponse                                `json:"board,omitempty"`
	Attachments  *Collection[*ThreadCommentAttachmentResponse] `json:"attachments,omitempty"`
	IsLiked      *bool                                         `json:"isLiked,omitempty"`
}

type NewThreadResponseParams struct {
	Thread                                                                *model.Thread
	Limit, Offset                                                         int
	UserID, CommentCount                                                  *int
	IncludeAttachments, IncludeComments, IncludeBoard, IncludeTagNameList bool
}

func NewThreadResponse(params NewThreadResponseParams) *ThreadResponse {
	var tagNameList []string
	if params.IncludeTagNameList {
		for _, tag_i := range params.Thread.EntThread.Edges.Tags {
			tagNameList = append(tagNameList, tag_i.Name)
		}
	}

	var boardResponse *BoardResponse
	if params.IncludeBoard {
		if params.Thread.EntThread.Edges.Board != nil {
			boardResponse = &BoardResponse{
				Id:    params.Thread.EntThread.Edges.Board.ID,
				Title: params.Thread.EntThread.Edges.Board.Title,
			}
		}
	}

	var commentResponseList []*ThreadCommentResponse
	var attachmentResponseList []*ThreadCommentAttachmentResponse
	var commentCollection *Collection[*ThreadCommentResponse]
	var attachmentCollection *Collection[*ThreadCommentAttachmentResponse]
	if params.IncludeComments {
		for _, comment_i := range params.Thread.EntThread.Edges.Comments {
			replyCount := len(comment_i.Edges.Replies)
			commentResponseList = append(commentResponseList, NewThreadCommentResponse(NewThreadCommentResponseParams{
				ThreadComment: model.NewThreadComment(model.NewThreadCommentParams{
					EntThreadComment: comment_i,
				}),
				ReplyCount:         &replyCount,
				UserID:             params.UserID,
				IncludeAttachments: true,
			}))

			if params.IncludeAttachments {
				for _, attachment_i := range comment_i.Edges.Attachments {
					attachmentResponseList = append(attachmentResponseList, NewThreadCommentAttachmentResponse(NewThreadCommentAttachmentResponseParams{
						ThreadCommentAttachment: model.NewThreadCommentAttachment(model.NewThreadCommentAttachmentParams{
							EntAttachment: attachment_i,
						}),
					}))
				}
			}
		}
		if params.CommentCount != nil {
			commentCollection = NewCollection(NewCollectionParams[*ThreadCommentResponse]{
				Data:       commentResponseList,
				TotalCount: *params.CommentCount,
				Limit:      params.Limit,
				Offset:     params.Offset,
			})

			attachmentCollection = NewCollection(NewCollectionParams[*ThreadCommentAttachmentResponse]{
				Data:       attachmentResponseList,
				TotalCount: *params.CommentCount,
				Limit:      params.Limit,
				Offset:     params.Offset,
			})
		}
	}

	var isLiked *bool
	if params.UserID != nil {
		liked := false
		for _, likedUser_i := range params.Thread.EntThread.Edges.LikedUsers {
			if likedUser_i.ID == *params.UserID {
				liked = true
				break
			}
		}
		isLiked = &liked
	}

	return &ThreadResponse{
		ID:           params.Thread.EntThread.ID,
		Board:        boardResponse,
		Title:        params.Thread.EntThread.Title,
		Description:  params.Thread.EntThread.Description,
		ThumbnailURL: params.Thread.EntThread.ThumbnailURL,
		TagNameList:  &tagNameList,
		CommentCount: params.CommentCount,
		Comments:     commentCollection,
		Attachments:  attachmentCollection,
		IsLiked:      isLiked,
	}
}
