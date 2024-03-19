package controller

import (
	"net/http"

	"github.com/MrBooi/go_chat_backend/bootstrap"
	"github.com/MrBooi/go_chat_backend/domain"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserController struct {
	UserUsecase domain.UserUsecase
	Env         *bootstrap.Env
}

/**
 * Gets the current user's profile
 */
func (uc *UserController) Profile(c *gin.Context) {
	userID := c.GetString("x-user-id")
	_, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	user, err := uc.UserUsecase.GetProfile(c, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

/**
 * Update current user's account
 */
func (uc *UserController) UpdateUser(c *gin.Context) {
	var body domain.UpdateUserRequest

	err := c.ShouldBind(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	userID := c.GetString("x-user-id")
	_, err = primitive.ObjectIDFromHex(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	user, err := uc.UserUsecase.UpdateUser(c, userID, body)

	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}
