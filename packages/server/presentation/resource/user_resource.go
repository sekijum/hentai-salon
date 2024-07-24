package resource

import (
	"server/domain/model"
	"time"
)

type UserResource struct {
	ID          int                                  `json:"id"`
	Name        string                               `json:"name"`
	Email       string                               `json:"email"`
	AvatarURL   *string                              `json:"avatarUrl,omitempty"`
	ProfileLink *string                              `json:"profileLink,omitempty"`
	Role        string                               `json:"role"`
	RoleLabel   string                               `json:"roleLabel"`
	CreatedAt   string                               `json:"createdAt"`
	UpdatedAt   string                               `json:"updatedAt"`
	Comments    ListResource[*ThreadCommentResource] `json:"comments"`
	Attachments []*ThreadCommentAttachmentResource   `json:"attachments"`
	Threads     ListResource[*ThreadResource]        `json:"threads,omitempty"`
}

type NewUserResourceParams struct {
	User          *model.User
	Limit, Offset int
}

func NewUserResource(params NewUserResourceParams) *UserResource {
	var avatarURL *string
	if params.User.EntUser.AvatarURL != nil {
		avatarURL = params.User.EntUser.AvatarURL
	}
	var profileLink *string
	if params.User.EntUser.ProfileLink != nil {
		profileLink = params.User.EntUser.ProfileLink
	}

	var commentResourceList []*ThreadCommentResource
	var attachments []*ThreadCommentAttachmentResource
	for _, comment := range params.User.EntUser.Edges.Comments {
		commentResource := NewThreadCommentResource(NewThreadCommentResourceParams{
			ThreadComment: &model.ThreadComment{EntThreadComment: comment},
			CommentIDs:    nil,
			Offset:        params.Offset,
		})
		commentResourceList = append(commentResourceList, commentResource)

		for _, attachment := range comment.Edges.Attachments {
			threadCommentAttachment := &model.ThreadCommentAttachment{EntAttachment: attachment}
			attachments = append(attachments, &ThreadCommentAttachmentResource{
				Url:          threadCommentAttachment.EntAttachment.URL,
				DisplayOrder: threadCommentAttachment.EntAttachment.DisplayOrder,
				Type:         threadCommentAttachment.TypeToString(),
				CommentID:    comment.ID,
			})
		}
	}

	comments := ListResource[*ThreadCommentResource]{
		TotalCount: params.User.ThreadCommentCount,
		Limit:      params.Limit,
		Offset:     params.Offset,
		Data:       commentResourceList,
	}

	var threads []*ThreadResource
	for _, thread := range params.User.EntUser.Edges.Threads {
		threadResource := NewThreadResource(NewThreadResourceParams{
			Thread: &model.Thread{EntThread: thread},
			Limit:  params.Limit,
			Offset: params.Offset,
		})
		threads = append(threads, threadResource)
	}

	threadList := ListResource[*ThreadResource]{
		TotalCount: params.User.ThreadCount,
		Limit:      params.Limit,
		Offset:     params.Offset,
		Data:       threads,
	}

	return &UserResource{
		ID:          params.User.EntUser.ID,
		Name:        params.User.EntUser.Name,
		Email:       params.User.EntUser.Email,
		AvatarURL:   avatarURL,
		ProfileLink: profileLink,
		Role:        params.User.RoleToString(),
		RoleLabel:   params.User.RoleToLabel(),
		CreatedAt:   params.User.EntUser.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   params.User.EntUser.UpdatedAt.Format(time.RFC3339),
		Comments:    comments,
		Attachments: attachments,
		Threads:     threadList,
	}
}
