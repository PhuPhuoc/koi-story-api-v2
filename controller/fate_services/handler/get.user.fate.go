package fatehandler

import (
	"net/http"

	faterepository "github.com/PhuPhuoc/koi-story-api-v2/controller/fate_services/repository"
	"github.com/PhuPhuoc/koi-story-api-v2/utils"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

//	@BasePath		/api/v1
//	@Summary		get user fate
//	@Description	get user fate
//	@Tags			fates
//	@Accept			json
//	@Produce		json
//	@Param			user_id	path		string					true	"User ID"
//	@Success		200		{object}	map[string]interface{}	"message success"
//	@Failure		400		{object}	error					"Bad request error"
//	@Router			/fates/user/{user_id} [get]
func getUserFateHandler(db *sqlx.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		user_id := c.Param("user_id")
		if user_id == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "User ID is required"})
			return
		}
		repo := faterepository.NewFateStore(db)
		data, err := repo.GetUserFate(user_id)
		if err != nil {
			utils.SendError(c, http.StatusBadRequest, "cannot get user fate", err.Error())
			return
		}
		utils.SendSuccess(c, http.StatusOK, "get user's fate successful", data, nil)
	}
}
