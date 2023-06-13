package routes

import (
	"hallo-corona/handlers"
	"hallo-corona/pkg/middleware"
	"hallo-corona/pkg/mysql"
	"hallo-corona/repositories"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.RouterGroup) {
	authRepository := repositories.RepositoryAuth(mysql.DB)
	h := handlers.HandlerAuth(authRepository)

	r.POST("/register", h.Register)
	r.POST("/login", h.Login)
	r.GET("/check-auth", middleware.Auth(h.CheckAuth))
}
