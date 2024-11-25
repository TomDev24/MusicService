package main

import (
	"log"
	"github.com/TomDev24/MusicServiceGo/internal/config"
	"github.com/TomDev24/MusicServiceGo/internal/server"
	"github.com/TomDev24/MusicServiceGo/internal/model"
	"github.com/TomDev24/MusicServiceGo/internal/database"
)

func main() {
	log.Println("Loading config...")
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Database initialization...")
	db, err := database.Init(cfg)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Database migration...")
    err = db.Conn.AutoMigrate(&model.Song{})
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Starting server...")
	serv := server.NewServer(cfg, db)
	serv.Start()
}
