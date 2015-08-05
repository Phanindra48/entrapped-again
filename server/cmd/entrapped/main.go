package main

import (
	"github.com/kgthegreat/entrapped-again/server"
	"os"
	"fmt"
)

func main() {
	envPort := os.Getenv("PORT")
	if len(envPort) != 0 {
		envPort = ":" + envPort
	}
	fmt.Println(envPort)
	entrapped.Start(envPort)
}
