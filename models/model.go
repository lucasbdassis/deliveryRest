package models

import (
	"gorm.io/gorm"
	"time"
)

// https://gorm.io/docs/models.html

type Product struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	Name        string         `json:"first_name"`
	Size        string         `json:"size"`
	Description string         `json:"description"`
	MediumPrice float32        `json:"medium_price"`
	Pasta       Pasta          `json:"pasta" gorm:"-"`
	ImageURL    string         `json:"img_url"`
	CreatedAt   time.Time      `json:"created"`
	UpdatedAt   time.Time      `json:"updated"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted"`
}

type Pasta struct {
	Traditional []string `json:"traditional"`
}
