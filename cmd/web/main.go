package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/fiqriardiansyah/user-shopping-api-golang/internal/config"
	_ "github.com/joho/godotenv/autoload"
	"github.com/sirupsen/logrus"
)

func init() {
	num64, err := strconv.ParseUint(os.Getenv("LOG_LEVEL"), 10, 32)
	if err != nil {
		logrus.Error(err)
	}
	num32 := uint32(num64)
	logrus.SetLevel(logrus.Level(num32))
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
		ForceColors:   true,
	})
}

func main() {
	db := config.NewDB()
	validate := config.NewValidator()
	fiber, err := config.NewFiber()
	cfg := config.NewConfig()
	if err != nil {
		panic(err)
	}

	config.NewApp(&config.AppConfig{
		App:      fiber,
		Db:       db,
		Validate: validate,
		Config:   cfg,
	})

	webport := os.Getenv("PORT")
	if err := fiber.Listen(fmt.Sprintf(":%s", webport)); err != nil {
		logrus.Fatalf("Failed to start server: %v", err)
	}
}
