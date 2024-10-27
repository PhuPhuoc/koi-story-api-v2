package fatehandler

import (
	"net/http"

	faterepository "github.com/PhuPhuoc/koi-story-api-v2/controller/fate_services/repository"
	"github.com/PhuPhuoc/koi-story-api-v2/utils"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

//	@BasePath		/api/v1
//	@Summary		get list fates
//	@Description	get list fates
//	@Tags			fates
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	map[string]interface{}	"data"
//	@Failure		400	{object}	error					"Bad request error"
//	@Router			/fates [get]
func getFatesHandler(db *sqlx.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		repo := faterepository.NewFateStore(db)
		list_fate, err := repo.GetFates()
		if err != nil {
			utils.SendError(c, http.StatusBadRequest, err.Error())
			return
		}
		utils.SendSuccess(c, http.StatusOK, "Get data successfully", list_fate, nil)
	}
}
