package imagehandler

import (
	"net/http"

	imagemodel "github.com/PhuPhuoc/koi-story-api-v2/controller/image_services/model"
	imagerepository "github.com/PhuPhuoc/koi-story-api-v2/controller/image_services/repository"
	"github.com/PhuPhuoc/koi-story-api-v2/utils"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

// @BasePath		/api/v1
// @Summary		add new image into post
// @Description	add new image into post
// @Tags			images
// @Accept			json
// @Produce		json
// @Param			post_id	path		string						true	"Post ID"
// @Param			images	body		imagemodel.AddImageModel	true	"image url"
// @Success		200		{object}	map[string]interface{}		"message success"
// @Failure		400		{object}	error						"Bad request error"
// @Router			/posts/{post_id}/images [post]
func addNewImageHandler(db *sqlx.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		post_id := c.Param("post_id")
		if post_id == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "post ID is required"})
			return
		}
		var req imagemodel.AddImageModel
		if err := c.ShouldBindJSON(&req); err != nil {
			utils.SendError(c, http.StatusBadRequest, "invalid request body", "")
			return
		}
		repo := imagerepository.NewImageStore(db)
		if err := repo.AddNewImage(post_id, &req); err != nil {
			utils.SendError(c, http.StatusBadRequest, "cannot add new image into post", err.Error())
			return
		}
		utils.SendSuccess(c, http.StatusOK, "image added successfully", nil, nil)
	}
}
