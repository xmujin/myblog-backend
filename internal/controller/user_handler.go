package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xmujin/myblog-backend/internal/model"
	"github.com/xmujin/myblog-backend/internal/service"
)

type UserController struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) UserController {
	return UserController{
		userService: userService,
	}
}

func (u UserController) Register() gin.HandlerFunc {
	return func(c *gin.Context) {
		var userDto model.UserRegisterDto
		if err := c.ShouldBindJSON(&userDto); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err := u.userService.Register(&userDto)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"注册失败": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"mesage": "注册成功",
		})
	}
}

func (u UserController) Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		var userDto model.UserLoginDto
		if err := c.ShouldBindJSON(&userDto); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		token, err := u.userService.Login(&userDto)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"登录失败": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"token": token,
		})

	}
}
