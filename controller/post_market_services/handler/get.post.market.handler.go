package markethandler

import (
	"net/http"

	marketrepository "github.com/PhuPhuoc/koi-story-api-v2/controller/post_market_services/repository"
	"github.com/PhuPhuoc/koi-story-api-v2/utils"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

// @BasePath		/api/v1
// @Summary		get market's posts
// @Description	get market's posts
// @Tags			post market
// @Accept			json
// @Produce		json
// @Success		200	{object}	map[string]interface{}	"data"
// @Failure		400	{object}	error					"Bad request error"
// @Router			/post-market [get]
func getPostMarketHandler(db *sqlx.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		repo := marketrepository.NewPostMarketStore(db)
		list_post, err := repo.GetListPostMarket()
		if err != nil {
			utils.SendError(c, http.StatusBadRequest, "cannot get list post market", err.Error())
			return
		}
		utils.SendSuccess(c, http.StatusOK, "Get data successfully", list_post, nil)
	}
}
