package controller

import (
	"admin-ais/model"
	"admin-ais/repository"
	"github.com/gin-gonic/gin"
	"net/http"
)


type UserController interface {
	AddUser(*gin.Context)
	GetAllUsers(*gin.Context)
	ChangeUser(*gin.Context)
	DeleteUser(*gin.Context)
}

type userController struct {
	userRepo repository.UserRepository
}

func NewUserController(userRepo repository.UserRepository) UserController {
	return &userController{userRepo: userRepo}
}

func (u userController) AddUser(c *gin.Context) {
	var user model.User

	err := c.ShouldBind(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	addUser, err := u.userRepo.AddUser(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": addUser})
}

func (u userController) GetAllUsers(c *gin.Context) {

	users, err := u.userRepo.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": users})

}

func (u userController) ChangeUser(c *gin.Context) {

	var user model.User
	idParam := c.Param("id")
	err := c.ShouldBind(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	err = u.userRepo.ChangeUser(user, idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Данные успешно обновлены"})
}

func (u userController) DeleteUser(c *gin.Context) {

	idParam := c.Param("id")
	err := u.userRepo.DeleteUser(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Пользователь удален"})
}
