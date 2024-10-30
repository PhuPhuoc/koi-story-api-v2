package fatehandler

import (
	"net/http"

	fatemodel "github.com/PhuPhuoc/koi-story-api-v2/controller/fate_services/model"
	faterepository "github.com/PhuPhuoc/koi-story-api-v2/controller/fate_services/repository"
	"github.com/PhuPhuoc/koi-story-api-v2/utils"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

// @BasePath		/api/v1
// @Summary		generate user fate
// @Description	generate user fate
// @Tags			fates
// @Accept			json
// @Produce		json
// @Param			user	body		fatemodel.GenerateUserFengshui	true	"User info for generate feng shui"
// @Success		200		{object}	map[string]interface{}			"message success"
// @Failure		400		{object}	error							"Bad request error"
// @Router			/fates/user [post]
func generateUserFateHandler(db *sqlx.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req fatemodel.GenerateUserFengshui
		if err := c.ShouldBindJSON(&req); err != nil {
			utils.SendError(c, http.StatusBadRequest, "invalid request body", "")
			return
		}
		repo := faterepository.NewFateStore(db)
		if err := repo.CalculateUserFate(&req); err != nil {
			utils.SendError(c, http.StatusBadRequest, "", err.Error())
			return
		}
		utils.SendSuccess(c, http.StatusOK, "user's fate generation successful", nil, nil)
	}
}
