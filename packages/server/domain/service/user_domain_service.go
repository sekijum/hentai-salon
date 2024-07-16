package service

import (
	"context"
	"server/infrastructure/datasource"
)

type UserDomainService struct {
	userDatasource *datasource.UserDatasource
}

func NewUserDomainService(userDatasource *datasource.UserDatasource) *UserDomainService {
	return &UserDomainService{userDatasource: userDatasource}
}

type NewUserDomainServiceIsUserAdminParams struct {
	Ctx    context.Context
	UserID int
}

func (svc *UserDomainService) IsUserAdmin(params NewUserDomainServiceIsUserAdminParams) (bool, error) {
	user, err := svc.userDatasource.FindByID(datasource.UserDatasourceFindByIDParams{
		Ctx:    params.Ctx,
		UserID: params.UserID,
	})
	if err != nil {
		return false, err
	}
	return user.IsAdmin(), nil
}
