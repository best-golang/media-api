package infrastructure

import (
	"api/interfaces/controllers"

	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func init() {
	router := gin.Default()

	sqlHandler := NewSqlHandler()
	Migrate(sqlHandler.Conn)
	// controller
	userController := controllers.NewUserController(sqlHandler)

	// Define routes
	router.POST("/users", func(c *gin.Context) { userController.Create(c) })
	router.GET("/users", func(c *gin.Context) { userController.Index(c) })
	Router = router
}
