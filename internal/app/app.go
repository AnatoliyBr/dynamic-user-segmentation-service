package app

import (
	"flag"
	"log"

	"github.com/AnatoliyBr/dynamic-user-segmentation-service/internal/controller/httpserver"
	"github.com/AnatoliyBr/dynamic-user-segmentation-service/internal/repository"
	"github.com/AnatoliyBr/dynamic-user-segmentation-service/internal/repository/sqlrepository"
	"github.com/AnatoliyBr/dynamic-user-segmentation-service/internal/usecase"
	"github.com/BurntSushi/toml"
	"github.com/joho/godotenv"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config-path", "configs/httpserver.toml", "path to config file")
}

func Run() {

	// PostgreSQL
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	configDB := repository.NewConfig()
	db, err := repository.NewDB(configDB)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Repository
	r := sqlrepository.NewSegmentRepository(db)

	// UseCase
	uc := usecase.NewAppUseCase(r)

	// Controller
	flag.Parse()
	configServer := httpserver.NewConfig()
	_, err = toml.DecodeFile(configPath, configServer)
	if err != nil {
		log.Fatal(err)
	}

	s := httpserver.NewServer(configServer, uc)
	if err := s.StartServer(); err != nil {
		log.Fatal(err)
	}
}
