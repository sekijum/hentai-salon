// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"github.com/google/wire"
	"server/application/service"
	"server/infrastructure/aws"
	"server/infrastructure/datasource"
	"server/infrastructure/ent"
	"server/infrastructure/mailpit"
	"server/infrastructure/minio"
	"server/presentation/controller"
)

// Injectors from wire.go:

func InitializeControllers() (*ControllersSet, func(), error) {
	client, cleanup, err := ent.ProvideClient()
	if err != nil {
		return nil, nil, err
	}
	boardDatasource := datasource.NewBoardDatasource(client)
	boardApplicationService := service.NewBoardApplicationService(boardDatasource)
	boardController := controller.NewBoardController(boardApplicationService)
	userDatasource := datasource.NewUserDatasource(client)
	mailpitClient := mailpit.NewMailpitClient()
	userApplicationService := service.NewUserApplicationService(userDatasource, mailpitClient)
	userController := controller.NewUserController(userApplicationService)
	threadDatasource := datasource.NewThreadDatasource(client)
	tagDatasource := datasource.NewTagDatasource(client)
	threadApplicationService := service.NewThreadApplicationService(client, threadDatasource, tagDatasource)
	threadController := controller.NewThreadController(threadApplicationService)
	threadCommentDatasource := datasource.NewThreadCommentDatasource(client)
	threadCommentApplicationService := service.NewThreadCommentApplicationService(threadCommentDatasource)
	threadCommentController := controller.NewThreadCommentController(threadCommentApplicationService)
	tagApplicationService := service.NewTagApplicationService(tagDatasource)
	tagController := controller.NewTagController(tagApplicationService)
	minioClient, err := minio.NewMinioClient()
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	s3Client, err := aws.NewS3Client()
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	storageApplicationService := service.NewStorageApplicationService(minioClient, s3Client)
	storageController := controller.NewStorageController(storageApplicationService)
	userAdminDatasource := datasource.NewUserAdminDatasource(client)
	userAdminApplicationService := service.NewUserAdminApplicationService(userAdminDatasource)
	userAdminController := controller.NewUserAdminController(userAdminApplicationService)
	boardAdminDatasource := datasource.NewBoardAdminDatasource(client)
	boardAdminApplicationService := service.NewBoardAdminApplicationService(boardAdminDatasource)
	boardAdminController := controller.NewBoardAdminController(boardAdminApplicationService)
	threadAdminDatasource := datasource.NewThreadAdminDatasource(client)
	threadAdminApplicationService := service.NewThreadAdminApplicationService(threadAdminDatasource)
	threadAdminController := controller.NewThreadAdminController(threadAdminApplicationService)
	controllersSet := &ControllersSet{
		BoardController:         boardController,
		UserController:          userController,
		ThreadController:        threadController,
		ThreadCommentController: threadCommentController,
		TagController:           tagController,
		StorageController:       storageController,
		UserAdminController:     userAdminController,
		BoardAdminController:    boardAdminController,
		ThreadAdminController:   threadAdminController,
	}
	return controllersSet, func() {
		cleanup()
	}, nil
}

// wire.go:

var externalServiceSet = wire.NewSet(ent.ProvideClient, aws.NewS3Client, minio.NewMinioClient, mailpit.NewMailpitClient)

var controllerSet = wire.NewSet(controller.NewBoardController, controller.NewUserController, controller.NewThreadController, controller.NewThreadCommentController, controller.NewTagController, controller.NewStorageController, controller.NewUserAdminController, controller.NewBoardAdminController, controller.NewThreadAdminController)

var serviceSet = wire.NewSet(service.NewBoardApplicationService, service.NewUserApplicationService, service.NewThreadApplicationService, service.NewThreadCommentApplicationService, service.NewTagApplicationService, service.NewStorageApplicationService, service.NewUserAdminApplicationService, service.NewBoardAdminApplicationService, service.NewThreadAdminApplicationService)

var datasourceSet = wire.NewSet(datasource.NewBoardDatasource, datasource.NewUserDatasource, datasource.NewThreadDatasource, datasource.NewThreadCommentDatasource, datasource.NewTagDatasource, datasource.NewUserAdminDatasource, datasource.NewBoardAdminDatasource, datasource.NewThreadAdminDatasource)

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
