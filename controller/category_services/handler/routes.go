package categoryhandler

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func RegisterCategoryRoutes(rg *gin.RouterGroup, db *sqlx.DB) {
	eg := rg.Group("/categories")
	{
		eg.GET("", getCategoriesHandler(db))
		eg.POST("", createNewCategoryHandler(db))
		eg.PUT("/:cate_id", updateCategoryHandler(db))
	}
}
