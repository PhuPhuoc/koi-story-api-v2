package bloghandler

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func RegisterPostBlogRoutes(rg *gin.RouterGroup, db *sqlx.DB) {
	eg := rg.Group("/post-blog")
	{
		eg.POST("", createNewPostBlogHandler(db))
	}
}
