package main

import (
	"github.com/Ahmad940/dropify/platform/db"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	// connecting to database and initialize migrations
	db.InitializeMigration()

	// starting app
	StartApp()

}
