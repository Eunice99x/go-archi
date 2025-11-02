package di

import (
	"github.com/eunice99x/dingoSuck/internal/repository"
	"github.com/eunice99x/dingoSuck/internal/service"
	"github.com/sarulabs/dingo/v4"
)

const (
	appService = "user_service"
)

func getAppService() []dingo.Def {
	return []dingo.Def{
		{
			Name: appService,
			Build: func(repo interface{}) (service.UserService, error) {
				return service.NewUserService(repo.(repository.UserRepository)), nil
			},
			Params: dingo.Params{
				"0": dingo.Service("user_repo"),
			},
		},
	}
}