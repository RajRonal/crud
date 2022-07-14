package main

import (
	"awesomeProject1/database"
	"awesomeProject1/server"
	"fmt"
)

func main() {
	err := database.ConnectAndMigrate("localhost", "5433", "tutorial", "local", "local", database.SSLModeDisable)
	if err != nil {
		panic(err)
	}

	fmt.Println("connected")
	srv := server.SetupRoutes()
	err = srv.Run(":8080")
	if err != nil {
		panic(err)
	}
}
