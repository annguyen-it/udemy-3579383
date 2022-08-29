package dbrepo

import (
	"database/sql"
	"learn-golang/internal/config"
	"learn-golang/internal/repository"
)

type postgresDBRepo struct {
	App *config.AppConfig
	DB  *sql.DB
}

func NewPostgresDBRepo(db *sql.DB, a *config.AppConfig) repository.DatabaseRepo {
	return &postgresDBRepo{
		App: a,
		DB:  db,
	}
}
