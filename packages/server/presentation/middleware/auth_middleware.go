package middleware

import (
	"errors"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "認証トークンが必要です"})
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		secretKey := os.Getenv("JWT_SECRET_KEY")
		if secretKey == "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "サーバーエラー: 秘密鍵が設定されていません"})
			c.Abort()
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, gin.Error{
					Err:  errors.New("予期しない署名方法です"),
					Type: gin.ErrorTypePublic,
				}
			}
			return []byte(secretKey), nil
		})

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "認証に失敗しました: " + err.Error()})
			c.Abort()
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			userID := int(claims["userID"].(float64))
			role := int(claims["role"].(float64))
			c.Set("role", role)
			c.Set("userID", userID)
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "トークンが無効です"})
			c.Abort()
			return
		}

		c.Next()
	}
}
