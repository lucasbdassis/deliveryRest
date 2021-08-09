package database

import (
	"delivery/database/migrations"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"time"
)

var db *gorm.DB

func RunDb(){


	str := "host=localhost port=25432 user=admin dbname=books sslmode=disable password=123456"

	database , err := gorm.Open(postgres.Open(str), &gorm.Config{})
	if err != nil {
		log.Fatal("error: ",err)
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