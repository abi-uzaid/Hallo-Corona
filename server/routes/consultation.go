package routes

import (
	"hallo-corona/handlers"
	"hallo-corona/pkg/middleware"
	"hallo-corona/pkg/postgres"
	"hallo-corona/repositories"

	"github.com/gin-gonic/gin"
)

func ConsultationRoutes(r *gin.RouterGroup) {
	consultationRepository := repositories.RepositoryConsultation(postgres.DB)
	h := handlers.HandlerConsultation(consultationRepository)

	r.GET("/consultations", middleware.Auth(h.FindConsultations))
	r.GET("/consultation/:id", middleware.Auth(h.GetConsultation))
	r.GET("/consultations/:id", middleware.Auth(h.FindMyConsultation))
	r.POST("/consultation", middleware.Auth(h.CreateConsultation))
	r.PATCH("/consultation/:id", middleware.Auth(h.UpdateConsultation))
	r.DELETE("/consultation/:id", middleware.Auth(h.DeleteConsultation))
}
