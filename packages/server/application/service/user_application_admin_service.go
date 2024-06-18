package service

import (
    "context"
    "server/infrastructure/datasource"
    "server/infrastructure/ent"
)

type UserApplicationAdminService struct {
    userRepo *datasource.UserAdminDatasource
}

func NewUserApplicationAdminService(userRepo *datasource.UserAdminDatasource) *UserApplicationAdminService {
    return &UserApplicationAdminService{userRepo: userRepo}
}

func (s *UserApplicationAdminService) GetUsers(ctx context.Context) ([]*ent.User, error) {
    return s.userRepo.GetAllUsers(ctx)
}
