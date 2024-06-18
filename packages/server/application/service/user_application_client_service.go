package service

import (
    "context"
    "server/infrastructure/datasource"
    "server/infrastructure/ent"
)

type UserApplicationClientService struct {
    userRepo *datasource.UserClientDatasource
}

func NewUserApplicationClientService(userRepo *datasource.UserClientDatasource) *UserApplicationClientService {
    return &UserApplicationClientService{userRepo: userRepo}
}

func (s *UserApplicationClientService) GetUsers(ctx context.Context) ([]*ent.User, error) {
    return s.userRepo.GetAllUsers(ctx)
}
