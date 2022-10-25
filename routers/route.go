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
	r.GET("/", ctl.HomeController)
	r.POST("/register", ctl.UserController.Register)
	r.POST("/login", ctl.UserController.Login)

	userRouter := r.Group("/user")
	{
		userRouter.Use(middlewares.Authentication())
		userRouter.GET("/:userName", ctl.UserController.GetUser)
		userRouter.PUT("/:userName", middlewares.UserAuthorization(), ctl.UserController.Update)
	}

	return r.Run(":" + PORT)
}
