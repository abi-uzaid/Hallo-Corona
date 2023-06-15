package routes

import (
	"hallo-corona/handlers"
	"hallo-corona/pkg/middleware"
	"hallo-corona/pkg/mysql"
	"hallo-corona/repositories"

	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.RouterGroup) {
	userRepository := repositories.RepositoryUser(mysql.DB)
	h := handlers.HandlerUser(userRepository)

	r.GET("/users", h.FindUsers)
	r.GET("/user/:id", h.GetUser)
	r.PATCH("/user", middleware.Auth(middleware.UploadFile(h.UpdateUser)))
	r.DELETE("/user", middleware.Auth(h.DeleteUser))
	r.PATCH("/change-image/:id", middleware.Auth(middleware.UploadFile(h.ChangeImage)))
}
