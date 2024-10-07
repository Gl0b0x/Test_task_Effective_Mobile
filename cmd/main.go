package main

import (
	"log"

	"EffectiveMobile/config"
	_ "EffectiveMobile/docs"
	app "EffectiveMobile/internal"

	_ "github.com/joho/godotenv/autoload"
)

// @title Swagger Example API
// @version 1.0
// @description Song library
// @host localhost:8080
// @BasePath /
func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}
	app.Run(cfg)
}
