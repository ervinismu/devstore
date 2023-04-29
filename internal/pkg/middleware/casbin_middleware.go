package middleware

import (
	"net/http"

	"github.com/casbin/casbin/v2"
	"github.com/ervinismu/devstore/internal/pkg/handler"
	"github.com/gin-gonic/gin"
)

func AuthorizationMiddleware(sub string, obj string, act string, enforcer *casbin.Enforcer) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// sub the user that wants to access a resource.
		// obj the resource that is going to be accessed.
		// act the operation that the user performs on the resource.

		res, _ := enforcer.Enforce(sub, obj, act)
		if res {
			// permit alice to read data1
			ctx.Next()
		} else {
			// deny the request, show an error
			handler.ResponseError(ctx, http.StatusUnauthorized, "you are not authorized to perform this action.")
			ctx.Abort()
			return
		}
	}
}
