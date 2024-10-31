package consulthandler

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func RegisterPostConsultRoutes(rg *gin.RouterGroup, db *sqlx.DB) {
	eg := rg.Group("/post-consult")
	{
		eg.POST("", createNewPostMarketHandler(db))
		eg.PUT("/:post_id", updatePostMarketHandler(db))
		eg.GET("", getPostConsultHandler(db))
		eg.GET("/user/:user_id", getMyPostConsultHandler(db))
	}
}
