package imagehandler

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func RegisterImageRoutes(rg *gin.RouterGroup, db *sqlx.DB) {
	rg.POST("/posts/:post_id/images", addNewImageHandler(db))
	rg.DELETE("/images/:image_id", deleteImageHandler(db))

}
