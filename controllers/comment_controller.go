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

func GetAllComment(c *gin.Context) {
	db := database.GetDB()

	comments := []models.Comment{}
	photoId, _ := strconv.Atoi(c.Param("photoId"))

	if err := db.Find(&comments).Where("photo_id = ?", photoId).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Internal Server Error",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, comments)
}

func GetCommentById(c *gin.Context) {
	db := database.GetDB()

	comment := models.Comment{}

	commentId, _ := strconv.Atoi(c.Param("commentId"))
	photoId, _ := strconv.Atoi(c.Param("photoId"))

	if err := db.Find(&comment).Where("photo_id = ?", photoId).Where("id = ?", commentId).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Internal Server Error",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, comment)
}

func CreateComment(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	comment := models.Comment{}
	contentType := helpers.GetContentType(c)

	userId := uint(userData["id"].(float64))
	photoId, _ := strconv.Atoi(c.Param("photoId"))

	var err error
	if contentType == appJSON {
		err = c.ShouldBindBodyWithJSON(&comment)
	} else {
		err = c.ShouldBind(&comment)
	}

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	comment.UserId = userId
	comment.PhotoId = uint(photoId)

	if err = db.Debug().Create(&comment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Internal Server Error",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, comment)
}

func UpdateComment(c *gin.Context) {
	db := database.GetDB()
	comment := models.Comment{}
	commentId, _ := strconv.Atoi(c.Param("commentId"))
	contentType := helpers.GetContentType(c)

	userData := c.MustGet("userData").(jwt.MapClaims)
	photoId, _ := strconv.Atoi(c.Param("photoId"))
	userId := uint(userData["id"].(float64))

	var err error
	if contentType == appJSON {
		err = c.ShouldBindBodyWithJSON(&comment)
	} else {
		err = c.ShouldBind(&comment)
	}

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	comment.UserId = userId
	comment.PhotoId = uint(photoId)
	comment.ID = uint(commentId)

	dataToUpdate := models.Comment{
		Message: comment.Message,
	}

	if err = db.Model(&comment).Where("id = ?", commentId).Where("photo_id = ?", photoId).Updates(dataToUpdate).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Internal Server Error",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, comment)
}

func DeleteComment(c *gin.Context) {
	db := database.GetDB()
	comment := models.Comment{}
	photoId, _ := strconv.Atoi(c.Param("photoId"))
	commentId, _ := strconv.Atoi(c.Param("commentId"))
	if err := db.Where("id = ?", commentId).Model(&comment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Internal Server Error",
			"message": err.Error(),
		})
		return
	}

	comment.ID = uint(commentId)
	comment.PhotoId = uint(photoId)

	if err := db.Where("id = ?", commentId).Delete(&comment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Internal Server Error",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, "Data Berhasil Dihapus")
}
