package consulthandler

import (
	"net/http"

	consultrepository "github.com/PhuPhuoc/koi-story-api-v2/controller/post_consult_services/repository"
	"github.com/PhuPhuoc/koi-story-api-v2/utils"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

//	@BasePath		/api/v1
//	@Summary		get my consult's posts
//	@Description	get my consult's posts
//	@Tags			post consult
//	@Accept			json
//	@Produce		json
//	@Param			user_id	path		string					true	"User ID"
//	@Success		200		{object}	map[string]interface{}	"data"
//	@Failure		400		{object}	error					"Bad request error"
//	@Router			/post-consult/user/{user_id} [get]
func getMyPostConsultHandler(db *sqlx.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		user_id := c.Param("user_id")
		if user_id == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "User ID is required"})
			return
		}
		repo := consultrepository.NewPostConsultStore(db)
		list_post, err := repo.GetMyListPostConsult(user_id)
		if err != nil {
			utils.SendError(c, http.StatusBadRequest, "cannot get my list post consult", err.Error())
			return
		}
		utils.SendSuccess(c, http.StatusOK, "Get data successfully", list_post, nil)
	}
}
