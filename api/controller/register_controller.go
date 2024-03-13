package controller

import (
	"net/http"

	"github.com/MrBooi/go_chat_backend/bootstrap"
	"github.com/MrBooi/go_chat_backend/domain"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RegisterController struct {
	RegisterUsecase domain.RegisterUsecase
	Env             *bootstrap.Env
}

func (rc *RegisterController) Register(c *gin.Context) {
	var request domain.RegisterRequest

	err := c.ShouldBind(&request)

	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	_, err = rc.RegisterUsecase.GetUserByUuidOrEmail(c, request.Uuid, request.Email)

	if err == nil {
		c.JSON(http.StatusConflict, domain.ErrorResponse{Message: "User already exists with the given Uuid or Email"})
		return
	}

	user := domain.User{
		ID:       primitive.NewObjectID(),
		Name:     request.Name,
		Email:    request.Email,
		PhotoUrl: request.PhotoUrl,
	}

	err = rc.RegisterUsecase.Create(c, &user)

	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, "user created successfully")

}
