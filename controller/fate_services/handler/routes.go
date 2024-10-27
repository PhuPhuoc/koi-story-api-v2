package fatehandler

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func RegisterFatesRoutes(rg *gin.RouterGroup, db *sqlx.DB) {
	eg := rg.Group("/fates")
	{
		eg.GET("", getFatesHandler(db))
		eg.POST("/user", generateUserFateHandler(db))
	}
}
