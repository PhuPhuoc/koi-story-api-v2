package userhandler

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func RegisterUserRoutes(rg *gin.RouterGroup, db *sqlx.DB) {
	eg := rg.Group("/users")
	{
		eg.GET("/hello", sayHello)
		eg.POST("/login", loginHandler(db))
		eg.POST("/register", registerHandler(db))
		eg.POST("/:user_id/sellers", registerSellerHandler(db))
		eg.GET("/:user_id/sellers", getSellersHandler(db))

	}
}
