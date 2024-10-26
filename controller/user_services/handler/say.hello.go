package userhandler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @BasePath		/api/v1
// @Summary		hello
// @Description	hello
// @Tags			users
// @Accept			json
// @Produce		json
// @Success		200	{object}	string	"say hi"
// @Router			/users/hello [get]
func sayHello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"hello": "This is greeting from Phu Phuoc linux server"})
}
