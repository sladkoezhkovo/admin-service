package pg

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/sladkoezhkovo/admin-service/internal/config"
	"os"
)

const (
	CityTable     = "city"
	UnitTable     = "units"
	DistrictTable = "district"
)

func Setup(cfg *config.SqlConfig) (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), cfg.Db, cfg.SSL))
	if err != nil {
		return nil, err
	}

	return db, nil
}
