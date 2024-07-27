package datasource

import (
	"context"
	"server/domain/model"
	"server/infrastructure/ent"
	"server/infrastructure/ent/user"
	"time"
)

type UserAdminDatasource struct {
	client *ent.Client
}

func NewUserAdminDatasource(client *ent.Client) *UserAdminDatasource {
	return &UserAdminDatasource{client: client}
}

type UserAdminDatasourceGetUserCountParams struct {
	Ctx     context.Context
	Keyword *string
	Role    *int
}

func (ds *UserAdminDatasource) GetUserCount(params UserAdminDatasourceGetUserCountParams) (int, error) {
	query := ds.client.User.Query()

	if params.Keyword != nil && *params.Keyword != "" {
		query = query.Where(user.Or(
			user.NameContains(*params.Keyword),
			user.EmailContains(*params.Keyword),
		))
	}

	if params.Role != nil && *params.Role != 0 {
		query = query.Where(user.RoleEQ(*params.Role))
	}

	userCount, err := query.Count(params.Ctx)
	if err != nil {
		return 0, err
	}
	return userCount, nil
}

type UserAdminDatasourceFindByIDParams struct {
	Ctx    context.Context
	UserID int
}

func (ds *UserAdminDatasource) FindByID(params UserAdminDatasourceFindByIDParams) (*model.User, error) {
	entUser, err := ds.client.User.Get(params.Ctx, params.UserID)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, err
		}
	}

	user := model.NewUser(model.NewUserParams{EntUser: entUser})

	return user, nil
}

type UserAdminDatasourceFindAllParams struct {
	Ctx       context.Context
	Limit     int
	Offset    int
	SortKey   *string
	SortOrder *string
	Keyword   *string
	Role      *int
}

func (ds *UserAdminDatasource) FindAll(params UserAdminDatasourceFindAllParams) ([]*model.User, error) {
	query := ds.client.User.Query()

	sortKey := user.FieldID
	if params.SortKey != nil && *params.SortKey != "" {
		sortKey = *params.SortKey
	}

	if params.SortOrder != nil && *params.SortOrder == "asc" {
		query = query.Order(ent.Asc(sortKey))
	} else {
		query = query.Order(ent.Desc(sortKey))
	}

	if params.Keyword != nil && *params.Keyword != "" {
		query = query.Where(user.NameContains(*params.Keyword))
	}

	if params.Role != nil && *params.Role != 0 {
		query = query.Where(user.RoleEQ(*params.Role))
	}

	query = query.Limit(params.Limit)
	query = query.Offset(params.Offset)

	entUserList, err := query.All(params.Ctx)
	if err != nil {
		return nil, err
	}

	var modelUsers []*model.User
	for _, entUser_i := range entUserList {
		modelUsers = append(modelUsers, model.NewUser(model.NewUserParams{EntUser: entUser_i}))
	}

	return modelUsers, nil
}

type UserAdminDatasourceUpdateParams struct {
	Ctx  context.Context
	User model.User
}

func (ds *UserAdminDatasource) Update(params UserAdminDatasourceUpdateParams) (*model.User, error) {
	update := ds.client.User.UpdateOneID(params.User.EntUser.ID)

	if params.User.EntUser.Role != 0 {
		update = update.SetRole(params.User.EntUser.Role)
	}
	if params.User.EntUser.Name != "" {
		update = update.SetName(params.User.EntUser.Name)
	}
	if params.User.EntUser.Email != "" {
		update = update.SetEmail(params.User.EntUser.Email)
	}
	if params.User.EntUser.Password != "" {
		update = update.SetPassword(params.User.EntUser.Password)
	}
	if params.User.EntUser.ProfileLink != nil {
		update = update.SetProfileLink(*params.User.EntUser.ProfileLink)
	}
	if params.User.EntUser.Status != 0 {
		update = update.SetStatus(params.User.EntUser.Status)
	}
	update = update.SetUpdatedAt(time.Now())

	entUser, err := update.Save(params.Ctx)
	if err != nil {
		return nil, err
	}

	user := model.NewUser(model.NewUserParams{
		EntUser: entUser,
	})

	return user, nil
}
