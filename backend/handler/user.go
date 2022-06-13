package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/tris-tux/go-task-gin/backend/helper"
	"github.com/tris-tux/go-task-gin/backend/schema"
	"github.com/tris-tux/go-task-gin/backend/service"
)

//UserHandler is a ....
type UserHandler interface {
	Update(context *gin.Context)
	Profile(context *gin.Context)
}

type userHandler struct {
	userService service.UserService
	jwtService  service.JWTService
}

//NewUserHandler is creating anew instance of UserControlller
func NewUserHandler(userService service.UserService, jwtService service.JWTService) UserHandler {
	return &userHandler{
		userService: userService,
		jwtService:  jwtService,
	}
}

func (c *userHandler) Update(context *gin.Context) {
	var userUpdateDTO schema.UserUpdateDTO
	errDTO := context.ShouldBind(&userUpdateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	// authHeader := context.GetHeader("Authorization")
	authorizationHeader := context.GetHeader("Authorization")
	if !strings.Contains(authorizationHeader, "Bearer") {
		helper.BuildErrorResponse("Invalid token", "Bareer", http.StatusBadRequest)
		return
	}

	authHeader := strings.Replace(authorizationHeader, "Bearer ", "", -1)

	token, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	id, err := strconv.ParseUint(fmt.Sprintf("%v", claims["user_id"]), 10, 64)
	if err != nil {
		panic(err.Error())
	}
	userUpdateDTO.ID = id
	u := c.userService.Update(userUpdateDTO)
	res := helper.BuildResponse(true, "OK!", u)
	context.JSON(http.StatusOK, res)
}

func (c *userHandler) Profile(context *gin.Context) {
	// authHeader := context.GetHeader("Authorization")
	authorizationHeader := context.GetHeader("Authorization")
	if !strings.Contains(authorizationHeader, "Bearer") {
		helper.BuildErrorResponse("Invalid token", "Bareer", http.StatusBadRequest)
		return
	}

	authHeader := strings.Replace(authorizationHeader, "Bearer ", "", -1)

	token, err := c.jwtService.ValidateToken(authHeader)
	if err != nil {
		panic(err.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["user_id"])
	user := c.userService.Profile(id)
	res := helper.BuildResponse(true, "OK", user)
	context.JSON(http.StatusOK, res)

}
