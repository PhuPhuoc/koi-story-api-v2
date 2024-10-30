package markethandler

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func RegisterPostMarketRoutes(rg *gin.RouterGroup, db *sqlx.DB) {
	eg := rg.Group("/post-market")
	{
		eg.GET("", getPostMarketHandler(db))
		eg.GET("/:post_id", getDetailPostMarketHandler(db))
		eg.POST("", createNewPostMarketHandler(db))

	}
}
