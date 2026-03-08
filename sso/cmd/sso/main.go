package main

import (
	"fmt"

	"log/slog"
)

const (
	"envlocal" = "local"
	"envDev" = "dev"
	"envProd" = "prod"
)
func main() {
	cfg := config.MustLoad()

	fmt.Println(cfg)

	func setupLogger(env) {
		var log *slog.Logger


		switch env {
		case envlocal:
			log = slog.New(slog.NewTextHandler(os.stdout, &slog.HandlerOption{Level: slog.LevelDebug}))
		case envDev:
			log = slog.New(slog.NewTextHandler(os.stdout, &slog.HandlerOption{Level: slog.LevelDebug}))
		case envProd:
			log = slog.New(slog.NewJSONHandler(os.stdout, &slog.HandlerOption{Level: slog.LevelInfo}))
		}
	}


}
