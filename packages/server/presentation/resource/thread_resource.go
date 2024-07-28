package resource

import (
	"server/domain/model"
)

type ThreadResource struct {
	ID           int                                           `json:"id"`
	Title        string                                        `json:"title"`
	Description  *string                                       `json:"description,omitempty"`
	ThumbnailURL *string                                       `json:"thumbnailUrl,omitempty"`
	TagNameList  *[]string                                     `json:"tagNameList,omitempty"`
	CommentCount *int                                          `json:"commentCount"`
	Comments     *Collection[*ThreadCommentResource]           `json:"comments,omitempty"`
	Board        *BoardResource                                `json:"board,omitempty"`
	Attachments  *Collection[*ThreadCommentAttachmentResource] `json:"attachments,omitempty"`
	IsLiked      *bool                                         `json:"isLiked,omitempty"`
}

type NewThreadResourceParams struct {
	Thread                                                                *model.Thread
	Limit, Offset                                                         int
	UserID, CommentCount                                                  *int
	IncludeAttachments, IncludeComments, IncludeBoard, IncludeTagNameList bool
}

func NewThreadResource(params NewThreadResourceParams) *ThreadResource {
	var tagNameList []string
	if params.IncludeTagNameList {
		for _, tag_i := range params.Thread.EntThread.Edges.Tags {
			tagNameList = append(tagNameList, tag_i.Name)
		}
	}

	var boardResource *BoardResource
	if params.IncludeBoard {
		if params.Thread.EntThread.Edges.Board != nil {
			boardResource = &BoardResource{
				Id:    params.Thread.EntThread.Edges.Board.ID,
				Title: params.Thread.EntThread.Edges.Board.Title,
			}
		}
	}

	var commentResourceList []*ThreadCommentResource
	var attachmentResourceList []*ThreadCommentAttachmentResource
	var commentCollection *Collection[*ThreadCommentResource]
	var attachmentCollection *Collection[*ThreadCommentAttachmentResource]
	if params.IncludeComments {
		for _, comment_i := range params.Thread.EntThread.Edges.Comments {
			replyCount := len(comment_i.Edges.Replies)
			commentResourceList = append(commentResourceList, NewThreadCommentResource(NewThreadCommentResourceParams{
				ThreadComment: model.NewThreadComment(model.NewThreadCommentParams{
					EntThreadComment: comment_i,
				}),
				ReplyCount:         &replyCount,
				UserID:             params.UserID,
				IncludeAttachments: true,
			}))

			if params.IncludeAttachments {
				for _, attachment_i := range comment_i.Edges.Attachments {
					attachmentResourceList = append(attachmentResourceList, NewThreadCommentAttachmentResource(NewThreadCommentAttachmentResourceParams{
						ThreadCommentAttachment: model.NewThreadCommentAttachment(model.NewThreadCommentAttachmentParams{
							EntAttachment: attachment_i,
						}),
					}))
				}
			}
		}
		if params.CommentCount != nil {
			commentCollection = NewCollection(NewCollectionParams[*ThreadCommentResource]{
				Data:       commentResourceList,
				TotalCount: *params.CommentCount,
				Limit:      params.Limit,
				Offset:     params.Offset,
			})

			attachmentCollection = NewCollection(NewCollectionParams[*ThreadCommentAttachmentResource]{
				Data:       attachmentResourceList,
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

	return &ThreadResource{
		ID:           params.Thread.EntThread.ID,
		Board:        boardResource,
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
