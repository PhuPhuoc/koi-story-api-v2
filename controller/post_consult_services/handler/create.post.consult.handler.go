package consulthandler

import (
	"net/http"

	consultmodel "github.com/PhuPhuoc/koi-story-api-v2/controller/post_consult_services/model"
	consultrepository "github.com/PhuPhuoc/koi-story-api-v2/controller/post_consult_services/repository"
	"github.com/PhuPhuoc/koi-story-api-v2/utils"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

//	@BasePath		/api/v1
//	@Summary		create new post consult
//	@Description	create new post consult
//	@Tags			post consult
//	@Accept			json
//	@Produce		json
//	@Param			consult	body		consultmodel.CreateConsultModel	true	"User register data"
//	@Success		200		{object}	map[string]interface{}			"message success"
//	@Failure		400		{object}	error							"Bad request error"
//	@Router			/post-consult [post]
func createNewPostConsultHandler(db *sqlx.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req consultmodel.CreateConsultModel
		if err := c.ShouldBindJSON(&req); err != nil {
			utils.SendError(c, http.StatusBadRequest, "invalid request body", "")
			return
		}
		repo := consultrepository.NewPostConsultStore(db)
		if err := repo.CreateNewConsultPost(&req); err != nil {
			utils.SendError(c, http.StatusBadRequest, "cannot create new market's post", err.Error())
			return
		}
		utils.SendSuccess(c, http.StatusOK, "new account registration successful", nil, nil)
	}
}
