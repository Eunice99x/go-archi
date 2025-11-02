package di

import (
	"github.com/eunice99x/dingoSuck/internal/api"
	"github.com/eunice99x/dingoSuck/internal/service"
	"github.com/sarulabs/dingo/v4"
)

const (
	appHandler = "user_handler"
)

func getAppHandler() []dingo.Def {
	return []dingo.Def{
		{
			Name: appHandler,
			Build: func(s interface{}) (*api.UserHandler, error) {
				return api.NewUserHandler(s.(service.UserService)), nil
			},
			Params: dingo.Params{
				"0": dingo.Service("user_service"),
			},
		},
	}
}