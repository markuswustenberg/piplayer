package main

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"

	"piplayer/player"
	"piplayer/server"
	"piplayer/tagreader"
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

	tr := tagreader.New(tagreader.Options{
		Debug: false,
		Log:   logger,
	})
	if err := tr.Open(getStringOrDefault("TAG_READER_PATH", "/dev/input/event0")); err != nil {
		logger.Println("Error opening tag reader:", err)
		return 1
	}
	defer func() {
		if err := tr.Close(); err != nil {
			logger.Println("Error closing tag reader:", err)
		}
	}()

	go func() {
		for {
			id, err := tr.Read()
			if err != nil {
				logger.Println("Error reading tag:", err)
			}
			p.PlayFromID(id)
		}
	}()

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
