package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Project contains the project all info
func Project(c *gin.Context) {
	projectPath := c.Query("path")
	s := ProjectName(projectPath)
	log.Println(s)
	c.HTML(http.StatusOK, "project.tmpl.html", gin.H{
		"Project": s,
	})
}
