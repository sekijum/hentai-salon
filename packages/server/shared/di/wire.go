//go:build wireinject
// +build wireinject

package di

import (
	applicationService "server/application/service"
	domainService "server/domain/service"
	aws "server/infrastructure/aws"
	"server/infrastructure/datasource"
	"server/infrastructure/ent"
	minio "server/infrastructure/minio"
	"server/presentation/controller"

	"github.com/google/wire"
)

var controllerSet = wire.NewSet(
	controller.NewBoardController,
	controller.NewUserController,
	controller.NewThreadController,
	controller.NewThreadCommentController,
	controller.NewTagController,
	controller.NewStorageController,
)

var applicationServiceSet = wire.NewSet(
	applicationService.NewBoardApplicationService,
	applicationService.NewUserApplicationService,
	applicationService.NewThreadApplicationService,
	applicationService.NewThreadCommentApplicationService,
	applicationService.NewTagApplicationService,
	applicationService.NewStorageApplicationService,
)

var domainServiceSet = wire.NewSet(
	domainService.NewBoardDomainService,
	domainService.NewUserDomainService,
	domainService.NewThreadDomainService,
)

var externalServiceSet = wire.NewSet(
	ent.ProvideClient,
	aws.NewS3Client,
	minio.NewMinioClient,
)

var datasourceSet = wire.NewSet(
	datasource.NewBoardDatasource,
	datasource.NewUserDatasource,
	datasource.NewThreadDatasource,
	datasource.NewThreadCommentDatasource,
	datasource.NewTagDatasource,
)

type ControllersSet struct {
	BoardController         *controller.BoardController
	UserController          *controller.UserController
	ThreadController        *controller.ThreadController
	ThreadCommentController *controller.ThreadCommentController
	TagController           *controller.TagController
	StorageController       *controller.StorageController
}

func InitializeControllers() (*ControllersSet, func(), error) {
	wire.Build(
		controllerSet,
		applicationServiceSet,
		domainServiceSet,
		externalServiceSet,
		datasourceSet,
		wire.Struct(new(ControllersSet), "*"),
	)
	return &ControllersSet{}, nil, nil
}
