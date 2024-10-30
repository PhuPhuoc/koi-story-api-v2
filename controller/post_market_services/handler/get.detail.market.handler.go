package markethandler

import (
	"net/http"

	marketrepository "github.com/PhuPhuoc/koi-story-api-v2/controller/post_market_services/repository"
	"github.com/PhuPhuoc/koi-story-api-v2/utils"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

//	@BasePath		/api/v1
//	@Summary		get detail market's posts
//	@Description	get detail market's posts
//	@Tags			post market
//	@Accept			json
//	@Produce		json
//	@Param			post_id	path		string					true	"Post ID"
//	@Success		200		{object}	map[string]interface{}	"data"
//	@Failure		400		{object}	error					"Bad request error"
//	@Router			/post-market/{post_id} [get]
func getDetailPostMarketHandler(db *sqlx.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		post_id := c.Param("post_id")
		if post_id == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Post ID is required"})
			return
		}
		repo := marketrepository.NewPostMarketStore(db)
		post_details, err := repo.GetPostMarketDetails(post_id)
		if err != nil {
			utils.SendError(c, http.StatusBadRequest, "cannot get detail of post market", err.Error())
			return
		}
		utils.SendSuccess(c, http.StatusOK, "Get data successfully", post_details, nil)
	}
}
