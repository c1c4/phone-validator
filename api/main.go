package main

import (
	"api/app"
	"fmt"
)

func main() {
	fmt.Println("The phone validator api is starting")
	serverReady := make(chan bool)
	app := app.App{
		ServerReady: serverReady,
	}

	app.StartApp()
}
