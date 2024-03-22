package route

import (
	controller "finalassignment/controllers"
	"finalassignment/middleware"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func StartApp() *gin.Engine {
	r := gin.Default()
	userRouter := r.Group("/users")
	{
		userRouter.POST("/register", controller.UserRegis)
		userRouter.POST("/login", controller.UserLogin)
		userRouter.DELETE("/:userId", middleware.Authentication(), controller.DeleteUser)
		userRouter.PUT("/:userId", middleware.Authentication(), controller.UpdateUser)
	}

	photoRouter := r.Group("/photos")
	{
		photoRouter.Use(middleware.Authentication())

		photoRouter.GET("/", controller.ShowPhoto)
		photoRouter.POST("/", controller.CreatePhoto)
		photoRouter.PUT("/:photoId", middleware.PhotoAuthorization(), controller.UpdatePhoto)
		photoRouter.DELETE("/:photoId", middleware.PhotoAuthorization(), controller.DeletePhoto)
	}

	commentRouter := r.Group("/comments")
	{
		commentRouter.Use(middleware.Authentication())

		commentRouter.GET("/", controller.ShowComment)
		commentRouter.POST("/", controller.CreateComment)
		commentRouter.PUT("/:commentId", middleware.CommentAuthorization(), controller.UpdateComment)
		commentRouter.DELETE("/:commentId", middleware.CommentAuthorization(), controller.DeleteComment)
	}

	socialMediaRouter := r.Group("/socialmedias")
	{
		socialMediaRouter.Use(middleware.Authentication())

		socialMediaRouter.GET("/", controller.ShowSocialMedia)
		socialMediaRouter.POST("/", controller.CreateSocialMedia)
		socialMediaRouter.PUT("/:socialMediaId", middleware.SocialAuthorization(), controller.UpdateSocialMedia)
		socialMediaRouter.DELETE("/:socialMediaId", middleware.SocialAuthorization(), controller.DeleteSocialMedia)
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
