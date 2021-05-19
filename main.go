package main

import (
	"os"

	"front-end/router"
)

func main() {

	port := ":" + os.Getenv("PORT")

	routerInit := router.InitRouter()

	if port == ":" {
		routerInit.Run()
	} else {
		routerInit.Run(port)
	}
}
