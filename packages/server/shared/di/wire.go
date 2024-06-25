//go:build wireinject
// +build wireinject

// この2つがないとパッケージ内で競合する

package di

import (
	"server/application/service"
	"server/infrastructure/datasource"
	"server/infrastructure/ent"
	controller "server/presentation/controller"

	"github.com/google/wire"
)

var entSet = wire.NewSet(
	ent.ProvideClient,
)

var datasourceSet = wire.NewSet(
	datasource.NewBoardClientDatasource,
)

var serviceSet = wire.NewSet(
	service.NewBoardClientService,
)

var controllerSet = wire.NewSet(
	controller.NewBoardClientController,
)

type ControllersSet struct {
	BoardClientController *controller.BoardClientController
}

func InitializeControllers() (*ControllersSet, func(), error) {
	wire.Build(
		entSet,
		datasourceSet,
		serviceSet,
		controllerSet,
		wire.Struct(new(ControllersSet), "*"),
	)
	return &ControllersSet{}, nil, nil
}
