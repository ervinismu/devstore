package middleware

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func PaginationMiddleware(defaultPage int, defaultPageSize int) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		page, err := strconv.Atoi(ctx.Query("page"))
		if err != nil {
			page = defaultPage
		}

		pageSize, err := strconv.Atoi(ctx.Query("page_size"))
		if err != nil {
			pageSize = defaultPageSize
		}

		ctx.Set("page", page)
		ctx.Set("page_size", pageSize)

		ctx.Next()
	}

}
