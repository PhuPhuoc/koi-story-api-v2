package markethandler

import (
	"net/http"

	marketrepository "github.com/PhuPhuoc/koi-story-api-v2/controller/post_market_services/repository"
	"github.com/PhuPhuoc/koi-story-api-v2/utils"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

// @BasePath		/api/v1
// @Summary		get my market's posts
// @Description	get my market's posts
// @Tags			post market
// @Accept			json
// @Produce		json
// @Param			user_id	path		string					true	"User ID"
// @Success		200		{object}	map[string]interface{}	"data"
// @Failure		400		{object}	error					"Bad request error"
// @Router			/post-market/user/{user_id} [get]
func getMyPostMarketHandler(db *sqlx.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		user_id := c.Param("user_id")
		if user_id == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "User ID is required"})
			return
		}
		repo := marketrepository.NewPostMarketStore(db)
		list_post, err := repo.GetMyPostMarketList(user_id)
		if err != nil {
			utils.SendError(c, http.StatusBadRequest, "cannot get list post market", err.Error())
			return
		}
		utils.SendSuccess(c, http.StatusOK, "Get data successfully", list_post, nil)
	}
}
