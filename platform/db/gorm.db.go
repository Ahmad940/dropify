package db

import (
	"fmt"
	"os"

	"github.com/Ahmad940/dropify/app/model"
	"github.com/Ahmad940/dropify/pkg/config"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitializeMigration() {
	var err error
	DB, err = gorm.Open(sqlite.Open("app.db"), config.GormConfig())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	// migrations here
	err = DB.AutoMigrate(&model.User{})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Migration has failed: %v\n", err)
		os.Exit(1)
	}
}
