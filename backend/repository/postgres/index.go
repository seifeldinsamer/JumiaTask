package pg

import (
	"backend/pkgs"
	"fmt"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Instance struct {
	orm    *gorm.DB
	Logger *zerolog.Logger
}

func NewPostgresDB() *Instance {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable ",
		pkgs.EnvironmentValues.DBHost, pkgs.EnvironmentValues.DBUserName, pkgs.EnvironmentValues.DBPassword, pkgs.EnvironmentValues.DBName, pkgs.EnvironmentValues.DBPort)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(errors.Errorf("Failed to Intialize Db connection %s", err))
	}
	return &Instance{
		orm: db,
	}
}
