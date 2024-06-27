package datasource

import (
	"context"
	"server/domain/model"
	"server/infrastructure/ent"
	"server/infrastructure/ent/user"

	"github.com/mitchellh/mapstructure"
)

type UserDatasource interface {
	Create(ctx context.Context, user *model.User) (*model.User, error)
	GetUserByEmail(ctx context.Context, email string) (*model.User, error)
	GetUserByID(ctx context.Context, id int) (*model.User, error)
}

type userDatasource struct {
	client *ent.Client
}

func NewUserDatasource(client *ent.Client) UserDatasource {
	return &userDatasource{client: client}
}

func (ds *userDatasource) Create(ctx context.Context, u *model.User) (*model.User, error) {
	userBuilder := ds.client.User.Create().
		SetName(u.Name).
		SetEmail(u.Email).
		SetPassword(u.Password).
		SetStatus(u.Status).
		SetRole(u.Role)
	if u.DisplayName != nil {
		userBuilder.SetDisplayName(*u.DisplayName)
	}
	if u.DisplayName != nil {
		userBuilder.SetAvatarUrl(*u.AvatarUrl)
	}

	savedUser, err := userBuilder.Save(ctx)
	if err != nil {
		return nil, err
	}

	modelUser, err := entUserToModelUser(savedUser)
	if err != nil {
		return nil, err
	}

	return modelUser, nil
}

func (ds *userDatasource) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	entUser, err := ds.client.User.Query().Where(user.EmailEQ(email)).Only(ctx)
	if err != nil {
		return nil, err
	}

	modelUser, err := entUserToModelUser(entUser)
	if err != nil {
		return nil, err
	}

	return modelUser, nil
}

func (ds *userDatasource) GetUserByID(ctx context.Context, id int) (*model.User, error) {
	entUser, err := ds.client.User.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	modelUser, err := entUserToModelUser(entUser)
	if err != nil {
		return nil, err
	}

	return modelUser, nil
}

func entUserToModelUser(entUser *ent.User) (*model.User, error) {
	var modelUser model.User
	err := mapstructure.Decode(entUser, &modelUser)
	if err != nil {
		return nil, err
	}
	return &modelUser, nil
}
