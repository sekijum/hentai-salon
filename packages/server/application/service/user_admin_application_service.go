package service

import (
	"context"
	"server/domain/model"
	"server/infrastructure/datasource"
	"server/infrastructure/ent"
	"server/presentation/request"
	"server/presentation/resource"
)

type UserAdminApplicationService struct {
	userAdminDatasource *datasource.UserAdminDatasource
}

func NewUserAdminApplicationService(userAdminDatasource *datasource.UserAdminDatasource) *UserAdminApplicationService {
	return &UserAdminApplicationService{userAdminDatasource: userAdminDatasource}
}

type UserAdminApplicationServiceFindAllParams struct {
	Ctx context.Context
	Qs  request.UserAdminFindAllRequest
}

func (svc *UserAdminApplicationService) FindAll(params UserAdminApplicationServiceFindAllParams) (*resource.Collection[*resource.UserAdminResource], error) {
	userList, err := svc.userAdminDatasource.FindAll(datasource.UserAdminDatasourceFindAllParams{
		Ctx:       params.Ctx,
		Limit:     params.Qs.Limit,
		Offset:    params.Qs.Offset,
		SortKey:   params.Qs.SortKey,
		SortOrder: params.Qs.SortOrder,
		Keyword:   params.Qs.Keyword,
		Role:      params.Qs.Role,
	})
	if err != nil {
		return nil, err
	}

	userCount, err := svc.userAdminDatasource.GetUserCount(datasource.UserAdminDatasourceGetUserCountParams{
		Ctx:     params.Ctx,
		Keyword: params.Qs.Keyword,
		Role:    params.Qs.Role,
	})
	if err != nil {
		return nil, err
	}

	var userResourceList []*resource.UserAdminResource
	for _, user_i := range userList {
		userResourceList = append(userResourceList, resource.NewUserAdminResource(resource.NewUserAdminResourceParams{
			User: user_i,
		}))
	}

	dto := resource.NewCollection(resource.NewCollectionParams[*resource.UserAdminResource]{
		Data:       userResourceList,
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

func (svc *UserAdminApplicationService) Update(params UserAdminApplicationServiceUpdateParams) (*resource.UserAdminResource, error) {
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

	dto := resource.NewUserAdminResource(resource.NewUserAdminResourceParams{User: user})

	return dto, nil
}
