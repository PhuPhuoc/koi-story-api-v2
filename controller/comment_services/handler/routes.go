package commenthandler

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func RegisterCommentRoutes(rg *gin.RouterGroup, db *sqlx.DB) {
	eg := rg.Group("posts/:post_id/comments")
	{
		eg.POST("", createCommentHandler(db))
		eg.GET("", getAllCommentHandler(db))
	}
	rg.DELETE("comments/:comment_id", deleteCommentHandler(db))
	rg.PUT("comments/:comment_id", updateCommentHandler(db))

}
