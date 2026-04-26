package main

import (
	"fmt"
	"log"
	"log/slog"

	"github.com/student-assistant/internal/api"
)

func main() {
	s := api.New()
	address, port := "0.0.0.0", 8080
	slog.Info("Starting server", slog.String("address", fmt.Sprintf("%s:%d", address, port)))
	log.Fatal(s.Run(address, port))
}
