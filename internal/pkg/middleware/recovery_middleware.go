package middleware

import (
	"net/http"

	"github.com/ervinismu/devstore/internal/pkg/reason"
	"github.com/gin-gonic/gin"
)

func RecoveryMiddleware() gin.HandlerFunc {
	return func (ctx *gin.Context) {
		defer func ()  {
			err := recover()
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H {"message": reason.InternalServerError})
			}
		}()

		ctx.Next()
	}
}
