package service_admin

import (
	"context"
	"server/domain/model"
	datasource_admin "server/infrastructure/datasource/admin"
	"server/infrastructure/ent"
	request_admin "server/presentation/request/admin"
	response_admin "server/presentation/response/admin"
)

type UserApplicationService struct {
	userDatasource *datasource_admin.UserDatasource
}

func NewUserApplicationService(userDatasource *datasource_admin.UserDatasource) *UserApplicationService {
	return &UserApplicationService{userDatasource: userDatasource}
}

type UserApplicationServiceFindByIDParams struct {
	Ctx    context.Context
	UserID int
}

func (svc *UserApplicationService) FindByID(params UserApplicationServiceFindByIDParams) (*response_admin.UserResponse, error) {
	user, err := svc.userDatasource.FindByID(datasource_admin.UserDatasourceFindByIDParams{
		Ctx:    params.Ctx,
		UserID: params.UserID,
	})
	if err != nil {
		return nil, err
	}

	dto := response_admin.NewUserResponse(response_admin.NewUserResponseParams{
		User: user,
	})

	return dto, nil
}

type UserApplicationServiceFindAllParams struct {
	Ctx context.Context
	Qs  request_admin.UserFindAllRequest
}

func (svc *UserApplicationService) FindAll(params UserApplicationServiceFindAllParams) (*response_admin.Collection[*response_admin.UserResponse], error) {
	userList, err := svc.userDatasource.FindAll(datasource_admin.UserDatasourceFindAllParams{
		Ctx:     params.Ctx,
		Limit:   params.Qs.Limit,
		Offset:  params.Qs.Offset,
		Sort:    params.Qs.Sort,
		Order:   params.Qs.Order,
		Keyword: params.Qs.Keyword,
	})
	if err != nil {
		return nil, err
	}

	userCount, err := svc.userDatasource.GetUserCount(datasource_admin.UserDatasourceGetUserCountParams{
		Ctx:     params.Ctx,
		Keyword: params.Qs.Keyword,
	})
	if err != nil {
		return nil, err
	}

	var userResponseList []*response_admin.UserResponse
	for _, user_i := range userList {
		userResponseList = append(userResponseList, response_admin.NewUserResponse(response_admin.NewUserResponseParams{
			User: user_i,
		}))
	}

	dto := response_admin.NewCollection(response_admin.NewCollectionParams[*response_admin.UserResponse]{
		Data:       userResponseList,
		TotalCount: userCount,
		Limit:      params.Qs.Limit,
		Offset:     params.Qs.Offset,
	})

	return dto, nil
}

type UserApplicationServiceUpdateParams struct {
	Ctx    context.Context
	UserID int
	Body   request_admin.UserUpdateRequest
}

func (svc *UserApplicationService) Update(params UserApplicationServiceUpdateParams) (*response_admin.UserResponse, error) {
	user := model.NewUser(model.NewUserParams{
		EntUser: &ent.User{
			ID:    params.UserID,
			Name:  params.Body.Name,
			Email: params.Body.Email,
		},
		OptionList: []func(*model.User){
			model.WithUserStatus(model.UserStatus(params.Body.Status)),
			model.WithUserRole(model.UserRole(params.Body.Role)),
		},
	})

	user, err := svc.userDatasource.Update(datasource_admin.UserDatasourceUpdateParams{
		Ctx:  params.Ctx,
		User: *user,
	})
	if err != nil {
		return nil, err
	}

	dto := response_admin.NewUserResponse(response_admin.NewUserResponseParams{User: user})

	return dto, nil
}
