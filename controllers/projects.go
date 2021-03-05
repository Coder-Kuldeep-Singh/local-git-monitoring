package controllers

import (
	"fmt"
	"git-monitoring/utility"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

var (
	// ProjectList contains the projects list
	ProjectList = make(map[string]string)
)

// Projects loads the all projects till while running server
func Projects(c *gin.Context) {
	c.HTML(http.StatusOK, "projects.tmpl.html", gin.H{
		"project": ProjectList,
	})
}

// PostProject loads the new projects name and path
func PostProject(c *gin.Context) {
	err := c.Request.ParseForm()
	if err != nil {
		c.HTML(http.StatusOK, "projects.tmpl.html", gin.H{
			"project": ProjectList,
			"Error":   err.Error(),
		})
		return
	}
	project := c.Request.Form.Get("project")
	exists := isGit(project)
	if !exists {
		c.HTML(http.StatusOK, "projects.tmpl.html", gin.H{
			"project": ProjectList,
			"Error":   "Folder isn't a git repo",
		})
		return
	}
	msg := projectPath(project)
	c.HTML(http.StatusOK, "projects.tmpl.html", gin.H{
		"project": ProjectList,
		"Error":   msg,
	})
}

func isGit(projectPath string) bool {
	file, err := utility.OpenFile(projectPath)
	if err != nil {
		log.Println("error to open folder")
		return false
	}
	files := utility.GetAllFiles(file)
	return utility.CheckExpectedFile(".git", files)

}

// ProjectName returns the project name from the project path
func ProjectName(project string) string {
	ss := strings.Split(project, "/")
	s := ss[len(ss)-1]
	return s
}

func projectPath(project string) string {
	s := ProjectName(project)
	if ProjectList == nil {
		ProjectList[project] = s
		return ""
	}
	_, exists := ProjectList[project]
	if exists {
		return fmt.Sprintf("Path already Exists")
	}
	ProjectList[project] = s
	return ""
}
