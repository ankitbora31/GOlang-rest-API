package controller

import (
	"errors"
	"net/http"
	"rest/config"
	"rest/entity"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AddressWeb struct {
	City  string `json:"city" binding:"required"`
	Pin   string `json:"pin" binding:"required"`
	State string `json:"state" binding:"required"`
}

func GetAddress(c *gin.Context) {
	db := config.InitDb()
	var address []entity.Address
	var user entity.User

	id, _ := c.Params.Get("userid")

	err1 := db.First(&user, id).Error

	if errors.Is(err1, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, gin.H{"message": "no user"})
		return
	}

	err := db.Where("user_id=?", id).First(&address).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, gin.H{"message": "no address"})
		return
	} else {
		db.Where("user_id=?", id).Find(&address)
		c.JSON(http.StatusOK, gin.H{"data": &address})
	}
}
func CreateAddress(c *gin.Context) {
	db := config.InitDb()

	id, _ := c.Params.Get("userid")
	var input AddressWeb
	var user entity.User
	err := db.First(&user, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, gin.H{"message": "no user"})
		return
	} else {
		err := c.ShouldBindJSON(&input)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		address := entity.Address{
			City:   input.City,
			Pin:    input.Pin,
			State:  input.State,
			UserID: id,
		}

		db.Create(&address)
		c.JSON(http.StatusOK, gin.H{"data": address})
	}
}

func UpdateAddress(c *gin.Context) {
	db := config.InitDb()
	userid, _ := c.Params.Get("userid")
	addressid, _ := c.Params.Get("addressid")

	var address entity.Address

	err := db.Where("user_id=?", userid).First(&address, addressid).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, gin.H{"not found": userid})
		return
	} else {
		c.BindJSON(&address)
		db.Where("id=?", addressid).Save(&address)
		c.JSON(http.StatusOK, address)
	}
}

func GetAddressById(c *gin.Context) {
	db := config.InitDb()
	userid, _ := c.Params.Get("userid")
	addressid, _ := c.Params.Get("addressid")
	var address entity.Address
	var user entity.User

	err1 := db.First(&user, userid).Error
	if errors.Is(err1, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, gin.H{"message": "no user"})
		return
	} else {
		err := db.Where("user_id=?", userid).First(&address, addressid).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"message": "no address"})
			return
		} else {
			c.JSON(http.StatusOK, gin.H{"data": &address})
		}
	}
}

func DeleteAddress(c *gin.Context) {
	db := config.InitDb()
	var user entity.User
	var address entity.Address
	userid, _ := c.Params.Get("userid")
	addressid, _ := c.Params.Get("addressid")

	err := db.First(&user, userid).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, gin.H{"not found": userid})
		return
	} else {
		err1 := db.Where("user_id=?", userid).First(&address, addressid).Error
		if errors.Is(err1, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"message": "no address"})
			return
		} else {
			db.Where("user_id=?", userid).Delete(&address, addressid)
			c.JSON(http.StatusOK, gin.H{"message": "address deleted"})
		}
	}
}
