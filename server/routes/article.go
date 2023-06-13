package routes

import (
	"hallo-corona/handlers"
	"hallo-corona/pkg/middleware"
	"hallo-corona/pkg/mysql"
	"hallo-corona/repositories"

	"github.com/gin-gonic/gin"
)

func ArticleRoutes(r *gin.RouterGroup) {
	articleRepository := repositories.RepositoryArticle(mysql.DB)
	h := handlers.HandlerArticle(articleRepository)

	r.GET("/articles", h.FindArticles)
	r.GET("/article/:id", h.GetArticel)
	r.POST("/article", middleware.Auth(middleware.UploadFile(h.CreateArticle)))
	r.DELETE("/article/:id", middleware.Auth(h.DeleteArticle))
	// r.PATCH("/article/:id", middleware.Auth(middleware.UploadFile(h.UpdateArticle)))
	r.PATCH("/article/:id", middleware.Auth(middleware.UploadFile(h.UpdateArticle)))
	r.GET("/articles/:id", middleware.Auth(h.FindMyArticle))
}
