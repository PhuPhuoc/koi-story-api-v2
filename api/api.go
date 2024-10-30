package api

import (
	"fmt"

	categoryhandler "github.com/PhuPhuoc/koi-story-api-v2/controller/category_services/handler"
	commenthandler "github.com/PhuPhuoc/koi-story-api-v2/controller/comment_services/handler"
	fatehandler "github.com/PhuPhuoc/koi-story-api-v2/controller/fate_services/handler"
	markethandler "github.com/PhuPhuoc/koi-story-api-v2/controller/post_market_services/handler"
	userhandler "github.com/PhuPhuoc/koi-story-api-v2/controller/user_services/handler"
	docs "github.com/PhuPhuoc/koi-story-api-v2/docs"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type server struct {
	address string
	db      *sqlx.DB
}

func InitServer(addr string, db *sqlx.DB) *server {
	return &server{
		address: addr,
		db:      db,
	}
}

func (sv *server) RunApp() error {
	docs.SwaggerInfo.BasePath = "/api/v1"
	router := gin.New()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	router.Use(
		gin.LoggerWithWriter(gin.DefaultWriter, "/swagger/*any"),
		gin.Recovery(),
		cors.New(config),
	)

	v1 := router.Group("/api/v1")
	sv.registerRoutes(v1)
	sv.runLog()
	return router.Run(sv.address)
}

func (sv *server) runLog() {
	fmt.Printf("\nFor development: http://localhost%s/swagger/index.html\n\n", sv.address)
}

func (sv *server) registerRoutes(v1 *gin.RouterGroup) {
	userhandler.RegisterUserRoutes(v1, sv.db)
	fatehandler.RegisterFatesRoutes(v1, sv.db)
	categoryhandler.RegisterCategoryRoutes(v1, sv.db)
	markethandler.RegisterPostMarketRoutes(v1, sv.db)
	commenthandler.RegisterCommentRoutes(v1, sv.db)
}
