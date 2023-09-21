package controller

import (
	"AnaUserService/request"
	"AnaUserService/response"
	"AnaUserService/service"
	"AnaUserService/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController(userService *service.UserService) *UserController {
	return &UserController{userService}
}

// @Summary		Get User By Id
// @Description	Get User By Id
// @Produce		json
// @Param id path string true "user id"
// @Success		200	{object} model.User
// @Router			/users/{id} [get]
func (uc *UserController) GetUserByID(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	user, err := uc.userService.GetUserByID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch user"})
		return
	}

	c.JSON(http.StatusOK, user)
}

// @Summary		Get User By Id
// @Description	Get User By Id
// @Accept		json
// @Produce		json
// @Param       credentials body request.Login true "Credentials Info"
// @Success		200	{object} model.User
// @Router			/login [post]
func (uc *UserController) Login(c *gin.Context) {
	request := request.Login{}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, response.Response{
			Status:  "error",
			Message: err.Error(),
		})
	}

	user, err := uc.userService.Login(request.Username, request.Password)
	if err != nil {
		c.JSON(400, response.Response{
			Status:  "error",
			Message: err.Error(),
		})
	}

	data, err := util.GenerateJWTToken(user.ID, user.Email)
	if err != nil {
		c.JSON(400, response.Response{
			Status:  "error",
			Message: err.Error(),
		})
	}

	c.JSON(200, response.Response{
		Status:  "success",
		Data:    data,
		Message: "success",
	})
}
