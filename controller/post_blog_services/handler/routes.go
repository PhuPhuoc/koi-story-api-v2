package bloghandler

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func RegisterPostBlogRoutes(rg *gin.RouterGroup, db *sqlx.DB) {
	eg := rg.Group("/post-blog")
	{
		eg.POST("", createNewPostBlogHandler(db))
		eg.GET("/category/:category_id", getBlogByCategoryIDHandler(db))
		eg.GET("/:post_id", getDetailBlogHandler(db))
		eg.GET("/recommend/:user_id", getRecommendBlogHandler(db))
	}
}
