package main

import (
	"api-service/internal/server"
	"api-service/internal/store"
	"log"
	"os"
)

type Config struct {
	HttpPort    string
	HttpTimeout int
	StoreInit   string
}

type App struct {
	PostStore store.PostStore
	WebServer server.WebServer
}

func main() {
	// get service logger and configuration
	logger := getLogger()
	config := getConfig()

	logger.Println("starting API service")

	if err := do(config, logger); err != nil {
		logger.Println("finishing service", err)
		os.Exit(1)
	}

	logger.Println("finishing API service")
}

func getLogger() *log.Logger {
	return log.Default()
}

func getConfig() *Config {
	config := Config{
		HttpPort:    os.Getenv("HTTP_PORT"),
		HttpTimeout: 1, // seconds
		StoreInit:   os.Getenv("STORE_INIT"),
	}
	return &config
}

func do(config *Config, logger *log.Logger) error {
	// initialize application
	logger.Println("initializing application")
	memoryStore, err := store.NewMemoryPostStore(config.StoreInit)
	if err != nil {
		return err
	}

	app := App{
		PostStore: memoryStore,
		WebServer: server.NewWebServer(config.HttpPort),
	}

	// start HTTP server
	logger.Printf("starting http server on port %s\n", config.HttpPort)
	err = app.WebServer.Serve(app.routes())
	if err != nil {
		return err
	}

	return nil
}
