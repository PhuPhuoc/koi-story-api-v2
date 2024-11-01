package bloghandler

import (
	"net/http"

	blogmodel "github.com/PhuPhuoc/koi-story-api-v2/controller/post_blog_services/model"
	blogrepository "github.com/PhuPhuoc/koi-story-api-v2/controller/post_blog_services/repository"
	"github.com/PhuPhuoc/koi-story-api-v2/utils"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

//	@BasePath		/api/v1
//	@Summary		create new post blog
//	@Description	create new post blog
//	@Tags			post blog
//	@Accept			json
//	@Produce		json
//	@Param			consult	body		blogmodel.CreationBlogModel	true	"User register data"
//	@Success		200		{object}	map[string]interface{}		"message success"
//	@Failure		400		{object}	error						"Bad request error"
//	@Router			/post-blog [post]
func createNewPostBlogHandler(db *sqlx.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req blogmodel.CreationBlogModel
		if err := c.ShouldBindJSON(&req); err != nil {
			utils.SendError(c, http.StatusBadRequest, "invalid request body", "")
			return
		}
		repo := blogrepository.NewPostBlogStore(db)
		if err := repo.CreateNewBlogPost(&req); err != nil {
			utils.SendError(c, http.StatusBadRequest, "cannot create new market's post", err.Error())
			return
		}
		utils.SendSuccess(c, http.StatusOK, "new account registration successful", nil, nil)
	}
}
