package routers

import (
	"dibagi/controllers"
	"dibagi/middlewares"
	"os"

	"github.com/gin-gonic/gin"
)

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

	var PORT = os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	return r.Run(":" + PORT)
}
