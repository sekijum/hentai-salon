package middleware

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func PaginationDefaultsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		const defaultLimit = 20
		const defaultOffset = 0

		limitStr := c.DefaultQuery("limit", strconv.Itoa(defaultLimit))
		offsetStr := c.DefaultQuery("offset", strconv.Itoa(defaultOffset))

		limit, err := strconv.Atoi(limitStr)
		if err != nil || limit <= 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit parameter"})
			c.Abort()
			return
		}

		offset, err := strconv.Atoi(offsetStr)
		if err != nil || offset < 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid offset parameter"})
			c.Abort()
			return
		}

		c.Set("limit", limit)
		c.Set("offset", offset)

		c.Next()
	}
}
