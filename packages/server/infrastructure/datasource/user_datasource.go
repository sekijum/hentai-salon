package datasource

import (
	"context"
	"errors"
	"server/domain/model"
	"server/infrastructure/ent"
	"server/infrastructure/ent/threadcomment"
	"server/infrastructure/ent/user"
	"server/infrastructure/ent/userthreadlike"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
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

func (ds *UserDatasource) FindByID(params UserDatasourceFindByIDParams) (*model.User, error) {
	entUser, err := ds.client.
		User.
		Query().
		Where(user.ID(params.UserID)).
		Where(user.StatusEQ(0)).
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

	user := model.NewUser(model.NewUserParams{EntUser: entUser})

	return user, nil
}

type UserDatasourceIsEmailDuplicatedParams struct {
	Ctx       context.Context
	Email     string
	ExcludeID *int // 更新時に現在のユーザーIDを除外するためのフィールド
}

func (ds *UserDatasource) IsEmailDuplicated(params UserDatasourceIsEmailDuplicatedParams) (bool, error) {
	q := ds.client.
		User.
		Query().
		Where(user.EmailEQ(params.Email)).
		Where(user.StatusEQ(0))
	if params.ExcludeID != nil {
		q = q.Where(user.IDNEQ(*params.ExcludeID))
	}

	_, err := q.First(params.Ctx)
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
	entUser, err := ds.client.
		User.
		Query().
		Where(user.EmailEQ(params.Email)).
		Where(user.StatusEQ(0)).
		Only(params.Ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, nil
		}
		return nil, err
	}

	user := model.NewUser(model.NewUserParams{EntUser: entUser})

	return user, nil
}

type UserDatasourceCreateParams struct {
	Ctx  context.Context
	User *model.User
}

func (ds *UserDatasource) Create(params UserDatasourceCreateParams) (*model.User, error) {
	q := ds.client.User.Create().
		SetName(params.User.EntUser.Name).
		SetEmail(params.User.EntUser.Email).
		SetPassword(params.User.EntUser.Password).
		SetStatus(params.User.EntUser.Status).
		SetRole(params.User.EntUser.Role)

	if params.User.EntUser.ProfileLink != nil {
		q.SetProfileLink(*params.User.EntUser.ProfileLink)
	}

	entUser, err := q.Save(params.Ctx)
	if err != nil {
		if ent.IsConstraintError(err) {
			if strings.Contains(err.Error(), "Duplicate entry") && strings.Contains(err.Error(), "for key 'users.email'") {
				return nil, errors.New("このメールアドレスは既に使用されています。")
			}
		}
		return nil, errors.New("データの制約に違反しています。")
	}

	user := model.NewUser(model.NewUserParams{EntUser: entUser})

	return user, nil
}

type UserDatasourceUpdateParams struct {
	Ctx  context.Context
	User *model.User
}

func (ds *UserDatasource) Update(params UserDatasourceUpdateParams) (*model.User, error) {
	q := ds.client.User.UpdateOneID(params.User.EntUser.ID)

	q = q.
		SetName(params.User.EntUser.Name).
		SetEmail(params.User.EntUser.Email).
		SetStatus(params.User.EntUser.Status).
		SetUpdatedAt(time.Now())
	if params.User.EntUser.ProfileLink != nil {
		q.SetProfileLink(*params.User.EntUser.ProfileLink)
	}

	entUser, err := q.Save(params.Ctx)
	if err != nil {
		return nil, err
	}

	user := model.NewUser(model.NewUserParams{EntUser: entUser})

	return user, nil
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

	user := model.NewUser(model.NewUserParams{EntUser: entUser})

	return user, nil
}

type UserDatasourceFindLikedThreadsParams struct {
	Ctx    context.Context
	UserID int
	Limit  int
	Offset int
}

func (ds *UserDatasource) FindLikedThreads(params UserDatasourceFindLikedThreadsParams) ([]*model.Thread, error) {
	entUser, err := ds.client.
		User.
		Query().
		Where(user.ID(params.UserID)).
		Where(user.StatusEQ(0)).
		WithLikedThreads(func(q *ent.ThreadQuery) {
			q.WithTags().
				WithBoard().
				WithComments().
				Order(func(s *sql.Selector) {
					s.OrderBy(sql.Desc(userthreadlike.FieldLikedAt))
				}).
				Limit(params.Limit).Offset(params.Offset)
		}).
		Only(params.Ctx)
	if err != nil {
		return nil, err
	}

	var threads []*model.Thread
	for _, entThread_i := range entUser.Edges.LikedThreads {
		threads = append(threads, model.NewThread(model.NewThreadParams{EntThread: entThread_i}))
	}

	return threads, nil
}

type UserDatasourceFindLikedCommentsParams struct {
	Ctx    context.Context
	UserID int
	Limit  int
	Offset int
}

func (ds *UserDatasource) FindLikedComments(params UserDatasourceFindLikedCommentsParams) ([]*model.ThreadComment, error) {
	entUser, err := ds.client.
		User.
		Query().
		Where(user.ID(params.UserID)).
		Where(user.StatusEQ(0)).
		WithLikedComments(func(q *ent.ThreadCommentQuery) {
			q.WithReplies().
				WithAttachments().
				WithThread().
				WithLikedUsers().
				Order(func(s *sql.Selector) {
					s.OrderBy(sql.Desc(userthreadlike.FieldLikedAt))
				}).
				Limit(params.Limit).
				Offset(params.Offset)
		}).
		Only(params.Ctx)
	if err != nil {
		return nil, err
	}

	var comments []*model.ThreadComment
	for _, entComment_i := range entUser.Edges.LikedComments {
		comments = append(comments, model.NewThreadComment(model.NewThreadCommentParams{EntThreadComment: entComment_i}))
	}

	return comments, nil
}

type UserDatasourceFindThreadsParams struct {
	Ctx    context.Context
	UserID int
	Limit  int
	Offset int
}

func (ds *UserDatasource) FindThreads(params UserDatasourceFindThreadsParams) ([]*model.Thread, error) {
	entUser, err := ds.client.
		User.
		Query().
		Where(user.ID(params.UserID)).
		Where(user.StatusEQ(0)).
		WithThreads(func(q *ent.ThreadQuery) {
			q.WithTags().
				WithBoard().
				WithComments()
			q.Limit(params.Limit).Offset(params.Offset).
				WithComments(func(q *ent.ThreadCommentQuery) {
					q.Select(threadcomment.FieldID)
				})
		}).
		Only(params.Ctx)
	if err != nil {
		return nil, err
	}

	var threads []*model.Thread
	for _, entThread_i := range entUser.Edges.Threads {
		threads = append(threads, model.NewThread(model.NewThreadParams{EntThread: entThread_i}))
	}

	return threads, nil
}

type UserDatasourceFindUserCommentsParams struct {
	Ctx    context.Context
	UserID int
	Limit  int
	Offset int
}

func (ds *UserDatasource) FindUserComments(params UserDatasourceFindUserCommentsParams) ([]*model.ThreadComment, error) {
	entUser, err := ds.client.
		User.
		Query().
		Where(user.ID(params.UserID)).
		Where(user.StatusEQ(0)).
		WithComments(func(q *ent.ThreadCommentQuery) {
			q.WithThread().
				WithAuthor().
				WithAttachments().
				WithParentComment().
				WithReplies().
				WithLikedUsers().
				Limit(params.Limit).
				Offset(params.Offset).
				Order(ent.Desc(threadcomment.FieldCreatedAt))
		}).
		Only(params.Ctx)
	if err != nil {
		return nil, err
	}

	var comments []*model.ThreadComment
	for _, entComment_i := range entUser.Edges.Comments {
		comments = append(comments, model.NewThreadComment(model.NewThreadCommentParams{EntThreadComment: entComment_i}))
	}

	return comments, nil
}

type UserDatasourceSuspendedParams struct {
	Ctx    context.Context
	UserID int
}

func (ds *UserDatasource) Suspended(params UserDatasourceSuspendedParams) error {
	err := ds.client.
		User.
		UpdateOneID(params.UserID).
		SetStatus(2).
		Exec(params.Ctx)
	if err != nil {
		return err
	}

	return nil
}
