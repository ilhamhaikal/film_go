package main

import (
	"log"

	"github.com/ilhamhaikal/film_go/config"
	"github.com/ilhamhaikal/film_go/internal/models"
	"github.com/ilhamhaikal/film_go/internal/server"
)

func main() {
	db, err := config.InitDB()
	if err != nil {
		log.Fatal(err)
	}

	// Auto migrate the schema
	err = db.AutoMigrate(&models.Cinema{})
	if err != nil {
		log.Fatal(err)
	}

	srv := server.NewServer(db)
	srv.SetupRoutes()

	log.Fatal(srv.Run(":8080"))
}
