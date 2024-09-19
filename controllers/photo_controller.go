package controllers

import (
	"MyGram/database"
	"MyGram/helpers"
	"MyGram/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func GetAllPhoto(c *gin.Context) {
	db := database.GetDB()

	photos := []models.Photo{}

	if err := db.Find(&photos).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Internal Server Error",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, photos)
}

func GetPhotoById(c *gin.Context) {
	db := database.GetDB()

	photoId, _ := strconv.Atoi(c.Param("photoId"))
	photo := models.Photo{}

	if err := db.First(&photo, "id = ?", photoId).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Internal Server Error",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, photo)
}

func CreatePhoto(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	photo := models.Photo{}
	contentType := helpers.GetContentType(c)

	userId := uint(userData["id"].(float64))

	var err error
	if contentType == appJSON {
		err = c.ShouldBindBodyWithJSON(&photo)
	} else {
		err = c.ShouldBind(&photo)
	}

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	photo.UserId = userId

	if err = db.Debug().Create(&photo).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Internal Server Error",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, photo)
}

func UpdatePhoto(c *gin.Context) {
	db := database.GetDB()
	photo := models.Photo{}
	photoId, _ := strconv.Atoi(c.Param("photoId"))
	contentType := helpers.GetContentType(c)
	userData := c.MustGet("userData").(jwt.MapClaims)
	userId := uint(userData["id"].(float64))

	var err error
	if contentType == appJSON {
		err = c.ShouldBindBodyWithJSON(&photo)
	} else {
		err = c.ShouldBind(&photo)
	}

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	photo.UserId = userId
	photo.ID = uint(photoId)

	dataToUpdate := models.Photo{
		Title:    photo.Title,
		Caption:  photo.Caption,
		PhotoUrl: photo.PhotoUrl,
	}

	if err = db.Model(&photo).Where("id = ?", photoId).Updates(dataToUpdate).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Internal Server Error",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, photo)
}

func DeletePhoto(c *gin.Context) {
	db := database.GetDB()
	photo := models.Photo{}
	photoId, _ := strconv.Atoi(c.Param("photoId"))
	photo.ID = uint(photoId)

	if err := db.Where("id = ?", photoId).Delete(&photo).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Internal Server Error",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, "Foto Berhasil Dihapus")
}
