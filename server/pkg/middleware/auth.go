package middleware

import (
	jwtToken "hallo-corona/pkg/jwt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type Result struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

func Auth(next gin.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")

		if token == "" {
			c.JSON(http.StatusUnauthorized, Result{Code: http.StatusUnauthorized, Message: "unauthorized 1"})
			return
		}

		token = strings.Split(token, " ")[1]
		claims, err := jwtToken.DecodeToken(token)

		if err != nil {
			c.JSON(http.StatusUnauthorized, Result{Code: http.StatusUnauthorized, Message: "unauthorized 2"})
			return
		}

		c.Set("userLogin", claims)
		next(c)
	}
}
