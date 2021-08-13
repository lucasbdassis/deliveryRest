package migrations

import (
	"delivery/models"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	// ele da um migrate com o banco de dados
	db.AutoMigrate(&models.Product{},&models.Login{})

}
