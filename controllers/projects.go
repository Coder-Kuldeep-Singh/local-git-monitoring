package controllers

import (
	"fmt"
	"git-monitoring/contributors"
	"git-monitoring/utility"
	"log"
	"os"
	"strings"
)

var (
	// ProjectList contains the projects list
	ProjectList = make(map[string]string)
)

func isGit(projectPath string) bool {
	file, err := utility.OpenFile(projectPath)
	if err != nil {
		log.Println("error to open folder", err)
		return false
	}
	defer file.Close()
	files := utility.GetAllFiles(file)
	return utility.CheckExpectedFile(".git", files)

}

// ProjectName returns the project name from the project path
func ProjectName(project string) string {
	ss := strings.Split(project, "/")
	s := ss[len(ss)-1]
	return s
}

// ProjectPath checks the path already exists or not
func ProjectPath(project string) string {
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

// LoadProjectsHomePage loads the home page of all list projects
func LoadProjectsHomePage(path string) {
	Newpath := ProjectPath(path)
	if Newpath != "" {
		log.Println("path exists")
		os.Exit(2)
	}
	exists := isGit(path)
	if exists == false {
		log.Println("folder already into the list")
		os.Exit(2)
	}
	log.Println("Found git repo")
	for key, value := range ProjectList {
		log.Printf("%s ---> %s", key, value)
		list, err := contributors.ContributorList(key)
		if err != nil {
			log.Println("error to get the contributors", err)
			os.Exit(2)
		}
		d := utility.Distinct(list)
		log.Println(d)
	}
	// for key, value := range contributor {
	// 	log.Println(key, value)
	// }
}
