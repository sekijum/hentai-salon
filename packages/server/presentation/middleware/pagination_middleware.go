package middleware

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func PaginationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defaultLimit := 20
		defaultOffset := 0

		limitStr := c.Query("limit")
		offsetStr := c.Query("offset")

		limit, err := strconv.Atoi(limitStr)
		if err != nil || limit <= 0 {
			limit = defaultLimit
		}

		offset, err := strconv.Atoi(offsetStr)
		if err != nil || offset < 0 {
			offset = defaultOffset
		}

		c.Set("limit", limit)
		c.Set("offset", offset)

		c.Next()
	}
}
