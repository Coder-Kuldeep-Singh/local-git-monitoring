package main

import (
	"git-monitoring/routers"
)

func main() {
	app := routers.SetupRouter()
	app.Run(":8000")
}
