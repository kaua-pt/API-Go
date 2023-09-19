package main

import (
	"fmt"

	"main.go/config"
	"main.go/router"
)

func main() {
	e := config.Init()

	if e != nil {
		fmt.Println(e)
		return
	}

	router.Initialize()
}
