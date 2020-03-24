// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package di

import (
	"github.com/ajordi/todo/internal/api"
	"github.com/ajordi/todo/internal/dao"
	"github.com/ajordi/todo/pkg/adding"
	"github.com/ajordi/todo/pkg/authenticating"
	"github.com/ajordi/todo/pkg/deleting"
	"github.com/ajordi/todo/pkg/listing"
	"github.com/jinzhu/gorm"
)

// Injectors from wire.go:

func InitApp(db *gorm.DB) *api.Api {
	storage := dao.NewTaskStorage(db)
	taskRepository := dao.NewAddingTaskRepository(storage)
	taskService := adding.ProvideTaskService(taskRepository)
	userStorage := dao.NewUserStorage(db)
	userRepository := dao.NewAuthenticatingUserRepository(userStorage)
	authenticateService := authenticating.ProviderAuthenticatingService(userRepository)
	listingTaskRepository := dao.NewListingTaskRepository(storage)
	listingTaskService := listing.ProvideTaskService(listingTaskRepository)
	deletingTaskRepository := dao.NewDeletingTaskRepository(storage)
	deletingTaskService := deleting.ProvideTaskService(deletingTaskRepository)
	apiApi := api.New(taskService, authenticateService, listingTaskService, deletingTaskService)
	return apiApi
}
