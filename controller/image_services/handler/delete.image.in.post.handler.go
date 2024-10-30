package imagehandler

import (
	"net/http"

	imagerepository "github.com/PhuPhuoc/koi-story-api-v2/controller/image_services/repository"
	"github.com/PhuPhuoc/koi-story-api-v2/utils"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

// @BasePath		/api/v1
// @Summary		delete image
// @Description	delete image
// @Tags			images
// @Accept			json
// @Produce		json
// @Param			image_id	path		string					true	"Image ID"
// @Success		200			{object}	map[string]interface{}	"message success"
// @Failure		400			{object}	error					"Bad request error"
// @Router			/images/{image_id} [delete]
func deleteImageHandler(db *sqlx.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		image_id := c.Param("image_id")
		if image_id == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Comment ID is required"})
			return
		}
		repo := imagerepository.NewImageStore(db)
		if err := repo.DeleteImage(image_id); err != nil {
			utils.SendError(c, http.StatusBadRequest, "cannot delete image", err.Error())
			return
		}
		utils.SendSuccess(c, http.StatusOK, "image deleted successful", nil, nil)
	}
}
