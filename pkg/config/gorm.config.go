package config

import (
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func GormConfig() *gorm.Config {
	return &gorm.Config{
		PrepareStmt: true,
		Logger:      logger.Default,
	}
}
