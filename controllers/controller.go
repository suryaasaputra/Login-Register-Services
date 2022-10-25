package controllers

type Controller struct {
	UserController *userController
}

func NewController(userController *userController) Controller {
	return Controller{
		UserController: userController,
	}

}
