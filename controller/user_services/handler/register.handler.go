package userhandler

import (
	"net/http"

	usermodel "github.com/PhuPhuoc/koi-story-api-v2/controller/user_services/model"
	userrepository "github.com/PhuPhuoc/koi-story-api-v2/controller/user_services/repository"
	"github.com/PhuPhuoc/koi-story-api-v2/utils"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

// @BasePath		/api/v1
// @Summary		register new account
// @Description	register new account
// @Tags			users
// @Accept			json
// @Produce		json
// @Param			user	body		usermodel.Register		true	"User register data"
// @Success		200		{object}	map[string]interface{}	"message success"
// @Failure		400		{object}	error					"Bad request error"
// @Router			/users/register [post]
func registerHandler(db *sqlx.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req usermodel.Register
		if err := c.ShouldBindJSON(&req); err != nil {
			utils.SendError(c, http.StatusBadRequest, "invalid request body", "")
			return
		}
		repo := userrepository.NewUserStore(db)
		if err := repo.Register(&req); err != nil {
			utils.SendError(c, http.StatusBadRequest, "cannot register new user", err.Error())
			return
		}
		utils.SendSuccess(c, http.StatusOK, "new account registration successful", nil, nil)
	}
}
