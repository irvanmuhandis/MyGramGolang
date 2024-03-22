package middleware

import (
	"finalassignment/database"
	"finalassignment/models"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func PhotoAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.GetDB()

		photoId, err := strconv.Atoi(c.Param("photoId"))

		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":    "Bad Request",
				"messsage": "Invalid parameter",
			})
			return
		}
		userData := c.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))
		photo := models.Photos{}

		err = db.Select("user_id").First(&photo, uint(photoId)).Error
		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":    "Not Found",
				"messsage": "Photo not found",
			})
			return
		}

		if photo.UserID != userID {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":    "Unauthorized",
				"messsage": "You didnt have acces",
			})
			return
		}

		// Lanjut ke nedpoint berikutnya
		c.Next()
	}
}

func CommentAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.GetDB()

		commentId, err := strconv.Atoi(c.Param("commentId"))

		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":    "Bad Request",
				"messsage": "Invalid parameter",
			})
			return
		}
		userData := c.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))
		comment := models.Comments{}

		err = db.Select("user_id").First(&comment, uint(commentId)).Error
		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":    "Not Found",
				"messsage": "Comment not found",
			})
			return
		}

		if comment.UserID != userID {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":    "Unauthorized",
				"messsage": "You didnt have acces",
			})
			return
		}

		// Lanjut ke nedpoint berikutnya
		c.Next()
	}
}

func SocialAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.GetDB()

		socialId, err := strconv.Atoi(c.Param("socialMediaId"))

		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":    "Bad Request",
				"messsage": "Invalid parameter",
			})
			return
		}
		userData := c.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))
		social := models.SocialMedias{}

		err = db.Select("user_id").First(&social, uint(socialId)).Error
		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":    "Not Found",
				"messsage": "Social media not found",
			})
			return
		}

		if social.UserId != userID {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":    "Unauthorized",
				"messsage": "You didnt have acces",
			})
			return
		}

		// Lanjut ke nedpoint berikutnya
		c.Next()
	}
}
