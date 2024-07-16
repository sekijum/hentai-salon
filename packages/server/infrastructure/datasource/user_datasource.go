package datasource

import (
	"context"
	"errors"
	"server/domain/model"
	"server/infrastructure/ent"
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
	Ctx    context.Context
	UserID int
}

func (ds *UserDatasource) FindByID(params UserDatasourceFindByIDParams) (*model.User, error) {
	entUser, err := ds.client.User.Get(params.Ctx, params.UserID)
	if err != nil {
		return nil, err
	}

	modelUser := &model.User{EntUser: entUser}

	return modelUser, nil
}

type UserDatasourceFindByEmailParams struct {
	Ctx   context.Context
	Email string
}

func (ds *UserDatasource) FindByEmail(params UserDatasourceFindByEmailParams) (*model.User, error) {
	entUser, err := ds.client.User.Query().Where(user.EmailEQ(params.Email)).Only(params.Ctx)
	if err != nil {
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

func (ds *UserDatasource) Update(params UserDatasourceUpdateParams) (*ent.User, error) {
	user, err := ds.client.User.Get(params.Ctx, params.User.EntUser.ID)
	if err != nil {
		return nil, err
	}

	update := user.Update()

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

	user, err = update.Save(params.Ctx)
	if err != nil {
		return nil, err
	}

	return user, nil
}
