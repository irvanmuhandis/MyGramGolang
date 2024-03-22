package middleware

import (
	"finalassignment/controllers/response"
	"finalassignment/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		verifyToken, err := helpers.VerifyToken(c)
		_ = verifyToken

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, response.ErrorMessage{
				Error:   "Unauthorized",
				Message: err.Error(),
			})
			return
		}
		//simpan token ke data request untuk diambil endpoint slanjutnya
		c.Set("userData", verifyToken)
		// Lanjut ke nedpoint berikutnya
		c.Next()
	}
}
