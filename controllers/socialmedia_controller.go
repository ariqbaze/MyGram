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

func GetAllSocialMedia(c *gin.Context) {
	db := database.GetDB()

	socialMedia := []models.SocialMedia{}

	if err := db.Find(&socialMedia).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Internal Server Error",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, socialMedia)
}

func GetSocialMediaById(c *gin.Context) {
	db := database.GetDB()

	socialMedia := models.SocialMedia{}

	socialMediaId, _ := strconv.Atoi(c.Param("socialMediaId"))

	if err := db.Find(&socialMedia).Where("id = ?", socialMediaId).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Internal Server Error",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, socialMedia)
}

func CreateSocialMedia(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	socialMedia := models.SocialMedia{}
	contentType := helpers.GetContentType(c)

	userId := uint(userData["id"].(float64))
	socialMediaId, _ := strconv.Atoi(c.Param("socialMediaId"))

	var err error
	if contentType == appJSON {
		err = c.ShouldBindBodyWithJSON(&socialMedia)
	} else {
		err = c.ShouldBind(&socialMedia)
	}

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	socialMedia.UserId = userId
	socialMedia.ID = uint(socialMediaId)

	if err = db.Debug().Create(&socialMedia).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Internal Server Error",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, socialMedia)
}

func UpdateSocialMedia(c *gin.Context) {
	db := database.GetDB()
	socialMedia := models.SocialMedia{}
	socialMediaId, _ := strconv.Atoi(c.Param("socialMediaId"))
	contentType := helpers.GetContentType(c)

	userData := c.MustGet("userData").(jwt.MapClaims)
	userId := uint(userData["id"].(float64))

	var err error
	if contentType == appJSON {
		err = c.ShouldBindBodyWithJSON(&socialMedia)
	} else {
		err = c.ShouldBind(&socialMedia)
	}

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	socialMedia.UserId = userId
	socialMedia.ID = uint(socialMediaId)

	dataToUpdate := models.SocialMedia{
		Name:           socialMedia.Name,
		SocialMediaUrl: socialMedia.SocialMediaUrl,
	}

	if err = db.Model(&socialMedia).Where("id = ?", socialMediaId).Updates(dataToUpdate).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Internal Server Error",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, socialMedia)
}

func DeleteSocialMedia(c *gin.Context) {
	db := database.GetDB()
	socialMedia := models.SocialMedia{}
	socialMediaId, _ := strconv.Atoi(c.Param("socialMediaId"))
	if err := db.Where("id = ?", socialMediaId).Model(&socialMedia).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Internal Server Error",
			"message": err.Error(),
		})
		return
	}

	socialMedia.ID = uint(socialMediaId)

	if err := db.Where("id = ?", socialMediaId).Delete(&socialMedia).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Internal Server Error",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, "Data Berhasil Dihapus")
}
