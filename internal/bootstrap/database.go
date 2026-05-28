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

	db := database.NewDatabase(
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.Username,
		cfg.Database.Password,
		cfg.Database.DbName,
		cfg.Database.Sslmode,
		cfg.Database.MaxIdleConn,
		cfg.Database.MaxOpenConn,
		cfg.Database.MaxConnLifetime,
	)

	pool, err := db.Connect()
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	return pool
}
