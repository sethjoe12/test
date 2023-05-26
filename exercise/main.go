package main

import (
	"test/controllers"
	"test/database"

	// _ "test/docs"
	"test/middlewares"

	"github.com/gin-gonic/gin"
	// swagger_files "github.com/swaggo/files"
	// ginSwagger_test "github.com/swaggo/gin-swagger"
)

func main() {

	database.DbConnectAndMigrate()
	//database connect and migration

	router := initRouter()
	router.Run(":12000")

}
func initRouter() *gin.Engine {
	router := gin.Default()

	//swagger router
	// TODO router.GET("/docs/*any", ginSwagger_test.WrapGandler(swagger_files.Handler))
	api := router.Group("/api")

	{
		// all v1 enpoint routes
		api.POST("/v1/user/register", controllers.RegisterUser)
		api.POST("/v1/token", controllers.GenerateToken)

		v2 := api.Group("/v2").Use(middlewares.Auth())
		{

			//all v2 enpoints routes

			v2.GET("/allusers", controllers.AllUser)           //get all registered users
			v2.GET("/user/:id", controllers.GetUserById)       //get single user by id
			v2.PUT("/user/:id", controllers.UpdateUserById)    //update single user by id
			v2.DELETE("/user/:id", controllers.DeleteUserById) //delete single user by id

		}

	}
	return router
}
