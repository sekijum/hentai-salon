package datasource

import (
	"context"
	"errors"
	"server/domain/model"
	"server/infrastructure/ent"
	"server/infrastructure/ent/threadcomment"
	"server/infrastructure/ent/user"
	"strings"
	"time"
)

type UserDatasource struct {
	client *ent.Client
}

func NewUserDatasource(client *ent.Client) *UserDatasource {
	return &UserDatasource{client: client}
}

type UserDatasourceFindByIDParams struct {
	Ctx                   context.Context
	UserID, Limit, Offset int
}

type ThreadCommentCount struct {
	ThreadID int `json:"thread_id"`
	Count    int `json:"count"`
}

type ThreadCommentReplyCount struct {
	ParentCommentID int `json:"parent_comment_id"`
	Count           int `json:"count"`
}

func (ds *UserDatasource) FindByID(params UserDatasourceFindByIDParams) (*model.User, error) {
	entUser, err := ds.client.User.Query().
		Where(user.ID(params.UserID)).
		WithComments(func(q *ent.ThreadCommentQuery) {
			q.WithAttachments().
				WithAuthor().
				WithParentComment().
				WithThread().
				Order(ent.Desc(threadcomment.FieldID)).
				Limit(params.Limit).
				Offset(params.Offset)
		}).
		WithThreads(func(q *ent.ThreadQuery) {
			q.WithBoard()
		}).
		Only(params.Ctx)
	if err != nil {
		return nil, err
	}

	modelUser := &model.User{
		EntUser: entUser,
	}

	return modelUser, nil
}

type UserDatasourceIsEmailDuplicatedParams struct {
	Ctx       context.Context
	Email     string
	ExcludeID *int // 更新時に現在のユーザーIDを除外するためのフィールド
}

func (ds *UserDatasource) IsEmailDuplicated(params UserDatasourceIsEmailDuplicatedParams) (bool, error) {
	query := ds.client.User.Query().Where(user.EmailEQ(params.Email))
	if params.ExcludeID != nil {
		query = query.Where(user.IDNEQ(*params.ExcludeID))
	}
	_, err := query.First(params.Ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

type UserDatasourceFindByEmailParams struct {
	Ctx   context.Context
	Email string
}

func (ds *UserDatasource) FindByEmail(params UserDatasourceFindByEmailParams) (*model.User, error) {
	entUser, err := ds.client.User.Query().Where(user.EmailEQ(params.Email)).Only(params.Ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, nil
		}
		return nil, err
	}

	modelUser := &model.User{EntUser: entUser}

	return modelUser, nil
}

type UserDatasourceCreateParams struct {
	Ctx  context.Context
	User *model.User
}

func (ds *UserDatasource) Create(params UserDatasourceCreateParams) (*model.User, error) {
	userBuilder := ds.client.User.Create().
		SetName(params.User.EntUser.Name).
		SetEmail(params.User.EntUser.Email).
		SetPassword(params.User.EntUser.Password).
		SetStatus(params.User.EntUser.Status).
		SetRole(params.User.EntUser.Role)

	if params.User.EntUser.AvatarURL != nil {
		userBuilder.SetAvatarURL(*params.User.EntUser.AvatarURL)
	}
	if params.User.EntUser.ProfileLink != nil {
		userBuilder.SetProfileLink(*params.User.EntUser.ProfileLink)
	}

	savedUser, err := userBuilder.Save(params.Ctx)
	if err != nil {
		if ent.IsConstraintError(err) {
			if strings.Contains(err.Error(), "Duplicate entry") && strings.Contains(err.Error(), "for key 'users.email'") {
				return nil, errors.New("このメールアドレスは既に使用されています。")
			}
		}
		return nil, errors.New("データの制約に違反しています。")
	}

	modelUser := &model.User{EntUser: savedUser}

	return modelUser, nil
}

type UserDatasourceUpdateParams struct {
	Ctx  context.Context
	User model.User
}

func (ds *UserDatasource) Update(params UserDatasourceUpdateParams) (*model.User, error) {
	entUser, err := ds.client.User.Get(params.Ctx, params.User.EntUser.ID)
	if err != nil {
		return nil, err
	}

	update := entUser.Update()

	if params.User.EntUser.Name != "" {
		update.SetName(params.User.EntUser.Name)
	}
	if params.User.EntUser.Email != "" {
		update.SetEmail(params.User.EntUser.Email)
	}
	if params.User.EntUser.Password != "" {
		update.SetPassword(params.User.EntUser.Password)
	}
	if params.User.EntUser.ProfileLink != nil {
		update.SetProfileLink(*params.User.EntUser.ProfileLink)
	}
	if params.User.EntUser.AvatarURL != nil {
		update.SetAvatarURL(*params.User.EntUser.AvatarURL)
	}
	if params.User.EntUser.Status != 0 {
		update.SetStatus(params.User.EntUser.Status)
	}
	update.SetUpdatedAt(time.Now())

	entUser, err = update.Save(params.Ctx)
	if err != nil {
		return nil, err
	}

	return &model.User{EntUser: entUser}, nil
}

type UserDatasourceUpdatePasswordParams struct {
	Ctx      context.Context
	UserID   int
	Password string
}

func (ds *UserDatasource) UpdatePassword(params UserDatasourceUpdatePasswordParams) (*model.User, error) {
	entUser, err := ds.client.User.Get(params.Ctx, params.UserID)
	if err != nil {
		return nil, err
	}

	entUser, err = entUser.Update().SetPassword(params.Password).SetUpdatedAt(time.Now()).Save(params.Ctx)
	if err != nil {
		return nil, err
	}

	return &model.User{EntUser: entUser}, nil
}
