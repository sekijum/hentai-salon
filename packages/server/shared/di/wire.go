//go:build wireinject
// +build wireinject

package di

import (
	"server/application/service"
	"server/infrastructure/aws"
	"server/infrastructure/datasource"
	"server/infrastructure/ent"
	"server/infrastructure/mailpit"
	"server/infrastructure/minio"
	"server/presentation/controller"

	"github.com/google/wire"
)

var externalServiceSet = wire.NewSet(
	ent.ProvideClient,
	aws.NewS3Client,
	aws.NewSESClient,
	minio.NewMinioClient,
	mailpit.NewMailpitClient,
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
	controller.NewContactController,
	controller.NewContactAdminController,
	controller.NewThreadCommentAdminController,
)

var serviceSet = wire.NewSet(
	service.NewBoardApplicationService,
	service.NewUserApplicationService,
	service.NewThreadApplicationService,
	service.NewThreadCommentApplicationService,
	service.NewTagApplicationService,
	service.NewStorageApplicationService,
	service.NewUserAdminApplicationService,
	service.NewBoardAdminApplicationService,
	service.NewThreadAdminApplicationService,
	service.NewContactApplicationService,
	service.NewContactAdminApplicationService,
	service.NewThreadCommentAdminApplicationService,
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
	datasource.NewContactDatasource,
	datasource.NewContactAdminDatasource,
	datasource.NewThreadCommentAdminDatasource,
)

type ControllersSet struct {
	BoardController              *controller.BoardController
	UserController               *controller.UserController
	ThreadController             *controller.ThreadController
	ThreadCommentController      *controller.ThreadCommentController
	TagController                *controller.TagController
	StorageController            *controller.StorageController
	UserAdminController          *controller.UserAdminController
	BoardAdminController         *controller.BoardAdminController
	ThreadAdminController        *controller.ThreadAdminController
	ContactController            *controller.ContactController
	ContactAdminController       *controller.ContactAdminController
	ThreadCommentAdminController *controller.ThreadCommentAdminController
}

func InitializeControllers() (*ControllersSet, func(), error) {
	wire.Build(
		controllerSet,
		serviceSet,
		externalServiceSet,
		datasourceSet,
		wire.Struct(new(ControllersSet), "*"),
	)
	return &ControllersSet{}, nil, nil
}
