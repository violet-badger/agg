package main

import (
	"example/web-service-gin/api"
	"os"
)

func main() {
	var exitCode int
	defer func() {
		os.Exit(exitCode)
	}()

	router := api.NewRouter()

	err := router.Run(":8080")
	if err != nil {
	}

}
