package main

import (
	"go.uber.org/fx"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/ditrit/badaas/orm"
	"github.com/ditrit/badaas/orm/logger"
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
		&gorm.Config{Logger: logger.Default.ToLogMode(logger.Info)},
	)
}
