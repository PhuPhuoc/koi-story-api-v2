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
//	@Summary		create new comment
//	@Description	create new comment
//	@Tags			comments
//	@Accept			json
//	@Produce		json
//	@Param			post_id	path		string						true	"Post ID"
//	@Param			comment	body		commentmodel.CreateComment	true	"comment information"
//	@Success		200		{object}	map[string]interface{}		"message success"
//	@Failure		400		{object}	error						"Bad request error"
//	@Router			/posts/{post_id}/comments [post]
func createCommentHandler(db *sqlx.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		post_id := c.Param("post_id")
		if post_id == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "User ID is required"})
			return
		}
		var req commentmodel.CreateComment
		if err := c.ShouldBindJSON(&req); err != nil {
			utils.SendError(c, http.StatusBadRequest, "invalid request body", "")
			return
		}
		repo := commentrepository.NewCommentStore(db)
		if err := repo.CreateNewComment(post_id, &req); err != nil {
			utils.SendError(c, http.StatusBadRequest, "cannot add new comment into post", err.Error())
			return
		}
		utils.SendSuccess(c, http.StatusOK, "new comment creation successful", nil, nil)
	}
}
