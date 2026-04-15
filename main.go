package main

import (
	"fmt"

	"id.diengs.backend/internal/config"
	"id.diengs.backend/internal/pkg"
)

func main() {
	viperConfig := config.NewViper()
	app := config.NewFiber(viperConfig)
	log := config.NewLogger(viperConfig)
	validate := config.NewValidator()
	mail := pkg.NewMail(viperConfig, log)
	db := config.NewDatabase(viperConfig, log)

	config.Bootstrap(&config.BootstrapConfig{
		App:      app,
		Log:      log,
		Validate: validate,
		Config:   viperConfig,
		Mail:     mail,
		DB:       db,
	})

	webPort := viperConfig.GetInt("web.port")
	if err := app.Listen(fmt.Sprintf(":%d", webPort)); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
