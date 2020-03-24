package application

import (
	"github.com/boilerplate/pkg/config"
	"github.com/boilerplate/pkg/db"
)

// Application holds commonly used app wide data, for ease of DI
type Application struct {
	DB  *db.DB
	Cfg *config.Config
}

// Get assembles and captures env vars, establishes DB connection and keeps
// reference both
func Get() (*Application, error) {
	cfg := config.Get()
	db, err := db.Get(cfg.GetDBConnStr())

	if err != nil {
		return nil, err
	}

	return &Application{
		DB:  db,
		Cfg: cfg,
	}, nil
}
