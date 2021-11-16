package controller

import (
	"admin-ais/model"
	"admin-ais/repository"
	"github.com/gin-gonic/gin"
	"net/http"
)

type RoleController interface {
	AddRole(*gin.Context)
	GetAllRoles(*gin.Context)
	ChangeRole(*gin.Context)
	DeleteRole(*gin.Context)
}

// задай себе вопрос, почему мне интерфейс?
type roleController struct {
	roleRepo repository.RoleRepository
}

func NewRoleController(roleRepo repository.RoleRepository) RoleController {
	return &roleController{roleRepo: roleRepo}
}

func (r roleController) AddRole(c *gin.Context) {
	var role model.Role

	err := c.ShouldBind(&role)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	addRole, err := r.roleRepo.AddRole(role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, addRole)
}

func (r roleController) GetAllRoles(c *gin.Context) {

	roles, err := r.roleRepo.GetAllRoles()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return

	}
	c.JSON(http.StatusOK, roles)
}

func (r roleController) ChangeRole(c *gin.Context) {
	idParam := c.Param("id")
	var role model.Role
	err := c.ShouldBindJSON(&role)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}

	err = r.roleRepo.ChangeRole(role, idParam)

	//if err.Error() == "record not found"{
	// c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// return
	//}
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return

	}

	c.JSON(http.StatusOK, gin.H{"message": "Данные успешно обновленны"})

}

func (r roleController) DeleteRole(c *gin.Context) {
	idParam := c.Param("id")
	var role model.Role
	err := c.ShouldBindJSON(&role)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}
	err = r.roleRepo.DeleteRole(role, idParam)
	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return

	}
	c.JSON(http.StatusOK, gin.H{"Status": "Успешно удаленно"})

}
