package controller

import (
	"final-project3/pkg/user/dto"
	"final-project3/pkg/user/usecase"
	jwt_local "final-project3/utils/jwt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type UserHTTPController struct {
	usecase usecase.UsecaseInterfaceUser
}

func InitControllerUser(uc usecase.UsecaseInterfaceUser) *UserHTTPController {
	return &UserHTTPController{
		usecase: uc,
	}
}

func (uc *UserHTTPController) Register(c *gin.Context) {
	var req dto.UserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "status": http.StatusBadRequest})
		return
	}

	user, err := uc.usecase.Register(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// res := ConvertToUserResponse(user)

	// c.JSON(http.StatusCreated, gin.H{
	// 	"id" : user.Id,
	// 	"full_name" : user.Fullname,
	// 	"email" : user.Email,
	// 	"created_at" : user.CreatedAt,
	// })

	c.JSON(http.StatusCreated, user)
}

func (uc *UserHTTPController) Login(c *gin.Context) {
	var req dto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "status": http.StatusBadRequest})
		return
	}

	user, err := uc.usecase.Login(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "status": http.StatusBadRequest})
		return
	}

	access_token, err := jwt_local.GenerateNewToken(jwt.MapClaims{
		"id":    user.Id,
		"email": user.Email,
	})

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": access_token,
	})
}

func (uc *UserHTTPController) UpdateUserById(c *gin.Context) {
	idString := c.Param("id")
	userId, err := strconv.Atoi(idString)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid Parsing User ID", "status": http.StatusBadRequest})
		return
	}
	var input dto.UserRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "something wrong",
		})
		return
	}

	user, err := uc.usecase.UpdateUserById(userId, input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "something wrong",
			"err":     err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (uc *UserHTTPController) DeleteUserById(c *gin.Context) {

	idString := c.Param("id")
	userId, err := strconv.Atoi(idString)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid Parsing Country ID", "status": http.StatusBadRequest})
		return
	}
	err = uc.usecase.DeleteUserById(userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "something wrong",
			"err":     err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "your account has been successfully deleted",
	})
}
