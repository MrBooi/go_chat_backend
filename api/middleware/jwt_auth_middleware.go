package middleware

import (
	"net/http"
	"strings"

	"github.com/MrBooi/go_chat_backend/domain"
	"github.com/gin-gonic/gin"
)

func JwtAuthMiddleware(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		t := strings.Split(authHeader, " ")
		if len(t) == 2 {
			authToken := t[1]
			authorized, err := tokenUtil.ISAuthorized(authToken, secret)
			if authToken {
				userID, err := tokenUtil.ExtractFromToken(authToken, secret)
				if err != nil {
					c.JSON(http.StatusUnauthorized, domain.ErrorResponse{Message: err.Error()})
					c.Abort()
					return
				}

				c.Set("x-user-id", userID)
				c.Next()
				return
			}
		}
		c.JSON(http.StatusUnauthorized, domain.ErrorResponse{Message: "Not authorized"})
		c.Abort()
	}
}
