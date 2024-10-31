package consulthandler

import (
	"net/http"

	consultrepository "github.com/PhuPhuoc/koi-story-api-v2/controller/post_consult_services/repository"
	"github.com/PhuPhuoc/koi-story-api-v2/utils"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

// @BasePath		/api/v1
// @Summary		get consult's posts
// @Description	get consult's posts
// @Tags			post consult
// @Accept			json
// @Produce		json
// @Success		200	{object}	map[string]interface{}	"data"
// @Failure		400	{object}	error					"Bad request error"
// @Router			/post-consult [get]
func getPostConsultHandler(db *sqlx.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		repo := consultrepository.NewPostConsultStore(db)
		list_post, err := repo.GetListPostConsult()
		if err != nil {
			utils.SendError(c, http.StatusBadRequest, "cannot get list post consult", err.Error())
			return
		}
		utils.SendSuccess(c, http.StatusOK, "Get data successfully", list_post, nil)
	}
}
