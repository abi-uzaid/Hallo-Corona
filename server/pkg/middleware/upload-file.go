package middleware

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/gin-gonic/gin"
)

func UploadFile(next gin.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		file, err := c.FormFile("image")
		// if file != nil {
			if err != nil {
				c.JSON(http.StatusBadRequest, err.Error())
				return
			}

			src, err := file.Open()
			if err != nil {
				c.JSON(http.StatusBadRequest, err.Error())
				return
			}
			defer src.Close()

			var ctx = context.Background()
			var CLOUD_NAME = os.Getenv("CLOUD_NAME")
			var API_KEY = os.Getenv("API_KEY")
			var API_SECRET = os.Getenv("API_SECRET")

			// Add your Cloudinary credentials ...
			cld, _ := cloudinary.NewFromParams(CLOUD_NAME, API_KEY, API_SECRET)

			// Upload file to Cloudinary ...
			resp, err := cld.Upload.Upload(ctx, src, uploader.UploadParams{Folder: "uploads"})

			if err != nil {
				fmt.Println(err.Error())
			}

			c.Set("dataFile", resp.SecureURL)
			next(c)
		}
		// c.Set("dataFile", "")
		// next(c)

	// }
}
