package controller

import (
	"admin-ais/model"
	"admin-ais/repository"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type RightController interface {
	AddRight(*gin.Context)
	GetAllRights(*gin.Context)
	ChangeRight(*gin.Context)
	DeleteRight(*gin.Context)
}


type rightController struct {
	rightRepo repository.RightRepository

}




func NewRightController(rightRepo repository.RightRepository) 	RightController{
	return &rightController{rightRepo: rightRepo}
}

func (r rightController) AddRight(c *gin.Context) {
	var right model.Right

	err := c.ShouldBind(&right)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	addRight, err := r.rightRepo.AddRight(right)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, addRight)
}

func (r rightController) GetAllRights(c *gin.Context) {

	rights, err := r.rightRepo.GetAllRights()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return

	}
	c.JSON(http.StatusOK, rights)
}

func (r rightController) ChangeRight(c *gin.Context) {
	idParam := c.Param("id")
	var right model.Right
	err := c.ShouldBindJSON(&right)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}
	fmt.Println(right, "обновленные данные")
	fmt.Println(idParam)
	 err = r.rightRepo.ChangeRight(right, idParam)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}

	c.JSON(http.StatusOK, gin.H{"message": "права успешно обновлены"})

}

func (r rightController) DeleteRight(c *gin.Context) {
	idParam := c.Param("id")
	//var right model.Right
	//err := c.ShouldBindJSON(&right)
	//if err != nil {
	//	c.JSON(http.StatusBadRequest, gin.H{"error": err})
	//}
	err := r.rightRepo.DeleteRight( idParam)
	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return

	}
	c.JSON(http.StatusOK, gin.H{"Status": "Успешно удаленно"})

}
