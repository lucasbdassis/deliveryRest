package controllers

import (
	"delivery/database"
	"delivery/models"
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetProduct(c *gin.Context) {
	// Irei pegar o id escrito na URL
	id := c.Param("id")

	// Converto o id para inteiro , e retorno erro se tiver algo fora do padrao
	newId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Id has to be integer",
		})

		return
	}
	// crio a conexao com o banco
	db := database.GetDatabase()

	var product models.Product
	// Faço uma busca no banco com o id, e retorno na struct com os devidos campos
	err = db.First(&product, newId).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot find product",
		})

		return
	}

	// retorna na pagina com um json da struct
	c.JSON(200, product)
}

func CreateProduct(c *gin.Context) {
	db := database.GetDatabase()

	var p models.Product

	err := c.ShouldBindJSON(&p)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		return
	}

	err = db.Create(&p).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot create product: " + err.Error(),
		})
		return
	}

	c.JSON(201, p)
}

func UpdateProduct(c *gin.Context) {
	db := database.GetDatabase()

	var p models.Product

	err := c.ShouldBindJSON(&p)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		return
	}

	err = db.Save(&p).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot update product: " + err.Error(),
		})
		return
	}

	c.JSON(200, p)
}

func DeleteProduct(c *gin.Context) {
	id := c.Param("id")

	newId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Id has to be integer",
		})

		return
	}

	db := database.GetDatabase()

	err = db.First(&models.Product{}, newId).Error
	if err != nil {
		c.JSON(204, gin.H{
			"error": "cannot delete this id , is because nonexistent : " + err.Error(),
		})
		return
	}
	err = db.Delete(&models.Product{}, &newId).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot delete product: " + err.Error(),
		})
		return
	}

	c.Status(200)
}

func GetAllProducts(c *gin.Context) {
	db := database.GetDatabase()

	var products []models.Product
	err := db.Find(&products).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot list products: " + err.Error(),
		})

		return
	}
	if len(products) == 0 {
		c.Status(204)

		return
	}

	c.JSON(200, products)
}
