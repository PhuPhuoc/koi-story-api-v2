package commenthandler

import (
	"net/http"

	commentrepository "github.com/PhuPhuoc/koi-story-api-v2/controller/comment_services/repository"
	"github.com/PhuPhuoc/koi-story-api-v2/utils"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

// @BasePath		/api/v1
// @Summary		Get all comments in a post
// @Description	Get all comments in a post
// @Tags			comments
// @Accept			json
// @Produce		json
// @Param			post_id	path		string					true	"Post ID"
// @Success		200		{object}	map[string]interface{}	"message success"
// @Failure		400		{object}	error					"Bad request error"
// @Router			/posts/{post_id}/comments [get]
func getAllCommentHandler(db *sqlx.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		post_id := c.Param("post_id")
		if post_id == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "User ID is required"})
			return
		}
		repo := commentrepository.NewCommentStore(db)
		data, err := repo.GetAllComment(post_id)
		if err != nil {
			utils.SendError(c, http.StatusBadRequest, "cannot add new comment into post", err.Error())
			return
		}
		utils.SendSuccess(c, http.StatusOK, "get data successfully", data, nil)
	}
}
