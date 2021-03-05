package main

import (
	"flag"
	"git-monitoring/controllers"
	"log"
	"os"
)

func main() {
	path := flag.String("p", "", "path of the git folder")
	flag.Parse()
	if *path == "" {
		flag.Usage()
		os.Exit(2)
	}
	log.Println(*path)
	controllers.LoadProjectsHomePage(*path)
}
