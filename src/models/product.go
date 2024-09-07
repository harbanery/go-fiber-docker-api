package models

import (
	"go-fiber-docker-api/src/config"
	"time"
)

type Product struct {
	ID        uint       `json:"id" gorm:"primarykey"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at" gorm:"index"`
	Name      string     `json:"name"`
	Price     float64    `json:"price"`
	Stock     int        `json:"stock"`
}

func GetAllProducts() []*Product {
	var products []*Product
	config.DB.Find(&products)
	return products
}

func GetProductbyID(id int) *Product {
	var product Product
	config.DB.First(&product, "id = ?", id)
	return &product
}

func CreateProduct(product *Product) error {
	result := config.DB.Create(&product)
	return result.Error
}

func UpdateProduct(id int, product *Product) error {
	result := config.DB.Model(&Product{}).Where("id = ?", id).Updates(product)
	return result.Error
}

func DeleteProduct(id int) error {
	result := config.DB.Delete(&Product{}, "id = ?", id)
	return result.Error
}
