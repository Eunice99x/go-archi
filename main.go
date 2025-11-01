package main

import (
	"github.com/eunice99x/dingoSuck/internal/db"
)

func main () {
	db.MigrateConn()
}