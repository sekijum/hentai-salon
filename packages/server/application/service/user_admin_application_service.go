package service

import (
	"context"
	"server/infrastructure/datasource"
	"server/presentation/request"
	"server/presentation/resource"
	"time"
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

func (svc *UserAdminApplicationService) FindAll(params UserAdminApplicationServiceFindAllParams) (*resource.ListResource[*resource.UserAdminResource], error) {
	users, err := svc.userAdminDatasource.FindAll(datasource.UserAdminDatasourceFindAllParams{
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

	var userAdminResourceList []*resource.UserAdminResource
	for _, user := range users {
		userAdminResourceList = append(userAdminResourceList, resource.NewUserAdminResource(resource.NewUserAdminResourceParams{
			User: user,
		}))
	}

	dto := &resource.ListResource[*resource.UserAdminResource]{
		TotalCount: userCount,
		Limit:      params.Qs.Limit,
		Offset:     params.Qs.Offset,
		Data:       userAdminResourceList,
	}

	return dto, nil
}

type UserAdminApplicationServiceUpdateParams struct {
	Ctx    context.Context
	UserID int
	Body   request.UserAdminUpdateRequest
}

func (svc *UserAdminApplicationService) Update(params UserAdminApplicationServiceUpdateParams) (*resource.UserAdminResource, error) {
	user, err := svc.userAdminDatasource.FindByID(datasource.UserAdminDatasourceFindByIDParams{
		Ctx:    params.Ctx,
		UserID: params.UserID,
	})
	if err != nil {
		return nil, err
	}

	if params.Body.Name != nil {
		user.EntUser.Name = *params.Body.Name
	}
	if params.Body.Email != nil {
		user.EntUser.Email = *params.Body.Email
	}
	if params.Body.Role != nil {
		user.EntUser.Role = *params.Body.Role
	}
	if params.Body.Status != nil {
		user.EntUser.Status = *params.Body.Status
	}
	user.EntUser.UpdatedAt = time.Now()

	updatedUser, err := svc.userAdminDatasource.Update(datasource.UserAdminDatasourceUpdateParams{
		Ctx:  params.Ctx,
		User: *user,
	})
	if err != nil {
		return nil, err
	}

	dto := resource.NewUserAdminResource(resource.NewUserAdminResourceParams{User: updatedUser})

	return dto, nil
}
