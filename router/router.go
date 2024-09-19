package router

import (
	"MyGram/controllers"
	"MyGram/middlewares"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	r := gin.Default()

	userRouter := r.Group("/users")

	userRouter.POST("/register", controllers.UserRegister)
	userRouter.POST("/login", controllers.UserLogin)

	photoRouter := r.Group("/photo")

	photoRouter.Use(middlewares.Authentication())
	photoRouter.GET("/", controllers.GetAllPhoto)
	photoRouter.POST("/", controllers.CreatePhoto)
	photoRouter.GET("/:photoId", controllers.GetPhotoById)
	photoRouter.PUT("/:photoId", middlewares.PhotoAuthorization(), controllers.UpdatePhoto)
	photoRouter.DELETE("/:photoId", middlewares.PhotoAuthorization(), controllers.DeletePhoto)

	commentRouter := r.Group("/photo/:photoId/comment")

	commentRouter.Use(middlewares.Authentication())
	commentRouter.GET("/", controllers.GetAllComment)
	commentRouter.POST("/", controllers.CreateComment)
	commentRouter.GET("/:commentId", controllers.GetCommentById)
	commentRouter.PUT("/:commentId", middlewares.CommentAuthorization(), controllers.UpdateComment)
	commentRouter.DELETE("/:commentId", middlewares.CommentAuthorization(), controllers.DeleteComment)

	socialMediaRouter := r.Group("/socialmedia")

	socialMediaRouter.Use(middlewares.Authentication())
	socialMediaRouter.GET("/", controllers.GetAllSocialMedia)
	socialMediaRouter.POST("/", controllers.CreateSocialMedia)
	socialMediaRouter.GET("/:socialMediaId", controllers.GetSocialMediaById)
	socialMediaRouter.PUT("/:socialMediaId", middlewares.SocialMediaAuthorization(), controllers.UpdateSocialMedia)
	socialMediaRouter.DELETE("/:socialMediaId", middlewares.SocialMediaAuthorization(), controllers.DeleteSocialMedia)

	return r
}
