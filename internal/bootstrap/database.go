package bootstrap

import (
	"github.com/arthurhzna/go-clean-architecture/internal/config"
	"github.com/arthurhzna/go-clean-architecture/internal/infrastructure/logging"
	"github.com/arthurhzna/go-clean-architecture/internal/infrastructure/persistence/database"
	"github.com/jmoiron/sqlx"
)

func NewDatabase(
	cfg *config.Config,
	log *logging.ZeroLogger,
) *sqlx.DB {

	dbCfg := cfg.Database

	db := database.NewDatabase(
		dbCfg.Host,
		dbCfg.Port,
		dbCfg.Username,
		dbCfg.Password,
		dbCfg.DbName,
		dbCfg.Sslmode,
		dbCfg.MaxIdleConn,
		dbCfg.MaxOpenConn,
		dbCfg.MaxConnLifetime,
	)

	pool, err := db.Connect()
	if err != nil {
		log.Fatal(err)
	}

	return pool
}
