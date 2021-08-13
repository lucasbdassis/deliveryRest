package controllers

import (
	"delivery/database"
	"delivery/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm/utils"
	"net/http"
)

func LoginPost(c *gin.Context) {
	var (
		json   models.Login
		jsonDb models.Login
	)

	db := database.GetDatabase()
	err := c.ShouldBindJSON(&json)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		return
	}
	err = db.Find(&jsonDb, &json).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "User is not already registered: " + err.Error(),
		})
		return
	}
	fmt.Println("123", err)

	if !utils.AssertEqual(jsonDb.User, json.User) || !utils.AssertEqual(jsonDb.Password, json.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
}

func Register(c *gin.Context) {
	db := database.GetDatabase()

	var requestRegister models.Login

	err := c.ShouldBindJSON(&requestRegister)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		return
	}

	err = db.Create(&requestRegister).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot create a user: " + err.Error(),
		})
		return
	}

	c.JSON(201, requestRegister)
}
