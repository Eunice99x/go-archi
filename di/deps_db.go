package di

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/sarulabs/dingo/v4"
)

const dbService = "db"

func getAppDB() []dingo.Def {
	return []dingo.Def{
		{
			Name: dbService,
			Build: func() (*sql.DB, error) {
				db, err := sql.Open("postgres", "postgres://postgres:eunice99x@localhost:5432/dblearn?sslmode=disable")
				if err != nil {
					log.Println("failed to connect to db:", err)
					return nil, err
				}
				return db, nil
			},
		},
	}
}
