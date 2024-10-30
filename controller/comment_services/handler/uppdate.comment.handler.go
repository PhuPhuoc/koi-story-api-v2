package commenthandler

import (
	"net/http"

	commentmodel "github.com/PhuPhuoc/koi-story-api-v2/controller/comment_services/model"
	commentrepository "github.com/PhuPhuoc/koi-story-api-v2/controller/comment_services/repository"
	"github.com/PhuPhuoc/koi-story-api-v2/utils"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

//	@BasePath		/api/v1
//	@Summary		update comment
//	@Description	update comment
//	@Tags			comments
//	@Accept			json
//	@Produce		json
//	@Param			comment_id	path		string						true	"Comment ID"
//	@Param			comment		body		commentmodel.UpdateComment	true	"comment's content update"
//	@Success		200			{object}	map[string]interface{}		"message success"
//	@Failure		400			{object}	error						"Bad request error"
//	@Router			/comments/{comment_id} [put]
func updateCommentHandler(db *sqlx.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		comment_id := c.Param("comment_id")
		if comment_id == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Comment ID is required"})
			return
		}
		var req commentmodel.UpdateComment
		if err := c.ShouldBindJSON(&req); err != nil {
			utils.SendError(c, http.StatusBadRequest, "invalid request body", "")
			return
		}
		repo := commentrepository.NewCommentStore(db)
		if err := repo.UpdateComment(comment_id, &req); err != nil {
			utils.SendError(c, http.StatusBadRequest, "cannot update comment", err.Error())
			return
		}
		utils.SendSuccess(c, http.StatusOK, "comment updated successful", nil, nil)
	}
}
