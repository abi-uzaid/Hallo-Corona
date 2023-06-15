package handlers

import (
	dto "hallo-corona/dto/result"
	usersdto "hallo-corona/dto/users"
	"hallo-corona/models"
	"hallo-corona/repositories"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
)

type handler struct {
	UserRepository repositories.UserRepository
}

func HandlerUser(UserRepository repositories.UserRepository) *handler {
	return &handler{UserRepository: UserRepository}
}

func (h *handler) FindUsers(c *gin.Context) {
	users, err := h.UserRepository.FindUsers()
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: users})

}

func (h *handler) GetUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	user, err := h.UserRepository.GetUser(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Message: "User data successfully obtained", Data: user})

}

func (h *handler) UpdateUser(c *gin.Context) {
	userLogin := c.MustGet("userLogin")
	userId := userLogin.(jwt.MapClaims)["id"].(float64)

	request := new(usersdto.UpdateUserRequest)
	if err := c.Bind(&request); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
		return
	}

	user, err := h.UserRepository.GetUser(int(userId))

	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
		return
	}

	if request.Fullname != "" {
		user.Fullname = request.Fullname
	}

	if request.Username != "" {
		user.Username = request.Username
	}

	if request.Email != "" {
		user.Email = request.Email
	}

	if request.Password != "" {
		user.Password = request.Password
	}

	if request.Gender != "" {
		user.Gender = request.Gender
	}

	if request.Phone != "" {
		user.Phone = request.Phone
	}
	if request.Address != "" {
		user.Address = request.Address
	}
	if request.Image != "" {
		user.Address = request.Image
	}

	data, err := h.UserRepository.UpdateUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Message: "User data updated successfully", Data: convertResponse(data)})

}

func (h *handler) DeleteUser(c *gin.Context) {
	userLogin := c.MustGet("userLogin")
	userId := userLogin.(jwt.MapClaims)["id"].(float64)

	// id, _ := strconv.Atoi(c.Param("id"))

	user, err := h.UserRepository.GetUser(int(userId))
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
		return
	}

	data, err := h.UserRepository.DeleteUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Message: "Delete Success", Data: convertResponse(data)})

}

func convertResponse(u models.User) usersdto.UserResponse {
	return usersdto.UserResponse{
		ID:       u.ID,
		Fullname: u.Fullname,
		Username: u.Username,
		Email:    u.Email,
		Password: u.Password,
		Gender:   u.Gender,
		Phone:    u.Phone,
		Address:  u.Address,
	}
}

func (h *handler) ChangeImage(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
		return
	}

	dataFile := c.MustGet("dataFile").(string)

	request := usersdto.ChangeImageRequest{
		Image: dataFile,
	}

	validation := validator.New()
	err = validation.Struct(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
		return
	}

	user, err := h.UserRepository.GetUser(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})

	}

	data, err := h.UserRepository.ChangeImage(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: data})

}
