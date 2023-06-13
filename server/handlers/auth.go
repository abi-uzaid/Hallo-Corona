package handlers

import (
	// "fmt"
	authdto "hallo-corona/dto/auth"
	dto "hallo-corona/dto/result"
	"hallo-corona/models"
	"hallo-corona/pkg/bcrypt"
	jwtToken "hallo-corona/pkg/jwt"
	"hallo-corona/repositories"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
)

type handlerAuth struct {
	AuthRepository repositories.AuthRepository
}

func HandlerAuth(AuthRepository repositories.AuthRepository) *handlerAuth {
	return &handlerAuth{AuthRepository}
}

func (h *handlerAuth) Register(c *gin.Context) {
	request := new(authdto.RegisterRequest)
	if err := c.Bind(request); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
		return
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
		return
	}

	password, err := bcrypt.HashingPassword(request.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
		return
	}

	user := models.User{
		Fullname: request.Fullname,
		Username: request.Username,
		Email:    request.Email,
		Password: password,
		// Role:     "",
		ListAs:  request.ListAs,
		Gender:  request.Gender,
		Phone:   request.Phone,
		Address: request.Address,
	}

	data, err := h.AuthRepository.Register(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Message: "Your registration is success", Data: data})
	return
}

func (h *handlerAuth) Login(c *gin.Context) {
	request := new(authdto.LoginRequest)
	if err := c.Bind(request); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
		return
	}

	data := models.User{
		Username: request.Username,
		Password: request.Password,
	}

	userLogin, err := h.AuthRepository.Login(data.Username)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: "username not registered"})
		return
	}

	isValid := bcrypt.CheckPasswordHash(request.Password, userLogin.Password)
	if !isValid {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: "Incorrect email or password"})
		return
	}

	claims := jwt.MapClaims{}
	claims["id"] = userLogin.ID
	claims["listAs"] = userLogin.ListAs
	// claims["role"] = userLogin.Role
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // 4 hours expired

	token, generateTokenErr := jwtToken.GenerateToken(&claims)
	if generateTokenErr != nil {
		log.Println(generateTokenErr)
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error" : "Unauthorized"})
		return
	}

	loginResponse := authdto.LoginResponse{
		ID:       userLogin.ID,
		Fullname: userLogin.Fullname,
		Username: userLogin.Username,
		Email:    userLogin.Email,
		ListAs:   userLogin.ListAs,
		// Role:     userLogin.Role,
		Gender:  userLogin.Gender,
		Phone:   userLogin.Phone,
		Address: userLogin.Address,
		Token:   token,
	}

	c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: loginResponse})
	return
}

func (h *handlerAuth) CheckAuth(c *gin.Context) {
	userLogin := c.MustGet("userLogin")
	userId := userLogin.(jwt.MapClaims)["id"].(float64)

	user, _ := h.AuthRepository.CheckAuth(int(userId))

	c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: responseCheckAuth(user)})
	return
}

// type resp struct {
// 	ID       uint
// 	FullName string
// 	Username string
// 	Email    string
// 	ListAs   string
// 	// Role    string
// 	Gender  string
// 	Phone   string
// 	Address string
// 	Image   string
// }

func responseCheckAuth(u models.User) authdto.CheckAuthResponse {
	return authdto.CheckAuthResponse{
		ID:       u.ID,
		Fullname: u.Fullname,
		Username: u.Username,
		Email:    u.Email,
		ListAs:   u.ListAs,
		// Role:    u.Role,
		Gender:  u.Gender,
		Phone:   u.Phone,
		Address: u.Address,
		// Image:   u.Image,
	}
}
