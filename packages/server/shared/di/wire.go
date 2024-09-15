//go:build wireinject
// +build wireinject

package di

import (
	"server/application/service"
	service_admin "server/application/service/admin"
	"server/infrastructure/aws"
	"server/infrastructure/datasource"
	datasource_admin "server/infrastructure/datasource/admin"
	"server/infrastructure/mailpit"
	"server/infrastructure/minio"
	"server/presentation/controller"
	controller_admin "server/presentation/controller/admin"

	"github.com/google/wire"
)

var externalServiceSet = wire.NewSet(
	datasource.ProvideClient,
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
	controller.NewContactController,
	controller.NewThreadCommentAttachmentController,
	controller.NewAdController,

	controller_admin.NewUserController,
	controller_admin.NewBoardController,
	controller_admin.NewThreadController,
	controller_admin.NewContactController,
	controller_admin.NewThreadCommentController,
	controller_admin.NewTagController,
	controller_admin.NewAdController,
)

var serviceSet = wire.NewSet(
	service.NewBoardApplicationService,
	service.NewUserApplicationService,
	service.NewThreadApplicationService,
	service.NewThreadCommentApplicationService,
	service.NewTagApplicationService,
	service.NewStorageApplicationService,
	service.NewContactApplicationService,
	service.NewThreadCommentAttachmentApplicationService,
	service.NewAdApplicationService,

	service_admin.NewUserApplicationService,
	service_admin.NewBoardApplicationService,
	service_admin.NewThreadApplicationService,
	service_admin.NewContactApplicationService,
	service_admin.NewThreadCommentApplicationService,
	service_admin.NewTagApplicationService,
	service_admin.NewAdApplicationService,
)

var datasourceSet = wire.NewSet(
	datasource.NewBoardDatasource,
	datasource.NewUserDatasource,
	datasource.NewThreadDatasource,
	datasource.NewThreadCommentDatasource,
	datasource.NewTagDatasource,
	datasource.NewContactDatasource,
	datasource.NewThreadCommentAttachmentDatasource,
	datasource.NewAdDatasource,

	datasource_admin.NewUserDatasource,
	datasource_admin.NewBoardDatasource,
	datasource_admin.NewThreadDatasource,
	datasource_admin.NewContactDatasource,
	datasource_admin.NewThreadCommentDatasource,
	datasource_admin.NewTagDatasource,
	datasource_admin.NewAdDatasource,
)

type ControllersSet struct {
	BoardController                   *controller.BoardController
	UserController                    *controller.UserController
	ThreadController                  *controller.ThreadController
	ThreadCommentController           *controller.ThreadCommentController
	TagController                     *controller.TagController
	StorageController                 *controller.StorageController
	ContactController                 *controller.ContactController
	ThreadCommentAttachmentController *controller.ThreadCommentAttachmentController
	AdController                      *controller.AdController

	UserAdminController          *controller_admin.UserController
	BoardAdminController         *controller_admin.BoardController
	ThreadAdminController        *controller_admin.ThreadController
	ContactAdminController       *controller_admin.ContactController
	ThreadCommentAdminController *controller_admin.ThreadCommentController
	TagAdminController           *controller_admin.TagController
	AdAdminController            *controller_admin.AdController
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
