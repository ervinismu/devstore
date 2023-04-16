package middleware

import (
	"net/http"
	"strings"

	"github.com/ervinismu/devstore/internal/pkg/handler"
	"github.com/ervinismu/devstore/internal/pkg/reason"
	"github.com/ervinismu/devstore/internal/pkg/token"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := tokenFromHeader(ctx)
		if tokenString == "" {
			handler.ResponseError(ctx, http.StatusUnauthorized, reason.Unauthorized)
			ctx.Abort()
			return
		}
		err := token.VerifyToken(tokenString)
		if err != nil {
			handler.ResponseError(ctx, http.StatusUnauthorized, reason.Unauthorized)
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}

func tokenFromHeader(ctx *gin.Context) string {
	bearerToken := ctx.Request.Header.Get("Authorization")
	splitToken := strings.Split(bearerToken, " ")
	if len(splitToken) == 2 {
		return splitToken[1]
	}

	return ""
}
