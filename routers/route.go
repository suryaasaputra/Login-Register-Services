package routers

import (
	"dibagi/controllers"
	"dibagi/middlewares"
	"os"

	"github.com/gin-gonic/gin"
)

var PORT = os.Getenv("PORT")

func StartServer(ctl controllers.Controller) error {
	r := gin.Default()

	r.POST("/register", ctl.UserController.Register)
	r.POST("/login", ctl.UserController.Login)
	r.Use(middlewares.Authentication())
	r.GET("/:userName", ctl.UserController.GetUser)
	r.PUT("/:userName", middlewares.UserAuthorization(), ctl.UserController.Update)

	return r.Run(":" + PORT)
}
