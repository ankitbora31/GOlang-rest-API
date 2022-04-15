package controller

import (
	"errors"
	"net/http"
	"rest/config"
	"rest/entity"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserWeb struct {
	Name   string `json:"name"`
	Gender string `json:"gender"`
	Age    int    `json:"age"`
}

func GetUsers(c *gin.Context) {
	db := config.InitDb()
	var users []entity.User

	db.Find(&users)
	c.JSON(http.StatusOK, gin.H{"data": users})
}

func CreateUser(c *gin.Context) {
	db := config.InitDb()

	var input UserWeb

	err := c.ShouldBindJSON(&input)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := entity.User{
		Name:   input.Name,
		Gender: input.Gender,
		Age:    input.Age,
	}

	db.Create(&user)
	c.JSON(http.StatusOK, gin.H{"data": user})
}

func UpdateUser(c *gin.Context) {
	db := config.InitDb()
	userid, _ := c.Params.Get("userid")

	var user entity.User

	err := db.First(&user, userid).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, gin.H{"not found": userid})
		return
	} else {
		c.BindJSON(&user)
		db.Save(&user)

		c.JSON(http.StatusOK, user)
	}
}

func GetUserById(c *gin.Context) {
	db := config.InitDb()
	userid, _ := c.Params.Get("userid")
	var user entity.User
	err := db.First(&user, userid).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, gin.H{"not found": userid})
		return
	} else {
		c.JSON(http.StatusOK, &user)
	}
}

func DeleteUser(c *gin.Context) {
	db := config.InitDb()
	var user entity.User
	var address entity.Address
	userid, _ := c.Params.Get("userid")

	dbResult := db.First(&user, userid).Error
	if errors.Is(dbResult, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, gin.H{"not found": userid})
		return
	} else {
		db.Where("user_id=?", userid).Delete(&address)
		db.Where("id=?", userid).Delete(&user)
		c.JSON(http.StatusOK, gin.H{"message": "user deleted"})
	}
}
