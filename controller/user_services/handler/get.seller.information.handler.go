package userhandler

import (
	"net/http"

	userrepository "github.com/PhuPhuoc/koi-story-api-v2/controller/user_services/repository"
	"github.com/PhuPhuoc/koi-story-api-v2/utils"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

// @BasePath		/api/v1
// @Summary		get list seller information
// @Description	get list seller information
// @Tags			users
// @Accept			json
// @Produce		json
// @Param			user_id	path		string					true	"User ID"
// @Success		200		{object}	map[string]interface{}	"data"
// @Failure		400		{object}	error					"Bad request error"
// @Router			/users/{user_id}/sellers [get]
func getSellersHandler(db *sqlx.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		user_id := c.Param("user_id")
		if user_id == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "User ID is required"})
			return
		}
		repo := userrepository.NewUserStore(db)
		list_seller, err := repo.GetSellerInfo(user_id)
		if err != nil {
			utils.SendError(c, http.StatusBadRequest, err.Error())
			return
		}
		utils.SendSuccess(c, http.StatusOK, "Get data successfully", list_seller, nil)
	}
}
