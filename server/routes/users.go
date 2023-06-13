package routes

import (
	"hallo-corona/handlers"
	"hallo-corona/pkg/middleware"
	"hallo-corona/pkg/postgres"
	"hallo-corona/repositories"

	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.RouterGroup) {
	userRepository := repositories.RepositoryUser(postgres.DB)
	h := handlers.HandlerUser(userRepository)

	r.GET("/users", h.FindUsers)
	r.GET("/user/:id", h.GetUser)
	r.PATCH("/user", middleware.Auth(h.UpdateUser))
	r.DELETE("/user", middleware.Auth(h.DeleteUser))
	r.PATCH("/change-image/{id}", middleware.UploadFile(h.ChangeImage))
}
