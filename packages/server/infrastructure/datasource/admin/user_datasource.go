package datasource_admin

import (
	"context"
	"server/domain/model"
	"server/infrastructure/ent"
	"server/infrastructure/ent/user"
	"strconv"
	"time"
)

type UserDatasource struct {
	client *ent.Client
}

func NewUserDatasource(client *ent.Client) *UserDatasource {
	return &UserDatasource{client: client}
}

type UserDatasourceGetUserCountParams struct {
	Ctx     context.Context
	Keyword *string
}

func (ds *UserDatasource) GetUserCount(params UserDatasourceGetUserCountParams) (int, error) {
	q := ds.client.User.Query()

	if params.Keyword != nil && *params.Keyword != "" {
		switch {
		case len(*params.Keyword) > 4 && (*params.Keyword)[:5] == "role:":
			if role, err := strconv.Atoi((*params.Keyword)[5:]); err == nil {
				q = q.Where(user.RoleEQ(role))
			}
		case len(*params.Keyword) > 7 && (*params.Keyword)[:7] == "status:":
			if status, err := strconv.Atoi((*params.Keyword)[7:]); err == nil {
				q = q.Where(user.StatusEQ(status))
			}
		case len(*params.Keyword) > 3 && (*params.Keyword)[:3] == "id:":
			if id, err := strconv.Atoi((*params.Keyword)[3:]); err == nil {
				q = q.Where(user.IDEQ(id))
			}
		default:
			q = q.Where(user.Or(
				user.NameContainsFold(*params.Keyword),
				user.EmailContainsFold(*params.Keyword),
			))
		}
	}

	userCount, err := q.Count(params.Ctx)
	if err != nil {
		return 0, err
	}
	return userCount, nil
}

type UserDatasourceFindByIDParams struct {
	Ctx    context.Context
	UserID int
}

func (ds *UserDatasource) FindByID(params UserDatasourceFindByIDParams) (*model.User, error) {
	entUser, err := ds.client.
		User.
		Get(params.Ctx, params.UserID)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, err
		}
	}

	user := model.NewUser(model.NewUserParams{EntUser: entUser})

	return user, nil
}

type UserDatasourceFindAllParams struct {
	Ctx     context.Context
	Limit   int
	Offset  int
	Sort    *string
	Order   *string
	Keyword *string
}

func (ds *UserDatasource) FindAll(params UserDatasourceFindAllParams) ([]*model.User, error) {
	q := ds.client.User.Query()

	sort := user.FieldID
	order := "desc"

	if params.Sort != nil && *params.Sort != "" {
		sort = *params.Sort
	}
	if params.Order != nil && *params.Order != "" {
		order = *params.Order
	}

	if order == "asc" {
		q = q.Order(ent.Asc(sort))
	} else {
		q = q.Order(ent.Desc(sort))
	}

	if params.Keyword != nil && *params.Keyword != "" {
		switch {
		case len(*params.Keyword) > 4 && (*params.Keyword)[:5] == "role:":
			if role, err := strconv.Atoi((*params.Keyword)[5:]); err == nil {
				q = q.Where(user.RoleEQ(role))
			}
		case len(*params.Keyword) > 7 && (*params.Keyword)[:7] == "status:":
			if status, err := strconv.Atoi((*params.Keyword)[7:]); err == nil {
				q = q.Where(user.StatusEQ(status))
			}
		case len(*params.Keyword) > 3 && (*params.Keyword)[:3] == "id:":
			if id, err := strconv.Atoi((*params.Keyword)[3:]); err == nil {
				q = q.Where(user.IDEQ(id))
			}
		default:
			q = q.Where(user.Or(
				user.NameContainsFold(*params.Keyword),
				user.EmailContainsFold(*params.Keyword),
			))
		}
	}

	q = q.Limit(params.Limit)
	q = q.Offset(params.Offset)

	entUserList, err := q.All(params.Ctx)
	if err != nil {
		return nil, err
	}

	var modelUsers []*model.User
	for _, entUser_i := range entUserList {
		modelUsers = append(modelUsers, model.NewUser(model.NewUserParams{EntUser: entUser_i}))
	}

	return modelUsers, nil
}

type UserDatasourceUpdateParams struct {
	Ctx  context.Context
	User model.User
}

func (ds *UserDatasource) Update(params UserDatasourceUpdateParams) (*model.User, error) {
	q := ds.client.
		User.
		UpdateOneID(params.User.EntUser.ID)

	q = q.
		SetName(params.User.EntUser.Name).
		SetEmail(params.User.EntUser.Email).
		SetRole(params.User.EntUser.Role).
		SetStatus(params.User.EntUser.Status).
		SetUpdatedAt(time.Now())

	entUser, err := q.Save(params.Ctx)
	if err != nil {
		return nil, err
	}

	user := model.NewUser(model.NewUserParams{
		EntUser: entUser,
	})

	return user, nil
}

type UserDatasourceUpdateStatusParams struct {
	Ctx  context.Context
	User model.User
}

func (ds *UserDatasource) UpdateStatus(params UserDatasourceUpdateStatusParams) (*model.User, error) {
	q := ds.client.User.UpdateOneID(params.User.EntUser.ID)

	q = q.
		SetStatus(params.User.EntUser.Status).
		SetUpdatedAt(time.Now())

	entUser, err := q.Save(params.Ctx)
	if err != nil {
		return nil, err
	}

	user := model.NewUser(model.NewUserParams{
		EntUser: entUser,
	})

	return user, nil
}
