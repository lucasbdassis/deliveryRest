package database

import (
	"delivery/database/migrations"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"time"
)

var db *gorm.DB

func RunDb(x string) {

	database, err := gorm.Open(postgres.Open(x), &gorm.Config{})
	if err != nil {
		log.Fatal("error: ", err)
	}

	db = database

	config, err := db.DB()
	if err != nil {
		log.Fatal("Erro ao se conectar ao banco")
	}

	config.SetMaxIdleConns(10)
	config.SetMaxOpenConns(100)
	config.SetConnMaxLifetime(time.Hour)
	migrations.RunMigrations(db)

}

func GetDatabase() *gorm.DB {
	return db
}
