package controllers

import (
	"dibagi/helpers"
	"dibagi/models"
	"dibagi/repository"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userController struct {
	UserRepository repository.IUserRepository
}

func NewUserController(userRepository repository.IUserRepository) *userController {
	return &userController{
		UserRepository: userRepository,
	}
}

func (u userController) Register(ctx *gin.Context) {
	var User = models.User{}
	err := ctx.ShouldBindJSON(&User)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}

	resp, err := u.UserRepository.RegisterUser(User)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "success",
		"data":    resp,
	})

}

func (u userController) Login(ctx *gin.Context) {
	var User = models.UserLoginRequest{}
	err := ctx.ShouldBindJSON(&User)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}

	userResult := u.UserRepository.GetUserByEmail(User.Email)

	if userResult.Email == "" {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"status": http.StatusNotFound,
			"error":  fmt.Sprintf("User with email '%s' does not exist", User.Email),
		})
		return
	}

	isPasswordCorrect := helpers.ComparePassword([]byte(userResult.Password), []byte(User.Password))

	if !isPasswordCorrect {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"status": http.StatusUnauthorized,
			"error":  "Password incorrect",
		})
		return
	}

	token, err := helpers.GenerateToken(userResult.ID, userResult.UserName)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
			"error":  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

func (u userController) GetUser(ctx *gin.Context) {
	userNameURL := ctx.Param("userName")
	userResult := u.UserRepository.GetUserByUserName(userNameURL)
	if userResult.UserName == "" {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"status": http.StatusNotFound,
			"error":  fmt.Sprintf("User with username '%s' not found", userNameURL),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    userResult,
	})
}

func (u userController) Update(ctx *gin.Context) {
	userNameURL := ctx.Param("userName")
	var User = models.User{}
	err := ctx.ShouldBindJSON(&User)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  "body required",
		})
		return
	}

	result := u.UserRepository.EditUser(userNameURL, User)
	if result.UserName == "" {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"status": http.StatusNotFound,
			"error":  fmt.Sprintf("User with username '%s' not found", userNameURL),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "update success",
		"data":    result,
	})

}
