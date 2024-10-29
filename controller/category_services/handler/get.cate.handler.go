package categoryhandler

import (
	"net/http"

	categoryrepository "github.com/PhuPhuoc/koi-story-api-v2/controller/category_services/repository"
	"github.com/PhuPhuoc/koi-story-api-v2/utils"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

// @BasePath		/api/v1
// @Summary		get list categories
// @Description	get list categories
// @Tags			categories
// @Accept			json
// @Produce		json
// @Success		200	{object}	map[string]interface{}	"data"
// @Failure		400	{object}	error					"Bad request error"
// @Router			/categories [get]
func getCategoriesHandler(db *sqlx.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		repo := categoryrepository.NewCategoryStore(db)
		list_fate, err := repo.GetCategories()
		if err != nil {
			utils.SendError(c, http.StatusBadRequest, err.Error())
			return
		}
		utils.SendSuccess(c, http.StatusOK, "Get data successfully", list_fate, nil)
	}
}
