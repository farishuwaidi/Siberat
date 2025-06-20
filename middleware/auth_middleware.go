package middleware

import (
	errorhandler "Siberat/errorHandler"
	"Siberat/helper"

	"github.com/gin-gonic/gin"
)

func JWTMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			errorhandler.HandlerError(c, &errorhandler.UnAuthorizedError{Message: "Unauthorized"})
			c.Abort()
			return
		}

		userId, err := helper.ValidateToken(tokenString)
		if err != nil {
			errorhandler.HandlerError(c,&errorhandler.UnAuthorizedError{Message: err.Error()})
			c.Abort()
			return
		}

		c.Set("userID", userId)
		c.Next()
	}
}