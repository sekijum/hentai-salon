package service

import (
	"context"
	"server/domain/model"
	"server/infrastructure/datasource"
	"server/infrastructure/ent"
	"server/presentation/request"
	"server/presentation/response"
)

type UserAdminApplicationService struct {
	userAdminDatasource *datasource.UserAdminDatasource
}

func NewUserAdminApplicationService(userAdminDatasource *datasource.UserAdminDatasource) *UserAdminApplicationService {
	return &UserAdminApplicationService{userAdminDatasource: userAdminDatasource}
}

type UserAdminApplicationServiceFindByIDParams struct {
	Ctx    context.Context
	UserID int
}

func (svc *UserAdminApplicationService) FindByID(params UserAdminApplicationServiceFindByIDParams) (*response.UserAdminResponse, error) {
	user, err := svc.userAdminDatasource.FindByID(datasource.UserAdminDatasourceFindByIDParams{
		Ctx:    params.Ctx,
		UserID: params.UserID,
	})
	if err != nil {
		return nil, err
	}

	dto := response.NewUserAdminResponse(response.NewUserAdminResponseParams{
		User: user,
	})

	return dto, nil
}

type UserAdminApplicationServiceFindAllParams struct {
	Ctx context.Context
	Qs  request.UserAdminFindAllRequest
}

func (svc *UserAdminApplicationService) FindAll(params UserAdminApplicationServiceFindAllParams) (*response.Collection[*response.UserAdminResponse], error) {
	userList, err := svc.userAdminDatasource.FindAll(datasource.UserAdminDatasourceFindAllParams{
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

	userCount, err := svc.userAdminDatasource.GetUserCount(datasource.UserAdminDatasourceGetUserCountParams{
		Ctx:     params.Ctx,
		Keyword: params.Qs.Keyword,
	})
	if err != nil {
		return nil, err
	}

	var userResponseList []*response.UserAdminResponse
	for _, user_i := range userList {
		userResponseList = append(userResponseList, response.NewUserAdminResponse(response.NewUserAdminResponseParams{
			User: user_i,
		}))
	}

	dto := response.NewCollection(response.NewCollectionParams[*response.UserAdminResponse]{
		Data:       userResponseList,
		TotalCount: userCount,
		Limit:      params.Qs.Limit,
		Offset:     params.Qs.Offset,
	})

	return dto, nil
}

type UserAdminApplicationServiceUpdateParams struct {
	Ctx    context.Context
	UserID int
	Body   request.UserAdminUpdateRequest
}

func (svc *UserAdminApplicationService) Update(params UserAdminApplicationServiceUpdateParams) (*response.UserAdminResponse, error) {
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

	user, err := svc.userAdminDatasource.Update(datasource.UserAdminDatasourceUpdateParams{
		Ctx:  params.Ctx,
		User: *user,
	})
	if err != nil {
		return nil, err
	}

	dto := response.NewUserAdminResponse(response.NewUserAdminResponseParams{User: user})

	return dto, nil
}
