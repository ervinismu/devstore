package middleware

import (
	"net/http"
	"strings"

	"github.com/ervinismu/devstore/internal/pkg/handler"
	"github.com/ervinismu/devstore/internal/pkg/reason"
	"github.com/gin-gonic/gin"
)

type AccessTokenVerifier interface {
	VerifyAccessTokenToken(tokenString string) (string, error)
}

func AuthMiddleware(tokenMaker AccessTokenVerifier) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		accessToken := tokenFromHeader(ctx)
		if accessToken == "" {
			handler.ResponseError(ctx, http.StatusUnauthorized, reason.Unauthorized)
			ctx.Abort()
			return
		}
		sub, err := tokenMaker.VerifyAccessTokenToken(accessToken)
		if err != nil {
			handler.ResponseError(ctx, http.StatusUnauthorized, reason.Unauthorized)
			ctx.Abort()
			return
		}

		// attach to request
		ctx.Set("user_id", sub)

		// continue
		ctx.Next()
	}
}

func tokenFromHeader(ctx *gin.Context) string {
	var accessToken string

	bearerToken := ctx.Request.Header.Get("Authorization")
	fields := strings.Fields(bearerToken)

	if len(fields) != 0 && fields[0] == "Bearer" {
		accessToken = fields[1]
	}

	return accessToken
}
