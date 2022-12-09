package main

import (
	"fmt"
	"os"

	"github.com/steelWinds/animals-rest-go/internal/cmd"
)

func main() {
	var APP_PORT = os.Getenv("APP_PORT")

	engine, _ := cmd.InitApp()

	addr := fmt.Sprintf(":%s", APP_PORT)

	engine.Run(addr)
}
