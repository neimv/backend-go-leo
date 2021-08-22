package main

import (
	"flag"
	"fmt"

	utils "github.com/neimv/backend-go-leo/utilities"
)

var (
	host = flag.String("h", "localhost", "Host to start web server")
	port = flag.Int("p", 3000, "port to access web server")
)

func main() {
	fmt.Println("Hello")
	utils.PrintBlue("Starting WebServer")
}
