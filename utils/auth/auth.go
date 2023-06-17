package auth

import (
	jwt_local "final-project3/utils/jwt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func MiddlewareLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		// validate jwt token
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.AbortWithStatus(401)
			return
		}

		claims, err := jwt_local.ParseToken(tokenString)
		if err != nil {
			c.AbortWithStatus(401)
			return
		}

		c.Set("user_info", claims)
		c.Next()
	}
}

func AuthorizationAccess() gin.HandlerFunc {
	return func(c *gin.Context) {
		// validate jwt token
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.AbortWithStatus(401)
			return
		}

		claims, err := jwt_local.ParseToken(tokenString)
		if err != nil {
			c.AbortWithStatus(401)
			return
		}

		if claims["role"] != "Admin" {
			c.AbortWithStatus(401)
			return
		}

		c.Set("user_info", claims)
		c.Next()
	}
}

func AuthenticationUserId() gin.HandlerFunc {
	return func(c *gin.Context) {
		// validate jwt token
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.AbortWithStatus(401)
			return
		}

		claims, err := jwt_local.ParseToken(tokenString)
		if err != nil {
			c.AbortWithStatus(401)
			return
		}

		idString := c.Param("id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			c.AbortWithStatus(400)
			return
		}

		if int(claims["id"].(float64)) != id {
			c.AbortWithStatus(401)
			return
		}

		c.Set("user_info", claims)
		c.Next()
	}
}
