package ioc

import (
	"github.com/emreaknci/goauthexample/api/handler"
	"github.com/emreaknci/goauthexample/internal/config"
	"github.com/emreaknci/goauthexample/internal/repository"
	"github.com/emreaknci/goauthexample/internal/service"
	"go.uber.org/dig"
	"gorm.io/gorm"
)

func BuildContainer() (*dig.Container, error) {
	container := dig.New()

	if err := provideDatabase(container); err != nil {
		panic(err)
	}

	if err := provideRepositories(container); err != nil {
		panic(err)
	}

	if err := provideServices(container); err != nil {
		panic(err)
	}

	if err := provideHandlers(container); err != nil {
		panic(err)
	}

	return container, nil
}

func provideDatabase(container *dig.Container) error {
	db, err := config.SetupDatabase()
	if err != nil {
		panic(err)
	}

	err = container.Provide(func() *gorm.DB {
		return db
	})

	if err != nil {
		panic(err)
	}

	return nil
}

func provideRepositories(container *dig.Container) error {
	err := container.Provide(func(db *gorm.DB) repository.UserRepository {
		return repository.NewUserRepository(db)
	})
	if err != nil {
		panic(err)
	}

	return nil
}

func provideServices(container *dig.Container) error {
	err := container.Provide(func(repo repository.UserRepository) service.AuthService {
		return service.NewAuthService(repo)
	})

	if err != nil {
		panic(err)
	}

	return nil
}

func provideHandlers(container *dig.Container) error {
	err := container.Provide(func(service service.AuthService) handler.AuthHandler {
		return handler.NewAuthHandler(service)
	})
	if err != nil {
		panic(err)
	}

	return nil
}
