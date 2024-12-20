package categoryhandler

import (
	"net/http"

	categorymodel "github.com/PhuPhuoc/koi-story-api-v2/controller/category_services/model"
	categoryrepository "github.com/PhuPhuoc/koi-story-api-v2/controller/category_services/repository"
	"github.com/PhuPhuoc/koi-story-api-v2/utils"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

//	@BasePath		/api/v1
//	@Summary		register new category
//	@Description	register new category
//	@Tags			categories
//	@Accept			json
//	@Produce		json
//	@Param			user	body		categorymodel.CategoryForCreateAndUpdate	true	"User register data"
//	@Success		200		{object}	map[string]interface{}						"message success"
//	@Failure		400		{object}	error										"Bad request error"
//	@Router			/categories [post]
func createNewCategoryHandler(db *sqlx.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req categorymodel.CategoryForCreateAndUpdate
		if err := c.ShouldBindJSON(&req); err != nil {
			utils.SendError(c, http.StatusBadRequest, "invalid request body", "")
			return
		}
		repo := categoryrepository.NewCategoryStore(db)
		if err := repo.CreateCategory(&req); err != nil {
			utils.SendError(c, http.StatusBadRequest, "cannot create new category", err.Error())
			return
		}
		utils.SendSuccess(c, http.StatusOK, "new category creation successful", nil, nil)
	}
}
