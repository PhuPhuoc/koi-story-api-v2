package consulthandler

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func RegisterPostConsultRoutes(rg *gin.RouterGroup, db *sqlx.DB) {
	eg := rg.Group("/post-consult")
	{
		eg.POST("", createNewPostConsultHandler(db))
		eg.PUT("/:post_id", updatePostConsultHandler(db))
		eg.GET("", getPostConsultHandler(db))
		eg.GET("/user/:user_id", getMyPostConsultHandler(db))
	}
}
