package controllers

import (
	"delivery/database"
	"delivery/models"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm/utils"
	"net/http"
)

func LoginGet(c *gin.Context) {

	session := sessions.Default(c)
	fmt.Println("session: ", session)

	v := session.Get("count")
	if v == nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "NonAuthorized"})
		return

	}
	fmt.Printf("%T", v)

	c.JSON(200, gin.H{"Status": "Authorized"})

}

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

	if !utils.AssertEqual(jsonDb.User, json.User) || !utils.AssertEqual(jsonDb.Password, json.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
		return
	}
	session := sessions.Default(c)
	session.Set("count", json.User)
	session.Save()

	c.JSON(http.StatusOK, gin.H{"code : ": "Welcome " + jsonDb.User + " you're connected"})
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
