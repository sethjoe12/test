package controllers

import (
	"log"
	"net/http"
	"test/database"
	"test/models"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

//go:generate mockgen -destination=mocks/mock_Users.go -package=mocks test/controllers Iuser

type Iuser interface {
	AllUsers(context *gin.Context)
	GetUserById(context *gin.Context)
	UpdateUserById(context *gin.Context)
	DeleteUserById(context *gin.Context)
}

type User struct {
	Iuser Iuser
}

func AllUser(context *gin.Context) {
	var user []models.User

	Instance, err := database.DbConnectAndMigrate()
	if err != nil {
		log.Println(err)
	}

	if err := Instance.Find(&user).Error; err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, user)
}

func GetUserById(context *gin.Context) {
	var user models.User

	Instance, err := database.DbConnectAndMigrate()
	if err != nil {
		log.Println(err)
	}

	if err := Instance.Where("id= ?", context.Param("id")).First(&user).Error; err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	context.JSON(http.StatusOK, user)
}

func UpdateUserById(context *gin.Context) {

	var user models.User

	Instance, err := database.DbConnectAndMigrate()
	if err != nil {
		log.Println(err)
	}

	if err := Instance.Where("id= ?", context.Param("id")).First(&user).Error; err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	var updateUser models.UpdateUser

	if err := context.ShouldBindJSON(&updateUser); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := Instance.Model(&user).Updates(models.User{Name: updateUser.Name}).Error; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, user)

}

func DeleteUserById(context *gin.Context) {

	var user models.User

	Instance, err := database.DbConnectAndMigrate()
	if err != nil {
		log.Println(err)
	}

	if err := Instance.Where("id= ?", context.Param("id")).First(&user).Error; err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	if err := Instance.Delete(&user).Error; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "User deleted"})

}
