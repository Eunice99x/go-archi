package di

import (
	"database/sql"

	"github.com/eunice99x/dingoSuck/internal/repository"
	"github.com/sarulabs/dingo/v4"
)

const (
	appRepo = "user_repo"
)

func getAppRepo() []dingo.Def{
	return []dingo.Def{
		{
			Name: appRepo,
			Build: func(db interface{}) (repository.UserRepository, error) {
				return repository.NewUserRepository(db.(*sql.DB)), nil
			},
			Params: dingo.Params{
				"0": dingo.Service("db"),
			},
		},
	}
}