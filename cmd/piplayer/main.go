package main

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"

	"piplayer/player"
	"piplayer/server"
)

func main() {
	os.Exit(start())
}

func start() int {
	_ = godotenv.Load()

	logger := log.New(os.Stdout, "", 0)

	p := player.New(player.Options{
		Dir:      getStringOrDefault("PLAYER_DIR", "."),
		Host:     getStringOrDefault("PLAYER_HOST", "localhost"),
		Log:      logger,
		Password: getStringOrDefault("PLAYER_PASSWORD", ""),
		Port:     getIntOrDefault("PLAYER_PORT", 8080),
	})

	s := server.New(server.Options{
		Host:   getStringOrDefault("HOST", ""),
		Log:    logger,
		Player: p,
		Port:   getIntOrDefault("PORT", 8081),
	})

	if err := s.Start(); err != nil {
		logger.Println("Error starting:", err)
		return 1
	}
	return 0
}

func getStringOrDefault(name, defaultV string) string {
	v, ok := os.LookupEnv(name)
	if !ok {
		return defaultV
	}
	return v
}

func getIntOrDefault(name string, defaultV int) int {
	v, ok := os.LookupEnv(name)
	if !ok {
		return defaultV
	}
	vAsInt, err := strconv.Atoi(v)
	if err != nil {
		return defaultV
	}
	return vAsInt
}
