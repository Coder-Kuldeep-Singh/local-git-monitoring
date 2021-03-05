package routers

import (
	"git-monitoring/controllers"

	"github.com/gin-gonic/gin"
)

// SetupRouter initialize all the urls
func SetupRouter() *gin.Engine {
	router := gin.Default()
	// gin.SetMode(os.Getenv("GIN_MODE"))
	router.LoadHTMLGlob("./html/templates/*")
	// router.Static("/home", "./html/static")
	router.GET("/projects", controllers.Projects)
	router.POST("/projects", controllers.PostProject)
	router.GET("/project/init", controllers.Project)
	return router
}
