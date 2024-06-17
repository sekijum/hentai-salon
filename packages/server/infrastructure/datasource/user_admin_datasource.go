package datasource

import (
    "context"
    "server/infrastructure/ent"
)

type UserAdminDatasource struct {
    client *ent.Client
}

func NewUserAdminDatasource(client *ent.Client) *UserAdminDatasource {
    return &UserAdminDatasource{client: client}
}

func (d *UserAdminDatasource) GetAllUsers(ctx context.Context) ([]*ent.User, error) {
    return d.client.User.Query().All(ctx)
}
