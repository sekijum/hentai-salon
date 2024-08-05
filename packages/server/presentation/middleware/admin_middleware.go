package middleware

import (
	"net/http"
	"server/domain/model"

	"github.com/gin-gonic/gin"
)

func AdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get("role")
		if !exists {
			c.JSON(http.StatusForbidden, gin.H{"error": "アクセスが拒否されました: 権限情報が見つかりません"})
			c.Abort()
			return
		}

		if role != int(model.UserRoleAdmin) {
			c.JSON(http.StatusForbidden, gin.H{"error": "アクセスが拒否されました: 管理者のみアクセス可能です"})
			c.Abort()
			return
		}

		c.Next()
	}
}
