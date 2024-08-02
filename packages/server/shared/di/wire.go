//go:build wireinject
// +build wireinject

package di

import (
	"server/application/service"
	service_admin "server/application/service/admin"
	"server/infrastructure/aws"
	"server/infrastructure/datasource"
	datasource_admin "server/infrastructure/datasource/admin"
	"server/infrastructure/ent"
	"server/infrastructure/mailpit"
	"server/infrastructure/minio"
	"server/presentation/controller"
	controller_admin "server/presentation/controller/admin"

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
	controller_admin.NewUserController,
	controller_admin.NewBoardController,
	controller_admin.NewThreadController,
	controller.NewContactController,
	controller_admin.NewContactController,
	controller_admin.NewThreadCommentController,
)

var serviceSet = wire.NewSet(
	service.NewBoardApplicationService,
	service.NewUserApplicationService,
	service.NewThreadApplicationService,
	service.NewThreadCommentApplicationService,
	service.NewTagApplicationService,
	service.NewStorageApplicationService,
	service_admin.NewUserApplicationService,
	service_admin.NewBoardApplicationService,
	service_admin.NewThreadApplicationService,
	service.NewContactApplicationService,
	service_admin.NewContactApplicationService,
	service_admin.NewThreadCommentApplicationService,
)

var datasourceSet = wire.NewSet(
	datasource.NewBoardDatasource,
	datasource.NewUserDatasource,
	datasource.NewThreadDatasource,
	datasource.NewThreadCommentDatasource,
	datasource.NewTagDatasource,
	datasource_admin.NewUserDatasource,
	datasource_admin.NewBoardDatasource,
	datasource_admin.NewThreadDatasource,
	datasource.NewContactDatasource,
	datasource_admin.NewContactDatasource,
	datasource_admin.NewThreadCommentDatasource,
)

type ControllersSet struct {
	BoardController              *controller.BoardController
	UserController               *controller.UserController
	ThreadController             *controller.ThreadController
	ThreadCommentController      *controller.ThreadCommentController
	TagController                *controller.TagController
	StorageController            *controller.StorageController
	UserAdminController          *controller_admin.UserController
	BoardAdminController         *controller_admin.BoardController
	ThreadAdminController        *controller_admin.ThreadController
	ContactController            *controller.ContactController
	ContactAdminController       *controller_admin.ContactController
	ThreadCommentAdminController *controller_admin.ThreadCommentController
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
