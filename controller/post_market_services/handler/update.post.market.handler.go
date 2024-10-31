package markethandler

import (
	"net/http"

	marketmodel "github.com/PhuPhuoc/koi-story-api-v2/controller/post_market_services/model"
	marketrepository "github.com/PhuPhuoc/koi-story-api-v2/controller/post_market_services/repository"
	"github.com/PhuPhuoc/koi-story-api-v2/utils"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

// @BasePath		/api/v1
// @Summary		update market
// @Description	update market
// @Tags			post market
// @Accept			json
// @Produce		json
// @Param			post_id	path		string					true							"Post ID"
// @Param			market	info		body					marketmodel.MarketInfoUpdate	true	"User register data"
// @Success		200		{object}	map[string]interface{}	"message success"
// @Failure		400		{object}	error					"Bad request error"
// @Router			/post-market/{post_id} [put]
func updatePostMarketHandler(db *sqlx.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		post_id := c.Param("post_id")
		if post_id == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Post ID is required"})
			return
		}
		var req marketmodel.MarketInfoUpdate
		if err := c.ShouldBindJSON(&req); err != nil {
			utils.SendError(c, http.StatusBadRequest, "invalid request body", "")
			return
		}
		repo := marketrepository.NewPostMarketStore(db)
		if err := repo.UpdateMarket(post_id, &req); err != nil {
			utils.SendError(c, http.StatusBadRequest, "cannot create new market's post", err.Error())
			return
		}
		utils.SendSuccess(c, http.StatusOK, "new account registration successful", nil, nil)
	}
}
