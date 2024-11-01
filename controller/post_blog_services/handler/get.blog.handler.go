package bloghandler

import (
	"net/http"

	blogrepository "github.com/PhuPhuoc/koi-story-api-v2/controller/post_blog_services/repository"
	"github.com/PhuPhuoc/koi-story-api-v2/utils"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

//	@BasePath		/api/v1
//	@Summary		get blog's posts
//	@Description	get blog's posts
//	@Tags			post blog
//	@Accept			json
//	@Produce		json
//	@Param			category_id	path		string					true	"User ID"
//	@Success		200			{object}	map[string]interface{}	"data"
//	@Failure		400			{object}	error					"Bad request error"
//	@Router			/post-blog/category/{category_id} [get]
func getBlogByCategoryIDHandler(db *sqlx.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		category_id := c.Param("category_id")
		if category_id == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "User ID is required"})
			return
		}
		repo := blogrepository.NewPostBlogStore(db)
		list_post, err := repo.GetBlogByCategory(category_id)
		if err != nil {
			utils.SendError(c, http.StatusBadRequest, "cannot get list post consult", err.Error())
			return
		}
		utils.SendSuccess(c, http.StatusOK, "Get data successfully", list_post, nil)
	}
}
