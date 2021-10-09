package main

import (
	"backend/api/handler"
	"backend/pkgs"
	"backend/pkgs/files"
	"backend/pkgs/filters"
	pg "backend/repository/postgres"
	usersUseCase "backend/usecase/users"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	zeroLog "github.com/rs/zerolog/log"
	"log"
	"net/http"
)

var router *gin.Engine

func init() {
	pkgs.Must(pkgs.LoadEnvironment())

	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	logger := zeroLog.With().Logger()

	instance := pg.NewPostgresDB()
	instance.Migrate()
	//	load csvFile
	logger.Info().Msgf("Start loading...")
	users := files.ReadCSVFile("/root/users.csv")
	logger.Info().Msgf("Finish loading...")

	instance.LoadUsers(users,
		[]filters.Filter{
			filters.CountryHandler(),
		},
	)

	router = handler.NewGinHandler(usersUseCase.LoadService(instance, &logger))
}

func main() {
	log.Fatal(http.ListenAndServe(":3000", router))
}
