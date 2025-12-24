package middleware

import (
	"fmt"
	"knowtime/config"
	"knowtime/internal"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = config.JwtKey

type Claims struct {
	UserID uint `json:"user_id"`
	jwt.RegisteredClaims
}

// GenerateJWT 生成JWT Token
func GenerateJWT(userID uint) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    config.JwtIssuer,
			Subject:   "user_token",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

// JWTAuthMiddleware JWT验证中间件
func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, internal.BaseMsg{
				Code:    401,
				Message: "Authorization header is required",
			})
			c.Abort()
			return
		}

		parts := strings.Fields(authHeader)
		if len(parts) != 2 || !strings.EqualFold(parts[0], "Bearer") {
			c.JSON(http.StatusUnauthorized, internal.BaseMsg{
				Code:    401,
				Message: "Authorization header is required",
			})
			c.Abort()
			return
		}
		tokenString := parts[1]

		claims := &Claims{}
		// 限制有效算法
		parser := jwt.NewParser(jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))
		token, err := parser.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok || token.Method.Alg() != jwt.SigningMethodHS256.Alg() {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return jwtKey, nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, internal.BaseMsg{
				Code:    401,
				Message: "Invalid or expired token",
			})
			c.Abort()
			return
		}

		if claims.Subject != "user_token" {
			c.JSON(http.StatusUnauthorized, internal.BaseMsg{
				Code:    401,
				Message: "Invalid token subject",
			})
			c.Abort()
			return
		}

		if config.JwtIssuer != "" && claims.Issuer != config.JwtIssuer {
			c.JSON(http.StatusUnauthorized, internal.BaseMsg{
				Code:    401,
				Message: "Invalid token issuer",
			})
			c.Abort()
			return
		}

		// 将用户ID存储到上下文中
		c.Set("user_id", claims.UserID)
		c.Next()
	}
}
