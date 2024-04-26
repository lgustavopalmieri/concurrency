package user_controller

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/lgustavopalmieri/labs-go-expert-auctiont/configuration/rest_err"
	"github.com/lgustavopalmieri/labs-go-expert-auctiont/internal/usecase/user_usecase"
)

type UserController struct {
	userUsecase user_usecase.UserUseCaseInterface
}

func NewUserController(userUseCase user_usecase.UserUseCaseInterface) *UserController {
	return &UserController{
		userUsecase: userUseCase,
	}
}

func (u *UserController) FindUserById(c *gin.Context) {
	userId := c.Param("userId")
	if err := uuid.Validate(userId); err != nil {
		errRest := rest_err.NewBadRequestError("Invalid fields", rest_err.Causes{
			Field:   "userId",
			Message: "Invalid UUID value",
		})
		c.JSON(errRest.Code, errRest)
		return
	}
	userData, err := u.userUsecase.FindUserById(context.Background(), userId)
	if err != nil {
		errRest := rest_err.ConvertError(err)
		c.JSON(errRest.Code, errRest)
		return
	}
	c.JSON(http.StatusOK, userData)
}
