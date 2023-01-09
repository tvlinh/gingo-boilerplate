package api

import (
	"log"
	"os"
	"tvlinh/gingo-boilerplate/config"
	"tvlinh/gingo-boilerplate/internal/server"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func loadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
}

func Run() error {
	// load env
	loadEnv()

	cfgFile, err := config.LoadConfig(os.Getenv("CONFIG_PATH"))
	if err != nil {
		log.Fatalf("LoadConfig: %v", err)
	}

	cfg, err := config.ParseConfig(cfgFile)
	if err != nil {
		log.Fatalf("ParseConfig: %v", err)
	}

	s := server.NewGinServer(cfg, logrus.New(), nil)

	return s.Run()
}
