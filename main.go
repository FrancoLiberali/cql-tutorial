package main

import (
	"github.com/ditrit/badaas/orm"
	"go.uber.org/fx"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	fx.New(
		// connect to db
		fx.Provide(NewDBConnection),

		// execute tutorial
		fx.Invoke(tutorial),
	).Run()
}

func NewDBConnection() (*gorm.DB, error) {
	return orm.Open(
		sqlite.Open(orm.CreateSQLiteDSN("db")),
	)
}
