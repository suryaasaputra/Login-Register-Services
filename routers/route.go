package routers

import (
	"dibagi/config"
	"dibagi/controllers"
	"dibagi/middlewares"
	"fmt"

	"github.com/gin-gonic/gin"
)

func StartServer(ctl controllers.Controller) error {
	r := gin.Default()

	r.POST("/register", ctl.UserController.Register)
	r.POST("/login", ctl.UserController.Login)
	r.Use(middlewares.Authentication())
	r.GET("/:userName", ctl.UserController.GetUser)
	r.PUT("/:userName", middlewares.UserAuthorization(), ctl.UserController.Update)

	return r.Run(fmt.Sprintf(":%d", config.SERVER_PORT))
}
