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

func (svc *UserDomainService) IsUserAdmin(ctx context.Context, userID int) (bool, error) {
	user, err := svc.userDatasource.FindByID(ctx, userID)
	if err != nil {
		return false, err
	}
	return user.IsAdmin(), nil
}
