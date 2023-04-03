package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RecoveryMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func ()  {
			err := recover()
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H {"message": "internal server error"})
			}
		}()

		ctx.Next()
	}
}
