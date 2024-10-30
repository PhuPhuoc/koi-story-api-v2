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
// @Summary		create new post market
// @Description	create new post market
// @Tags			post market
// @Accept			json
// @Produce		json
// @Param			user	body		marketmodel.FormCreatePostMarket	true	"User register data"
// @Success		200		{object}	map[string]interface{}				"message success"
// @Failure		400		{object}	error								"Bad request error"
// @Router			/post-market [post]
func createNewPostMarketHandler(db *sqlx.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req marketmodel.FormCreatePostMarket
		if err := c.ShouldBindJSON(&req); err != nil {
			utils.SendError(c, http.StatusBadRequest, "invalid request body", "")
			return
		}
		repo := marketrepository.NewPostMarketStore(db)
		if err := repo.CreateNewPost(&req); err != nil {
			utils.SendError(c, http.StatusBadRequest, "cannot create new market's post", err.Error())
			return
		}
		utils.SendSuccess(c, http.StatusOK, "new account registration successful", nil, nil)
	}
}