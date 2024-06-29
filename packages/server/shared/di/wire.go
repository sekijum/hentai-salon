//go:build wireinject
// +build wireinject

// この2つがないとパッケージ内で競合する

package di

import (
	applicationService "server/application/service"
	domainService "server/domain/service"
	"server/infrastructure/datasource"
	"server/infrastructure/ent"
	"server/presentation/controller"

	"github.com/google/wire"
)

var controllerSet = wire.NewSet(
	controller.NewBoardAdminController,
	controller.NewBoardController,
	controller.NewUserController,
	controller.NewThreadController,
)

var applicationServiceSet = wire.NewSet(
	applicationService.NewBoardAdminApplicationService,
	applicationService.NewBoardApplicationService,
	applicationService.NewUserApplicationService,
	applicationService.NewThreadApplicationService,
)

var domainServiceSet = wire.NewSet(
	domainService.NewBoardDomainService,
	domainService.NewUserDomainService,
	domainService.NewThreadDomainService,
)

var entSet = wire.NewSet(
	ent.ProvideClient,
)

var datasourceSet = wire.NewSet(
	datasource.NewBoardDatasource,
	datasource.NewUserDatasource,
	datasource.NewThreadDatasource,
)

type ControllersSet struct {
	BoardController      *controller.BoardController
	BoardAdminController *controller.BoardAdminController
	UserController       *controller.UserController
	ThreadController     *controller.ThreadController
}

func InitializeControllers() (*ControllersSet, func(), error) {
	wire.Build(
		controllerSet,
		applicationServiceSet,
		domainServiceSet,
		entSet,
		datasourceSet,
		wire.Struct(new(ControllersSet), "*"),
	)
	return &ControllersSet{}, nil, nil
}
