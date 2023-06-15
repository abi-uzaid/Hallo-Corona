package handlers

import (
	"context"
	"fmt"
	articledto "hallo-corona/dto/article"
	dto "hallo-corona/dto/result"
	"hallo-corona/models"
	"hallo-corona/repositories"
	"net/http"
	"os"
	"strconv"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
)

type handlerArticle struct {
	ArticleRepository repositories.ArticleRepository
}

func HandlerArticle(ArticleRepository repositories.ArticleRepository) *handlerArticle {
	return &handlerArticle{ArticleRepository}
}

func (h *handlerArticle) FindArticles(c *gin.Context) {
	articles, err := h.ArticleRepository.FindArticles()
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: articles})
}

func (h *handlerArticle) GetArticel(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	products, err := h.ArticleRepository.GetArticle(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: convertResponseArticle(products)})
}

func (h *handlerArticle) CreateArticle(c *gin.Context) {
	// c.Header("Content-Type", "multipart/form-data")

	userLogin := c.MustGet("userLogin")
	userId := userLogin.(jwt.MapClaims)["id"].(float64)

	// if userId {
	dataFile := c.MustGet("dataFile").(string)
	fmt.Println("this is data file", dataFile)

	// UserID, _ := strconv.Atoi(c.Request.FormValue("userId"))

	request := articledto.CreateArticleRequest{
		Title:    c.Request.FormValue("title"),
		UserID:   int(userId),
		Image:    dataFile,
		Desc:     c.Request.FormValue("desc"),
		Category: c.Request.FormValue("category"),
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
		return
	}

	// var ctx = context.Background()
	// var CLOUD_NAME = os.Getenv("CLOUD_NAME")
	// var API_KEY = os.Getenv("API_KEY")
	// var API_SECRET = os.Getenv("API_SECRET")

	// // Add your Cloudinary credentials ...
	// cld, _ := cloudinary.NewFromParams(CLOUD_NAME, API_KEY, API_SECRET)

	// // Upload file to Cloudinary ...
	// resp, err := cld.Upload.Upload(ctx, dataFile, uploader.UploadParams{Folder: "uploads"})

	// if err != nil {
	// 	fmt.Println(err.Error())
	// }

	// submit to db article
	article := models.Article{
		Title:  request.Title,
		UserID: request.UserID,
		// User:     models.UserResponse{},
		Image:    request.Image,
		Desc:     request.Desc,
		Category: request.Category,
	}

	data, err := h.ArticleRepository.CreateArticle(article)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Message: "Create article success", Data: data})

	// } else {
	// 	c.JSON(http.StatusUnauthorized, dto.ErrorResult{Code: http.StatusUnauthorized, Message: "error"})
	// 	return
	// }

}

func (h *handlerArticle) UpdateArticle(c *gin.Context) {
	userLogin := c.MustGet("userLogin")
	userRole := userLogin.(jwt.MapClaims)["listAs"].(string)
	if userRole == "doctor" {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
			return
		}

		dataFile := c.MustGet("dataFile").(string)

		request := articledto.UpdateArticleRequest{
			Title: c.PostForm("title"),
			Image: dataFile,
			Desc:  c.PostForm("desc"),
		}

		validation := validator.New()
		err = validation.Struct(request)
		if err != nil {
			c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
			return
		}

		var ctx = context.Background()
		var CLOUD_NAME = os.Getenv("CLOUD_NAME")
		var API_KEY = os.Getenv("API_KEY")
		var API_SECRET = os.Getenv("API_SECRET")

		// Add your Cloudinary credentials ...
		cld, _ := cloudinary.NewFromParams(CLOUD_NAME, API_KEY, API_SECRET)

		// Upload file to Cloudinary ...
		resp, err := cld.Upload.Upload(ctx, dataFile, uploader.UploadParams{Folder: "hallo corona"})

		if err != nil {
			fmt.Println(err.Error())
		}

		article, err := h.ArticleRepository.GetArticle(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
			return
		}

		if request.Title != "" {
			article.Title = request.Title
		}

		if request.Image != "" {
			article.Image = resp.SecureURL
		}

		if request.Desc != "" {
			article.Desc = request.Desc
		}

		articleUpdate, err := h.ArticleRepository.UpdateArticle(article)
		if err != nil {
			c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
			return
		}
		c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: articleUpdate})
		return
	}
	c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: "Access denied"})
}

func (h *handlerArticle) DeleteArticle(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: "Invalid article ID"})
		return
	}

	article, err := h.ArticleRepository.GetArticle(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
		return
	}

	articleDel, err := h.ArticleRepository.DeleteArticle(article)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Message: "Delete article success", Data: articleDel})
}

func (h *handlerArticle) FindMyArticle(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
		return
	}

	articles, err := h.ArticleRepository.FindMyArticle(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Message: "Succes", Data: articles})
}

func convertResponseArticle(u models.Article) articledto.ArticleResponse {
	return articledto.ArticleResponse{
		ID:        u.ID,
		Title:     u.Title,
		Image:     u.Image,
		User:      u.User,
		Desc:      u.Desc,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
		Category:  u.Category,
	}
}
