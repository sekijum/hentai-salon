//go:build wireinject
// +build wireinject

package di

import (
	applicationService "server/application/service"
	domainService "server/domain/service"
	"server/infrastructure/aws"
	datasource "server/infrastructure/datasource"
	"server/infrastructure/ent"
	"server/infrastructure/minio"
	controller "server/presentation/controller"

	"github.com/google/wire"
)

var externalServiceSet = wire.NewSet(
	ent.ProvideClient,
	aws.NewS3Client,
	minio.NewMinioClient,
)

var controllerSet = wire.NewSet(
	controller.NewBoardController,
	controller.NewUserController,
	controller.NewThreadController,
	controller.NewThreadCommentController,
	controller.NewTagController,
	controller.NewStorageController,
	controller.NewUserAdminController,
	controller.NewBoardAdminController,
	controller.NewThreadAdminController,
)

var applicationServiceSet = wire.NewSet(
	applicationService.NewBoardApplicationService,
	applicationService.NewUserApplicationService,
	applicationService.NewThreadApplicationService,
	applicationService.NewThreadCommentApplicationService,
	applicationService.NewTagApplicationService,
	applicationService.NewStorageApplicationService,
	applicationService.NewUserAdminApplicationService,
	applicationService.NewBoardAdminApplicationService,
	applicationService.NewThreadAdminApplicationService,
)

var domainServiceSet = wire.NewSet(
	domainService.NewBoardDomainService,
	domainService.NewUserDomainService,
	domainService.NewThreadDomainService,
)

var datasourceSet = wire.NewSet(
	datasource.NewBoardDatasource,
	datasource.NewUserDatasource,
	datasource.NewThreadDatasource,
	datasource.NewThreadCommentDatasource,
	datasource.NewTagDatasource,
	datasource.NewUserAdminDatasource,
	datasource.NewBoardAdminDatasource,
	datasource.NewThreadAdminDatasource,
)

type ControllersSet struct {
	BoardController         *controller.BoardController
	UserController          *controller.UserController
	ThreadController        *controller.ThreadController
	ThreadCommentController *controller.ThreadCommentController
	TagController           *controller.TagController
	StorageController       *controller.StorageController
	UserAdminController     *controller.UserAdminController
	BoardAdminController    *controller.BoardAdminController
	ThreadAdminController   *controller.ThreadAdminController
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
