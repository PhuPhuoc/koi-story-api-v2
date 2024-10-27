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
// @Summary		register new seller information
// @Description	register new seller information
// @Tags			users
// @Accept			json
// @Produce		json
// @Param			user_id	path		string					true	"User ID"
// @Param			user	body		usermodel.SellerInfo	true	"Seller register data"
// @Success		200		{object}	map[string]interface{}	"message success"
// @Failure		400		{object}	error					"Bad request error"
// @Router			/users/{user_id}/sellers [post]
func registerSellerHandler(db *sqlx.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		user_id := c.Param("user_id")
		if user_id == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "User ID is required"})
			return
		}
		var req usermodel.SellerInfo
		if err := c.ShouldBindJSON(&req); err != nil {
			utils.SendError(c, http.StatusBadRequest, "invalid request body")
			return
		}
		repo := userrepository.NewUserStore(db)
		if err := repo.RegisterSeller(user_id, &req); err != nil {
			utils.SendError(c, http.StatusBadRequest, err.Error())
			return
		}
		utils.SendSuccess(c, http.StatusOK, "new account registration successful", nil, nil)
	}
}
