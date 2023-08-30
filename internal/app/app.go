package app

import (
	"flag"
	"log"

	"github.com/AnatoliyBr/dynamic-user-segmentation-service/internal/controller/httpserver"
	"github.com/BurntSushi/toml"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config-path", "configs/httpserver.toml", "path to config file")
}

func Run() {

	// Controller
	flag.Parse()
	configServer := httpserver.NewConfig()
	_, err := toml.DecodeFile(configPath, configServer)
	if err != nil {
		log.Fatal(err)
	}

	s := httpserver.NewServer(configServer)
	if err := s.StartServer(); err != nil {
		log.Fatal(err)
	}
}
