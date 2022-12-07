package main

import (
	"github.com/steelWinds/animals-rest-go/internal/cmd"
)

func main() {
	engine, _ := cmd.InitApp()

	engine.Run(":3000")
}
