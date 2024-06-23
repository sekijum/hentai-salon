package di

import (
	repository "server/application/repository"
	service "server/application/service"
	"server/infrastructure/datasource"
	"server/infrastructure/ent"
	controller "server/presentation/controller"

	"github.com/google/wire"
)

func InitializeClientControllers(client *ent.Client) (*controller.BoardClientController, error) {
	wire.Build(
		datasource.NewBoardClientDatasource,
		service.NewBoardClientService,
		controller.NewBoardClientController,
		wire.Bind(new(repository.BoardClientRepository), new(*datasource.BoardClientDatasource)),
	)
	return nil, nil
}
