package commenthandler

import (
	"net/http"

	commentrepository "github.com/PhuPhuoc/koi-story-api-v2/controller/comment_services/repository"
	"github.com/PhuPhuoc/koi-story-api-v2/utils"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

//	@BasePath		/api/v1
//	@Summary		delete comment
//	@Description	delete comment
//	@Tags			comments
//	@Accept			json
//	@Produce		json
//	@Param			comment_id	path		string					true	"Comment ID"
//	@Success		200			{object}	map[string]interface{}	"message success"
//	@Failure		400			{object}	error					"Bad request error"
//	@Router			/comments/{comment_id} [delete]
func deleteCommentHandler(db *sqlx.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		comment_id := c.Param("comment_id")
		if comment_id == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Comment ID is required"})
			return
		}
		repo := commentrepository.NewCommentStore(db)
		if err := repo.DeleteComment(comment_id); err != nil {
			utils.SendError(c, http.StatusBadRequest, "cannot delete comment", err.Error())
			return
		}
		utils.SendSuccess(c, http.StatusOK, "comment deleted successful", nil, nil)
	}
}
