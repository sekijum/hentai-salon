package datasource

import (
    "context"
    "server/infrastructure/ent"
)

type UserClientDatasource struct {
    client *ent.Client
}

func NewUserClientDatasource(client *ent.Client) *UserClientDatasource {
    return &UserClientDatasource{client: client}
}

func (d *UserClientDatasource) GetAllUsers(ctx context.Context) ([]*ent.User, error) {
    return d.client.User.Query().All(ctx)
}
