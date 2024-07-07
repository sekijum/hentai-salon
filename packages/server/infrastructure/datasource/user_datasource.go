package datasource

import (
	"context"
	"server/domain/model"
	"server/infrastructure/ent"
	"server/infrastructure/ent/user"
)

type UserDatasource struct {
	client *ent.Client
}

func NewUserDatasource(client *ent.Client) *UserDatasource {
	return &UserDatasource{client: client}
}

func (ds *UserDatasource) FindByID(ctx context.Context, id int) (*model.User, error) {
	entUser, err := ds.client.User.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	modelUser := &model.User{
		EntUser: entUser,
	}

	return modelUser, nil
}

func (ds *UserDatasource) FindByEmail(ctx context.Context, email string) (*model.User, error) {
	entUser, err := ds.client.User.Query().Where(user.EmailEQ(email)).Only(ctx)
	if err != nil {
		return nil, err
	}

	modelUser := &model.User{
		EntUser: entUser,
	}

	return modelUser, nil
}

func (ds *UserDatasource) Create(ctx context.Context, m *model.User) (*model.User, error) {
	userBuilder := ds.client.User.Create().
		SetName(m.EntUser.Name).
		SetEmail(m.EntUser.Email).
		SetPassword(m.EntUser.Password).
		SetStatus(m.EntUser.Status).
		SetRole(m.EntUser.Role)

	if m.EntUser.AvatarURL != nil {
		userBuilder.SetAvatarURL(*m.EntUser.AvatarURL)
	}

	savedUser, err := userBuilder.Save(ctx)
	if err != nil {
		return nil, err
	}

	modelUser := &model.User{
		EntUser: savedUser,
	}

	return modelUser, nil
}
