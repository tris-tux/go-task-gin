package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tris-tux/go-task-gin/backend/helper"
	"github.com/tris-tux/go-task-gin/backend/schema"
	"github.com/tris-tux/go-task-gin/backend/service"
)

//AuthHandler interface is a contract what this handler can do
type AuthHandler interface {
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
}

type authHandler struct {
	authService service.AuthService
	jwtService  service.JWTService
}

//NewAuthHandler creates a new instance of AuthHandler
func NewAuthHandler(authService service.AuthService, jwtService service.JWTService) AuthHandler {
	return &authHandler{
		authService: authService,
		jwtService:  jwtService,
	}
}

func (c *authHandler) Login(ctx *gin.Context) {
	var loginDTO schema.LoginDTO
	errDTO := ctx.ShouldBindJSON(&loginDTO)
	if errDTO != nil {
		response := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	authResult := c.authService.VerifyCredential(loginDTO.Email, loginDTO.Password)
	if v, ok := authResult.(schema.User); ok {
		generatedToken := c.jwtService.GenerateToken(strconv.FormatUint(v.ID, 10))
		v.Token = generatedToken
		response := helper.BuildResponse(true, "OK!", v)
		ctx.JSON(http.StatusOK, response)
		return
	}
	response := helper.BuildErrorResponse("Please check again your credential", "Invalid Credential", helper.EmptyObj{})
	ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
}

func (c *authHandler) Register(ctx *gin.Context) {
	var registerDTO schema.RegisterDTO
	errDTO := ctx.ShouldBindJSON(&registerDTO)
	if errDTO != nil {
		response := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	if !c.authService.IsDuplicateEmail(registerDTO.Email) {
		response := helper.BuildErrorResponse("Failed to process request", "Duplicate email", helper.EmptyObj{})
		ctx.JSON(http.StatusConflict, response)
	} else {
		createdUser := c.authService.CreateUser(registerDTO)
		token := c.jwtService.GenerateToken(strconv.FormatUint(createdUser.ID, 10))
		createdUser.Token = token
		response := helper.BuildResponse(true, "OK!", createdUser)
		ctx.JSON(http.StatusCreated, response)
	}
}
