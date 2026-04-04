package main

import (
	"fmt"

	"id.diengs.backend/internal/config"
)

func main() {
	viperConfig := config.NewViper()
	app := config.NewFiber(viperConfig)
	log := config.NewLogger(viperConfig)
	validate := config.NewValidator()

	config.Bootstrap(&config.BootstrapConfig{
		App:      app,
		Log:      log,
		Validate: validate,
		Config:   viperConfig,
	})

	webPort := viperConfig.GetInt("web.port")
	if err := app.Listen(fmt.Sprintf(":%d", webPort)); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
