package handlers

import (
	consultationdto "hallo-corona/dto/consultation"
	dto "hallo-corona/dto/result"
	"hallo-corona/models"
	"hallo-corona/repositories"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
)

type handlerConsultation struct {
	ConsultationRepository repositories.ConsultationRepository
}

func HandlerConsultation(ConsultationRepository repositories.ConsultationRepository) *handlerConsultation {
	return &handlerConsultation{ConsultationRepository}
}

func (h *handlerConsultation) FindConsultations(c *gin.Context) {
	consultations, err := h.ConsultationRepository.FindConsultations()
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Message: "Find consultation success", Data: consultations})
	return
}

func (h *handlerConsultation) GetConsultation(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	consult, err := h.ConsultationRepository.GetConsultation(int(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: consult})
	return
}

func (h *handlerConsultation) CreateConsultation(c *gin.Context) {
	userLogin := c.MustGet("userLogin")
	userId := userLogin.(jwt.MapClaims)["id"].(float64)

	var request consultationdto.CreateConsultationRequest
	err := c.Bind(&request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
		return
	}

	validation := validator.New()
	err = validation.Struct(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
		return
	}

	consultation := models.Consultation{
		Fullname:   request.Fullname,
		Phone:      request.Phone,
		BornDate:   request.BornDate,
		Age:        request.Age,
		Height:     request.Height,
		Weight:     request.Weight,
		Gender:     request.Gender,
		Subject:    request.Subject,
		LiveConsul: request.LiveConsul,
		Desc:       request.Desc,
		UserID:     int(userId),
		User:       models.UserResponse{},
		Status:     "pending",
	}

	data, err := h.ConsultationRepository.CreateConsultation(consultation)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Message: "Create consultation success", Data: data})
	return
}

func (h *handlerConsultation) DeleteConsultation(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	consultation, err := h.ConsultationRepository.GetConsultation(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
		return
	}

	data, err := h.ConsultationRepository.DeleteConsultation(consultation)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Message: "Delete consultation success", Data: convertResponseConsultation(data)})
	return

}

func (h *handlerConsultation) UpdateConsultation(c *gin.Context) {
	request := new(consultationdto.UpdateConsultationRequest)
	if err := c.Bind(request); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
		return
	}

	id, _ := strconv.Atoi(c.Param("id"))
	consultation, err := h.ConsultationRepository.GetConsultation(int(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
		return
	}

	if request.Fullname != "" {
		consultation.Fullname = request.Fullname
	}

	if request.Phone != "" {
		consultation.Phone = request.Phone
	}

	if request.BornDate != "" {
		consultation.BornDate = request.BornDate
	}

	if request.Age != 0 {
		consultation.Age = request.Age
	}

	if request.Height != 0 {
		consultation.Height = request.Height
	}

	if request.Weight != 0 {
		consultation.Weight = request.Weight
	}

	if request.Gender != "" {
		consultation.Gender = request.Gender
	}

	if request.Subject != "" {
		consultation.Subject = request.Subject
	}

	if request.LiveConsul != "" {
		consultation.LiveConsul = request.LiveConsul
	}

	if request.Desc != "" {
		consultation.Desc = request.Desc
	}

	if request.Reply != "" {
		consultation.Reply = request.Reply
	}

	if request.Reply != "" {
		consultation.Status = "waiting"
	}

	if request.LinkLive != "" {
		consultation.LinkLive = request.LinkLive
	}

	data, err := h.ConsultationRepository.UpdateConsultation(consultation)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: data})
	return
}

func (h *handlerConsultation) FindMyConsultation(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	consult, err := h.ConsultationRepository.FindMyConsultation(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: consult})
	return
}

func convertResponseConsultation(u models.Consultation) consultationdto.ConsultationResponse {
	return consultationdto.ConsultationResponse{
		ID:         u.ID,
		Fullname:   u.Fullname,
		Phone:      u.Phone,
		Age:        u.Age,
		Height:     u.Height,
		Weight:     u.Weight,
		Gender:     u.Gender,
		LiveConsul: u.LiveConsul,
		Desc:       u.Desc,
		User:       u.User,
		Subject:    u.Subject,
		Status:     u.Status,
		Reply:      u.Reply,
		LinkLive:   u.LinkLive,
		// Reservation: u.Reservation,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}
