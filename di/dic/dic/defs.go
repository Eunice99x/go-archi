package dic

import (
	"errors"

	"github.com/sarulabs/di/v2"
	"github.com/sarulabs/dingo/v4"

	sql "database/sql"

	api "github.com/eunice99x/dingoSuck/internal/api"
	repository "github.com/eunice99x/dingoSuck/internal/repository"
	service "github.com/eunice99x/dingoSuck/internal/service"
)

func getDiDefs(provider dingo.Provider) []di.Def {
	return []di.Def{
		{
			Name:  "db",
			Scope: "",
			Build: func(ctn di.Container) (interface{}, error) {
				d, err := provider.Get("db")
				if err != nil {
					var eo *sql.DB
					return eo, err
				}
				b, ok := d.Build.(func() (*sql.DB, error))
				if !ok {
					var eo *sql.DB
					return eo, errors.New("could not cast build function to func() (*sql.DB, error)")
				}
				return b()
			},
			Unshared: false,
		},
		{
			Name:  "user_handler",
			Scope: "",
			Build: func(ctn di.Container) (interface{}, error) {
				d, err := provider.Get("user_handler")
				if err != nil {
					var eo *api.UserHandler
					return eo, err
				}
				pi0, err := ctn.SafeGet("user_service")
				if err != nil {
					var eo *api.UserHandler
					return eo, err
				}
				p0, ok := pi0.(interface{})
				if !ok {
					var eo *api.UserHandler
					return eo, errors.New("could not cast parameter 0 to interface{}")
				}
				b, ok := d.Build.(func(interface{}) (*api.UserHandler, error))
				if !ok {
					var eo *api.UserHandler
					return eo, errors.New("could not cast build function to func(interface{}) (*api.UserHandler, error)")
				}
				return b(p0)
			},
			Unshared: false,
		},
		{
			Name:  "user_repo",
			Scope: "",
			Build: func(ctn di.Container) (interface{}, error) {
				d, err := provider.Get("user_repo")
				if err != nil {
					var eo repository.UserRepository
					return eo, err
				}
				pi0, err := ctn.SafeGet("db")
				if err != nil {
					var eo repository.UserRepository
					return eo, err
				}
				p0, ok := pi0.(interface{})
				if !ok {
					var eo repository.UserRepository
					return eo, errors.New("could not cast parameter 0 to interface{}")
				}
				b, ok := d.Build.(func(interface{}) (repository.UserRepository, error))
				if !ok {
					var eo repository.UserRepository
					return eo, errors.New("could not cast build function to func(interface{}) (repository.UserRepository, error)")
				}
				return b(p0)
			},
			Unshared: false,
		},
		{
			Name:  "user_service",
			Scope: "",
			Build: func(ctn di.Container) (interface{}, error) {
				d, err := provider.Get("user_service")
				if err != nil {
					var eo service.UserService
					return eo, err
				}
				pi0, err := ctn.SafeGet("user_repo")
				if err != nil {
					var eo service.UserService
					return eo, err
				}
				p0, ok := pi0.(interface{})
				if !ok {
					var eo service.UserService
					return eo, errors.New("could not cast parameter 0 to interface{}")
				}
				b, ok := d.Build.(func(interface{}) (service.UserService, error))
				if !ok {
					var eo service.UserService
					return eo, errors.New("could not cast build function to func(interface{}) (service.UserService, error)")
				}
				return b(p0)
			},
			Unshared: false,
		},
	}
}
